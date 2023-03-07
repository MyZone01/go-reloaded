package main

import (
	"fmt"
	"goreloaded"
)

func main() {
	inputFileName, outputFileName := goreloaded.GetArgs()
	fmt.Println(outputFileName)
	fileContent := goreloaded.GetFile(inputFileName)
	result := fileContent

	cmdList := goreloaded.ParseCmd(fileContent)
	for _, cmd := range cmdList {
		result = goreloaded.Format(result, cmd)
	}
	for i := len(cmdList) - 1; i >= 0; i-- {
		cmd := cmdList[i]
		result = result[:cmd.CmdStart-1] + result[cmd.CmdEnd+1:]
	}
	result = goreloaded.Clean(result);
	goreloaded.CreateResultFile(outputFileName, result)
}
