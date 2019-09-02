package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetFileInProject(fileName string) ([]byte, error) {
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var runPath string
	runPath, err = os.Getwd()
	var bytes []byte
	if bytes, err = ioutil.ReadFile(AppPath + string(os.PathSeparator) + fileName); err == nil {
		return bytes, nil
	}
	if bytes, err = ioutil.ReadFile(runPath + string(os.PathSeparator) + fileName); err == nil {
		return bytes, nil
	}
	return nil, err
}

func GetDirInProject(dir string) ([]os.FileInfo, error) {
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var runPath string
	runPath, err = os.Getwd()
	if infos, err := ioutil.ReadDir(AppPath + string(os.PathSeparator) + dir); err == nil {
		return infos, nil
	}
	if infos, err := ioutil.ReadDir(runPath + string(os.PathSeparator) + dir); err == nil {
		return infos, nil
	}
	return nil, err
}

func GetFileAbsPathInProject(fileName string) (string, error) {
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var runPath string
	runPath, err = os.Getwd()
	if _, err := ioutil.ReadFile(AppPath + string(os.PathSeparator) + fileName); err == nil {
		return AppPath + string(os.PathSeparator) + fileName, nil
	}
	if _, err := ioutil.ReadDir(runPath + string(os.PathSeparator) + fileName); err == nil {
		return runPath + string(os.PathSeparator) + fileName, nil
	}
	return "", err
}

func GetDirAbsPathInProject(dir string) (string, error) {
	AppPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var runPath string
	runPath, err = os.Getwd()
	if _, err := ioutil.ReadDir(AppPath + string(os.PathSeparator) + dir); err == nil {
		return AppPath + string(os.PathSeparator) + dir, nil
	}
	if _, err := ioutil.ReadDir(runPath + string(os.PathSeparator) + dir); err == nil {
		return runPath + string(os.PathSeparator) + dir, nil
	}
	return "", err
}
