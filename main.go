package main 

import (
	"fmt"
	"github.com/ahrtr/vfsgenDemo/data"
)

func getFileInfo(fileName string) {
	file, err := data.Assets.Open(fileName)

	if err != nil {
		fmt.Printf("Failed to open file, fileName:%s, err: %v\n", fileName, err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info, fileName: %s, err: %v\n", fileName, err)
		return
	}

	fmt.Printf("fileName: %s, File info: %v\n", fileName, fi)
}

func main() {
	getFileInfo("examplefile1.json")
	getFileInfo("subdir1/examplefile3.xml")
}
