# 执行外部命令



## 入门示例

```go
package main

import (
	"os"
	"os/exec"
)

func main() {
  // 执行命令
	cmd := exec.Command("ls", "-al")
  // 指定标准输出
	cmd.Stdout = os.Stdout
  // 运行
	cmd.Run()
}
```



## 获取命令执行结果

```go
package main

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

func execCommand(command string) (string, error) {
	var cmd *exec.Cmd

  // 区分环境
	if strings.Contains(runtime.GOOS, "windows") {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	var result bytes.Buffer
	cmd.Stdout = &result
	error := cmd.Run()
	return result.String(), error
}

func main() {
	result, _ := execCommand("ls -al")
	print(result)
}

```





## Read More

- [import "os/exec"](https://studygolang.com/pkgdoc)