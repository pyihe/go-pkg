| package  |  comment | 
| :---- | :----|
|[bytes](https://github.com/pyihe/go-pkg/tree/master/bytes)|byte相关|
|[certs](https://github.com/pyihe/go-pkg/tree/master/certs)|证书相关|
|[clone](https://github.com/pyihe/go-pkg/tree/master/clone)|拷贝相关|
|[encoding](https://github.com/pyihe/go-pkg/tree/master/encoding)|编解码相关|
|[encrypts](https://github.com/pyihe/go-pkg/tree/master/encrypts)|加解密相关|
|[errors](https://github.com/pyihe/go-pkg/tree/master/errors)|错误相关|
|[files](https://github.com/pyihe/go-pkg/tree/master/files)|文件相关|
|[https](https://github.com/pyihe/go-pkg/tree/master/https)|http相关|
|[list](https://github.com/pyihe/go-pkg/tree/master/list)|队列相关|
|[logs](https://github.com/pyihe/go-pkg/tree/master/logs)|日志相关|
|[maps](https://github.com/pyihe/go-pkg/tree/master/maps)|map相关|
|[maths](https://github.com/pyihe/go-pkg/tree/master/maths)|数学相关|
|[monitor](https://github.com/pyihe/go-pkg/tree/master/monitor)|监控相关|
|[nets](https://github.com/pyihe/go-pkg/tree/master/nets)|网络相关|
|[prt](https://github.com/pyihe/go-pkg/tree/master/ptr)|基础数据类型指针|
|[rands](https://github.com/pyihe/go-pkg/tree/master/rands)|随机函数相关|
|[redis](https://github.com/pyihe/go-pkg/tree/master/redis)|redis相关|
|[snowflakes](https://github.com/pyihe/go-pkg/tree/master/snowflakes)|snowflake相关|
|[sorts](https://github.com/pyihe/go-pkg/tree/master/sorts)|排序相关|
|[strings](https://github.com/pyihe/go-pkg/tree/master/strings)|字符串相关|
|[sync](https://github.com/pyihe/go-pkg/tree/master/syncs)|同步锁相关|
|[times](https://github.com/pyihe/go-pkg/tree/master/times)|定时器相关|
|[go-pkgs](https://github.com/pyihe/go-pkg/tree/master/go-pkgs)| 一些辅助函数|
|[zips](https://github.com/pyihe/go-pkg/tree/master/zips)|压缩相关| 

- golang编译命令

1. GOOS="target OS" GOARCH="target arch" go build -o "output file name"

   |   OS   | GOOS         |       GOARCH        |
   |:------|:------      | :---------------   |
   |Mac|darwin|386, amd64, arm, arm64|
   |DragonflyBSD|dragonfly|amd64|
   |FreeBSD|freebsd|386, amd64, arm|
   |Debian,RedHat,CentOs,Ubuntu|linux|386, amd64, arm, arm64, ppc64, ppc641e|
   |NetBSD|netbsd|386, amd64, arm|
   |OpenBSD|openbsd|386, amd64, arm|
   |Plan 9|plan9|386, amd64|
   |Solaris|solaris|amd64|
   |Win series|windows|386, amd64|

2. [gox](https://github.com/mitchellh/gox) -osarch="target os/target arch"<br>

   |target os|target arch|
   |:--------|:----------|
   |darwin|386|
   |darwin|amd64|
   |linux|386|
   |linux|amd64|
   |linux|arm|
   |freebsd|386|
   |freebsd|amd64|
   |freebsd|arm|
   |openbsd|386|
   |openbsd|amd64|
   |netbsd|386|
   |netbsd|amd64|
   |netbsd|arm|
   |plan9|386|
   |windows|386|
   |windows|amd64|