package main

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

func execCommand(command string) (string, string, error) {

	var cmd *exec.Cmd

	if strings.Contains(runtime.GOOS, "windows") {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("/bin/bash", "-c", command)
	}

	var result bytes.Buffer
	var fail bytes.Buffer

	cmd.Stdout = &result
	cmd.Stderr = &fail

	error := cmd.Run()
	return result.String(), fail.String(), error
}

func main() {
	result, fail, error := execCommand("lll")
	println(result)
	println(fail)
	println(error.Error())
}
