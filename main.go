package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func PrintRepreatFile(path string, fileNameSizeMap map[string]int64, exFileList []string) {

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fs {

		fmt.Println(file.Name())
		if file.IsDir() {

			PrintRepreatFile(path+"/"+file.Name(), fileNameSizeMap, exFileList)

		} else {

			if file.Size() > 1000000 {
				fileSize := fileNameSizeMap[file.Name()]

				if fileSize == file.Size() {
					fmt.Println(path + "/" + file.Name())

					exFileList = append(exFileList, path+file.Name())
				} else {

					fileNameSizeMap[file.Name()] = file.Size()
				}
			}

		}

	}
}

func main() {
	var path string
	fmt.Println("Введите путь")
	fmt.Scanf("%s\n", &path)
	fileNameSizeMap := make(map[string]int64, 10000)
	exFileList := make([]string, 100, 1000)
	PrintRepreatFile(path, fileNameSizeMap, exFileList)
	fmt.Println(fileNameSizeMap, exFileList)

}
