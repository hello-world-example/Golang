# 全平台编译



## MAC

```bash
# Linux 可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main.linux main.go

# Windows 可执行程序
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```



## Linux

```bash
# MAC 可执行程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o main.mac main.go

# Windows 可执行程序
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```



## Windows

```bash
# MAC 可执行程序
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o main.mac main.go

# Linux 可执行程序
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o main.linux main.go
```



## 变量设置

- `GOOS`：目标平台的操作系统（darwin、freebsd、linux、windows）
- `GOARCH`：目标平台的体系架构（386、amd64、arm）
- 交叉编译不支持 CGO 所以要禁用它 `CGO_ENABLED = 0`



## 当前系统变量

```bash
$ go env

GO111MODULE=""
# ❤
GOARCH="amd64"
GOBIN=""
...
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GONOPROXY=""
GONOSUMDB=""
# ❤
GOOS="darwin"
...
```

