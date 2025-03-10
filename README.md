[![ci](https://github.com/sinlov-go/zlog-zap-wrapper/actions/workflows/ci.yml/badge.svg)](https://github.com/sinlov-go/zlog-zap-wrapper/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/zlog-zap-wrapper?label=go.mod)](https://github.com/sinlov-go/zlog-zap-wrapper)
[![GoDoc](https://godoc.org/github.com/sinlov-go/zlog-zap-wrapper?status.png)](https://godoc.org/github.com/sinlov-go/zlog-zap-wrapper)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov-go/zlog-zap-wrapper)](https://goreportcard.com/report/github.com/sinlov-go/zlog-zap-wrapper)

[![GitHub license](https://img.shields.io/github/license/sinlov-go/zlog-zap-wrapper)](https://github.com/sinlov-go/zlog-zap-wrapper)
[![codecov](https://codecov.io/gh/sinlov-go/zlog-zap-wrapper/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/zlog-zap-wrapper)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov-go/zlog-zap-wrapper)](https://github.com/sinlov-go/zlog-zap-wrapper/tags)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/zlog-zap-wrapper)](https://github.com/sinlov-go/zlog-zap-wrapper/releases)

## for what

- this project used to fast config [zap logger](https://pkg.go.dev/go.uber.org/zap)

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/zlog-zap-wrapper)](https://github.com/sinlov-go/zlog-zap-wrapper/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov-go/zlog-zap-wrapper.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov-go/zlog-zap-wrapper
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov-go/zlog-zap-wrapper | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [x] wrapper zap lib, [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/uber-go/zap?label=zap%20latest%20go.md)](https://github.com/uber-go/zap)
- [x] log rolling package [![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/natefinch/lumberjack/v2.0?label=lumberjack%20v2%20go.mod)](https://github.com/natefinch/lumberjack)
- [x] load by config kit
  - [x] support load config by viper [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/spf13/viper?label=viper%20latest%20go.md)](https://github.com/spf13/viper)
- [x] support `zlog.LogsConfigFlavors` to init more flavors logger

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

### libs

- more libs see [go.mod](https://github.com/sinlov-go/zlog-zap-wrapper/blob/main/go.mod)

## usage

- more see [doc/LIB.md](doc/LIB.md)

## dev

- see [dev.md](doc-dev/dev.md)