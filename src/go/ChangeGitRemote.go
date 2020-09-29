package main

import (
	"runtime"
	"bytes"
	"os/exec"
	"fmt"
	"strings"
	"path/filepath"
	"container/list"
	"os"
	"bufio"
)

// 
func exec_shell(command string) (string, error) {
	//函数返回一个 *Cmd，用于使用给出的参数执行 name 指定的程序
	// 
	sysType := runtime.GOOS
	
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

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	
}

// 
func getFilelist(path string, l *list.List) {
	err := filepath.Walk(path, 
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() && strings.HasSuffix(path, ".git") {
				fmt.Printf("find :%s\n", path)
				l.PushBack(strings.Replace(path, ".git", "", -1 ))
			}

			return nil
		})

	if err != nil {
			fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func getInput() string {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}




func main(){
	fmt.Print("【必填】请输入 保存 gitlab 项目的路径，可以是工作空间，非项目本身（如：/Users/kail/IdeaProjects/_business）:")
	gitlab_root := getInput()
	fmt.Println(gitlab_root)
	if len(gitlab_root) <=0 {
		fmt.Println("【缺少必填】保存 gitlab 项目的路径")
		return
	}

	fmt.Print("【必填】请输入需要替换的 Git 远程地址（如： 172.16.2.2）: ")
	oldOrigin := getInput()
	fmt.Println(oldOrigin)
	if len(gitlab_root) <=0 {
		fmt.Println("【缺少必填】被替换的地址")
		return
	}

	fmt.Print("【必填】请输入 替换后的 Git 远程地址（如：gitlab.kail.cn ）: ")
	newOrigin := getInput()
	fmt.Println(newOrigin)
	if len(gitlab_root) <=0 {
		fmt.Println("【缺少必填】替换后的地址")
		return
	}


	gitpaths := list.List{}
	getFilelist(gitlab_root, &gitpaths)

	for x := gitpaths.Front(); x != nil; x = x.Next() {
		fmt.Println("======================================================")
		fmt.Println(x.Value)
		// fmt.Printf("output: %T\n",x.Value)
		os.Chdir(fmt.Sprintf("%s",x.Value))
		command := fmt.Sprintf("git remote get-url origin")
		// fmt.Println(command)

		origin, _ := exec_shell(command)
		fmt.Printf("origin: %s \n",origin )
		
		if strings.Contains(origin, oldOrigin) {
			originNew := strings.Replace(origin, oldOrigin, newOrigin, 1 )
			os.Chdir(fmt.Sprintf("%s",x.Value))
			command = fmt.Sprintf("git remote set-url origin %s", originNew)
			// fmt.Println(command)
			exec_shell(command)
			fmt.Printf("%s   %s  -->  %s \n", x.Value , origin , originNew)
		}	else {
			fmt.Printf("替换失败（未匹配）：%s   %s   被替换的地址:%s  替换后的地址:%s \n",x.Value, origin, oldOrigin ,newOrigin)
		}
	}	
}



// 获取命令执行结果

// 字符串替换

// 读取命令行参数

