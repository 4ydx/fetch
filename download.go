package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func directory(url string) string {
	splitOn := "://"
	index := strings.LastIndex(url, splitOn)
	return fmt.Sprintf("%s.html", url[index+len(splitOn):])
}

func download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	targetDir := directory(url)
	if err := os.Mkdir(targetDir, 0777); err != nil {
		return err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%s/index.html", targetDir), body, 0666); err != nil {
		return err
	}
	return nil
}
