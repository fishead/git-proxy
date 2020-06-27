# git-proxy

[English](README.md)

使 git 可以轻松使用网络代理的环境变量

## 如何安装

```shell
go get github.com/fishead/git-proxy
```

## 如何使用

只需要在`clone`、`fetch`或`push`之前加上`proxy`

```shell
# 设置 http 代理
export HTTP_PROXY=http://127.0.0.1:1087

# 或者设置 socks5 代理
export SOCKS_SERVER=127.0.0.1:1087

# 然后
git proxy clone git@github.com:fishead/git-proxy.git

# 或者
git proxy clone https://github.com/fishead/git-proxy.git
```

## Is it awesome

Yes.
