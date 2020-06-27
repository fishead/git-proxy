package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/net/proxy"
)

func main() {
	if os.Args[1] == "relay" {
		relay(os.Args[2], os.Args[3], os.Args[4])
		return
	}

	args := []string{}

	proxyServer, ok := getProxyEnv()
	if ok && len(proxyServer) > 0 {
		gitCmd, ok := findFirstNoneFlag(os.Args[1:])
		if !ok {
			help()
			os.Exit(2)
			return
		}

		var (
			repoURL string
		)
		switch gitCmd {
		case "clone":
			repoURL, ok = findFirstNoneFlag(os.Args[2:])
		case "fetch":
			repo, ok := findFirstNoneFlag(os.Args[2:])
			if ok {
				output, err := exec.Command("git", "remote", "get-url", repo).Output()
				if err == nil {
					repoURL = string(output)
				}
			}
		case "push":
			repo, ok := findFirstNoneFlag(os.Args[2:])
			if ok {
				output, err := exec.Command("git", "remote", "get-url", "--push", repo).Output()
				if err == nil {
					repoURL = string(output)
				}
			}
		}

		var protocol string
		if ok {
			protocol, _ = getProtocol(repoURL)
		}

		switch protocol {
		case ProtocolHTTP, ProtocolHTTPS:
			// for http or https protocol
			args = append(args, "-c", fmt.Sprintf(`http.proxy=%s`, proxyServer))
		case ProtocolSSH:
			args = append(args, "-c",
				fmt.Sprintf(`core.sshCommand=ssh -o ProxyCommand="git proxy relay %s %%h %%p"`, proxyServer))
		case ProtocolGit:
			args = append(args, "-c", fmt.Sprintf(`core.gitProxy="git proxy relay %s %%h %%p"`, proxyServer))
		}
	}

	args = append(args, os.Args[1:]...)
	callGit(args)
}

func getProxyEnv() (string, bool) {
	socksProxy, ok := os.LookupEnv("SOCKS_SERVER")
	if ok && len(socksProxy) > 0 {
		return socksProxy, ok
	}

	socksProxy, ok = os.LookupEnv("socks_server")
	if ok && len(socksProxy) > 0 {
		return socksProxy, ok
	}

	httpProxy, ok := os.LookupEnv("HTTP_PROXY")
	if ok && len(httpProxy) > 0 {
		return httpProxy, ok
	}

	httpProxy, ok = os.LookupEnv("http_proxy")
	if ok && len(httpProxy) > 0 {
		return httpProxy, ok
	}

	httpsProxy, ok := os.LookupEnv("HTTPS_PROXY")
	if ok && len(httpsProxy) > 0 {
		return httpsProxy, ok
	}

	return os.LookupEnv("https_proxy")
}

func help() {
	log.Println("help")
}

func findFirstNoneFlag(args []string) (string, bool) {
	for _, s := range args {
		if !strings.HasPrefix(s, "-") {
			return s, true
		}
	}

	return "", false
}

func callGit(arg []string) error {
	cmd := exec.Command("git", arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func relay(s string, h string, p string) error {
	var conn net.Conn
	if strings.HasPrefix(s, "http") {
		s := strings.TrimPrefix(s, "https://")
		s = strings.TrimPrefix(s, "http://")
		proxyConn, err := net.Dial("tcp", s)
		if err != nil {
			return err
		}
		defer proxyConn.Close()

		helo := fmt.Sprintf("CONNECT %s HTTP/1.1\r\nHost: %s \r\n", s, s)
		proxyConn.Write([]byte(helo))

		conn, err = net.Dial("tcp", net.JoinHostPort(h, p))
		if err != nil {
			return err
		}

		go func() {
			_, err = io.Copy(proxyConn, conn)
			if err != nil {
				log.Println(err)
			}
		}()

		go func() {
			_, err = io.Copy(conn, proxyConn)
			if err != nil {
				log.Println(err)
			}
		}()
	} else {
		// 假设是 socks5 代理
		u, err := url.Parse(fmt.Sprintf("socks5://%s", s))
		if err != nil {
			return err
		}

		dialer, err := proxy.FromURL(u, nil)
		if err != nil {
			return err
		}

		conn, err = dialer.Dial("tcp", net.JoinHostPort(h, p))
		if err != nil {
			return err
		}
	}
	defer conn.Close()

	var (
		inFinished  = make(chan bool)
		outFinished = make(chan bool)
	)

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Println(err)
		}
		inFinished <- true
	}()
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Println(err)
		}
		outFinished <- true
	}()

	<-inFinished
	<-outFinished

	return nil
}
