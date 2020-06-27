# git-proxy

[中文](README.zhCN.md)

make git consume proxy env

## How to install

```shell
go get github.com/fishead/git-proxy
```

## How to use

Just prepend `proxy` before `clone`, `fetch` or `push`

```shell
# set http proxy
export HTTP_PROXY=http://127.0.0.1:1087

# or socks5 proxy
export SOCKS_SERVER=127.0.0.1:1087

# then
git proxy clone git@github.com:fishead/git-proxy.git

# or
git proxy clone https://github.com/fishead/git-proxy.git
```

## Is it awesome

Yes.
