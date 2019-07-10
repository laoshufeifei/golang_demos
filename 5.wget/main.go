package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func checkFatalError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error is %v\n", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("hello wget")

	url := "https://www.baidu.com/img/bd_logo1.png"
	// url := "https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/releases/18.04/ubuntu-18.04.2-live-server-amd64.iso"

	response, err := http.Get(url)
	checkFatalError(err)
	defer response.Body.Close()

	// print header
	fmt.Printf("http header for %s is:\n", url)
	for k, v := range response.Header {
		fmt.Printf("%s: ", k)
		for i := range v {
			fmt.Printf("%s", v[i])
		}
		fmt.Printf("\n")
	}

	// write body to file
	fmt.Println("begin write body")
	baseName := filepath.Base(url)
	newFile, err := os.Create(baseName)
	checkFatalError(err)
	defer newFile.Close()

	io.Copy(newFile, response.Body)
	fmt.Println("finish save body to file:", baseName)
}
