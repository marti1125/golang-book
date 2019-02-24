package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("./chapter13/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
