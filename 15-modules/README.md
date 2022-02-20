#### GO MODULES

##### 1.go mod命令

| 命令            | 作用                           |
| --------------- | ------------------------------ |
| go mod init     | 生成mod.go 文件                |
| go mod download | 下载go.mod文件中指明的所有依赖 |
| go mod tidy     | 整理现有的依赖                 |
| go mod graph    | 查看现有的依赖结构             |
| go mod edit     | 编辑go.mod文件                 |
| go mod vendor   | 导出项目所有依赖到vendor目录   |
| go mod verify   | 校验一个模块是否被篡改过       |
| go mod why      | 查看为什么需要依赖某模块       |

##### 2.go mod环境变量

可以通过go env命令来进行查看
```shell
$ go env

GO111MODULE="on"
GOARCH="amd64"
GOBIN="/Users/lemoba/ranen/workspace/go/bin"
GOCACHE="/Users/lemoba/Library/Caches/go-build"
GOENV="/Users/lemoba/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/lemoba/ranen/workspace/go/pkg/mod"
GONOPROXY=""
...
```
#### GO111MODULE

go语言提供了GO111MODULE这个环境变量来作为Go modules的开关,其允许设置一下参数:
* auto: 只要项目包含了go.mod文件的话启用Go modules, 目前在Go11.1~Go1.14中仍然是默认值.
* on: 启用Go modules, 推荐设置,将会是未来版本中的默认值
* off: 禁用Go moudules, 不推荐设置.

可以通过设置
```shell
$ go env -w GO111MODULE=on
```
#### GOPROXY
这个环境变量主要是设置Go模块代理,其作用是用于使用Go在后续拉取模块版本时直接通过镜像站点来快速拉取.
GOPROXY的默认值:https://proxy.golang.org,direct
proxy.golang.org国内访问不了,需要设置国内的代理
> direct 用于指示Go回源到模块版本的源地址去抓取
* 阿里云
    https://mirrors.aliyun.com/goproxy/
* 七牛云
    https://goproxy.cn,direct
#### GOSUMDB
用来校验拉取的第三方库是否是完整的
#### GOPRIVATE
表示是私有仓库,不会进行GOPROXY下载和校验