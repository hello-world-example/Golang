package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//
func execShell(command string) (string, error) {
	sysType := runtime.GOOS

	/* Windows */
	if strings.Contains(sysType, "windows") {
		var out bytes.Buffer
		cmd := exec.Command("cmd", "/C", command)
		cmd.Stdout = &out
		err := cmd.Run()
		return out.String(), err
	} else {
		var out bytes.Buffer
		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = &out
		err := cmd.Run()
		return out.String(), err
	}
}

/* 获取 Git 路径 */
func findGitPath(path string, l *list.List) {
	err := filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			// 文件是路径，有 .git 文件
			if f.IsDir() && strings.HasSuffix(path, ".git") {
				fmt.Printf("find :%s\n", path)
				l.PushBack(strings.Replace(path, ".git", "", -1))
			}

			return nil
		})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func getInput() string {
	// 使用 os.Stdin 开启输入流
	// 函数原型 func NewReader(rd io.Reader) *Reader
	// NewReader创建一个具有默认大小缓冲、从 r 读取的 *Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)

	// in.ReadLine 函数具有三个返回值 []byte bool error
	// 分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	// 范围输入流
	return string(str)
}

/* 解析请求参数 */
func parseInput() (string, string, string) {
	fmt.Print("【必填】请输入 Git 项目的工作空间，非项目本身（如：/Users/kail/IdeaProjects/）:")
	gitlab_root := getInput()
	if len(gitlab_root) <= 0 {
		fmt.Println("【缺少必填】Git 项目的工作空间")
		return "", "", ""
	}
	fmt.Println(gitlab_root)

	fmt.Print("【选填，默认：172.16.2.2】请输入 需要替换的 Git 远程地址 : ")
	oldOrigin := getInput()
	if len(oldOrigin) <= 0 {
		oldOrigin = "172.16.2.2"
	}
	fmt.Println(oldOrigin)

	fmt.Print("【选填，默认：gitlab.kail.int】请输入 新的 Git 远程地址: ")
	newOrigin := getInput()
	if len(newOrigin) <= 0 {
		newOrigin = "gitlab.kail.int"
	}
	fmt.Println(newOrigin)

	return gitlab_root, oldOrigin, newOrigin
}

func main() {

	git_root, oldOrigin, newOrigin := parseInput()
	if len(git_root) <= 0 {
		return
	}

	// 找到的 Git 路径
	gitpaths := list.List{}
	findGitPath(git_root, &gitpaths)

	for x := gitpaths.Front(); x != nil; x = x.Next() {
		gitpath := x.Value

		fmt.Println("====================================================== ", gitpath)

		// 切换到 gitpath
		os.Chdir(fmt.Sprintf("%s", gitpath))
		origin, _ := execShell(fmt.Sprintf("git remote get-url origin"))
		origin = strings.Trim(strings.Trim(origin, "\r\n"), "\n")

		// 如果匹配到
		if strings.Contains(origin, oldOrigin) {
			// 替换为信息的地址
			originNew := strings.Replace(origin, oldOrigin, newOrigin, 1)

			// 切换到 gitpath
			os.Chdir(fmt.Sprintf("%s", gitpath))
			// 执行替换命令
			execShell(fmt.Sprintf("git remote set-url origin %s", originNew))

			// 打印结果
			fmt.Printf("Success: %s   %s  -->  %s \n", gitpath, origin, originNew)
		} else {

			// 打印失败原因
			fmt.Printf("Error: 替换失败：项目 %s 的 origin %s 未匹配到 %s \n", gitpath, origin, oldOrigin)
		}
	}
}

/*
go build -o ChangeGitRemote.mac ChangeGitRemote.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ChangeGitRemote.linux ChangeGitRemote.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ChangeGitRemote.exe ChangeGitRemote.go
*/
