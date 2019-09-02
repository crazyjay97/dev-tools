package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var (
		AppPath string
		err     error
	)
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	fmt.Println(err)
	fmt.Println(AppPath)
	dir, err := os.Getwd()
	fmt.Println(dir)
	getFileInProject("configs/config.json")
}

func getFileInProject(fileName string) []byte {
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	var runPath string
	runPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	if bytes, err := ioutil.ReadFile(AppPath + string(os.PathSeparator) + fileName); err != nil {
		return bytes
	}
	if bytes, err := ioutil.ReadFile(runPath + string(os.PathSeparator) + fileName); err != nil {
		return bytes
	}
	panic("Config Initialize Error")
}
