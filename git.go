package main

import (
	"errors"
	"regexp"
)

// protocols supported by git
const (
	ProtocolSSH   = "ssh"
	ProtocolGit   = "git"
	ProtocolHTTP  = "http"
	ProtocolHTTPS = "https"
	ProtocolFTP   = "ftp"
	ProtocolFTPS  = "ftps"
	ProtocolFile  = "file"
)

var (
	regexpGit   = regexp.MustCompile(`^git:`)
	regexpHTTP  = regexp.MustCompile(`^http:`)
	regexpHTTPS = regexp.MustCompile(`^https:`)
	regexpFTP   = regexp.MustCompile(`^ftp:`)
	regexpFTPS  = regexp.MustCompile(`^ftps:`)
	regexpSSH   = regexp.MustCompile(`^(ssh|git):|^([^@]*@)?[^./]+(\.[^./]+)+(\/[^/]*)+`)
	regexpFILE  = regexp.MustCompile(`^file:|(\.)?(\/[^/]+)+`)
)

func getProtocol(s string) (string, error) {
	b := []byte(s)

	if regexpGit.Match(b) {
		return ProtocolGit, nil
	}

	if regexpHTTP.Match(b) {
		return ProtocolHTTP, nil
	}

	if regexpHTTPS.Match(b) {
		return ProtocolHTTPS, nil
	}

	if regexpFTP.Match(b) {
		return ProtocolFTP, nil
	}

	if regexpFTPS.Match(b) {
		return ProtocolFTPS, nil
	}

	if regexpSSH.Match(b) {
		return ProtocolSSH, nil
	}

	if regexpFILE.Match(b) {
		return ProtocolFile, nil
	}

	return "", errors.New("error")
}
