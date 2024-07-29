package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(stdin, "11 22\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	//println("Err:", string(stderr.Bytes()))
	fmt.Println(out.String())
	println(out.String() == "33\n")

}
