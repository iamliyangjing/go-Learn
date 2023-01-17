package main

import (
	"fmt"
	"os"
	"os/exec"
)

//进程信息

func main() {
	// go run example/main,go a b c d
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH")) // /usr/local/go/bin
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("gerp", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
