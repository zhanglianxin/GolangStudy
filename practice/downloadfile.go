package main

import (
	"os"
	"path"
	"net/http"
	"io"
	"bufio"
)

func main() {
	targetPath := "xzx/"

	// https://gist.github.com/mattes/d13e273314c3b3ade33f
	if _, err := os.Stat(targetPath); err != nil {
		panic(err)
	}

	lines := readLines("/home/zhanglianxin/a")

	for _, file := range lines {
		if err := downloadFile(file, targetPath); err != nil {
			panic(err)
		}
	}

}

// https://gist.github.com/kendellfab/7417164
func readLines(filePath string) (lines []string) {
	inFile, _ := os.Open(filePath)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

// https://golangcode.com/download-a-file-from-a-url/
func downloadFile(url string, filePath string) error {
	if "" == filePath {
		filePath = "./"
	}
	// https://www.dotnetperls.com/path-go
	_, file := path.Split(url)

	out, err := os.Create(filePath + file)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}
