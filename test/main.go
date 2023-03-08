package main

import (
	"fmt"
	"goreloaded"
)

func main() {
	inputFileName, outputFileName := goreloaded.GetArgs()
	fmt.Println(outputFileName)
	fileContent := goreloaded.GetFile(inputFileName)
	result := goreloaded.Format(fileContent)
	goreloaded.CreateResultFile(outputFileName, result)
}
