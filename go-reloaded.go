package main

import (
	"fmt"
	"os"
)

func toUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			runes[i] = rune(runes[i] - 32)
		}
	}
	return string(runes)
}

func isVowel(char rune) bool {
	return char == 'a' ||
		char == 'A' ||
		char == 'e' ||
		char == 'E' ||
		char == 'o' ||
		char == 'O' ||
		char == 'u' ||
		char == 'U' ||
		char == 'i' ||
		char == 'I'
}

func isPunctuation(char rune) bool {
	return char == ',' || char == '.' || char == '!' || char == '?' || char == ':' || char == ';' || char == '\''
}

// Get execution's argument. Return the input and output file name
func getArgs() (string, string, string) {
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2], ""
	}
	return "", "", "Invalid number of argument"
}

// Get the input file content as a string
func getInputFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Input file not found")
		return ""
	}
	return string(content)
}

// Format the input file content with specified rules
func Format(fileContent string) string {
	runes := []rune{}
	buffer := []rune{}
	cmdFound := false
	cmd := ""
	isMark := false
	numberOfWords := 1
	base := map[string]string{
		"hex": "0123456789ABCDEF",
		"bin": "01",
	}
	for i := len(fileContent) - 1; i >= 0; i-- {
		char := rune(fileContent[i])
		if char == '(' {
			cmd = string(buffer)
			buffer = []rune{}
			i--
			cmdFound = false
		} else if char == ')' {
			cmdFound = true
		} else {
			if cmdFound {
				if char == ',' {
					numberOfWords = AtoiBase(string(buffer[1:]), "0123456789")
					buffer = []rune{}
				} else {
					buffer = append([]rune{char}, buffer...)
				}
			} else {
				if cmd != "" {
					if char == ' ' || i == 0 {
						if cmd == "hex" || cmd == "bin" {
							if i == 0 {
								buffer = append([]rune{char}, buffer...)
							}
							number := ConvertBase(toUpper(string(buffer)), base[cmd], "0123456789")
							if char == ' ' {
								number = " " + number
							}
							runes = append([]rune(number), runes...)
							buffer = []rune{}
							cmd = ""
							continue
						}
					}
					if char == ' ' {
						numberOfWords--
						if numberOfWords == 0 {
							numberOfWords = 1
							cmd = ""
						}
					} else {
						if cmd == "hex" {
							buffer = append([]rune{char}, buffer...)
							continue
						} else if cmd == "bin" {
							buffer = append([]rune{char}, buffer...)
							continue
						} else if cmd == "up" {
							if char >= 'a' && char <= 'z' {
								char -= 32
							}
						} else if cmd == "low" {
							if char >= 'A' && char <= 'Z' {
								char += 32
							}
						} else if cmd == "cap" && (i == 0 || fileContent[i-1] == ' ') {
							if char >= 'a' && char <= 'z' {
								char -= 32
							}
						}
					}
				}
				if isPunctuation(char) {
					if char == '\'' {
						if (i != len(fileContent)-1 && i >= 3 && fileContent[i-3:i+2] != "don't") || i == len(fileContent)-1 || i < 3 {
							if !isMark {
								if i != len(fileContent)-1 && fileContent[i+1] != ' ' {
									runes = append([]rune{' '}, runes...)
								} else if i != 0 && fileContent[i-1] == ' ' {
									i--
								}
							} else {
								if i != len(fileContent)-1 && fileContent[i+1] == ' ' {
									runes = runes[1:]
								} else if i != 0 && fileContent[i-1] != ' ' {
									runes = append([]rune{' '}, runes...)
								}
							}
							isMark = !isMark
						}
					} else {
						if i != len(fileContent)-1 && fileContent[i+1] != '.' && fileContent[i+1] != '?' && fileContent[i+1] != ' ' {
							runes = append([]rune{' '}, runes...)
						}
						if i != 0 && fileContent[i-1] == ' ' {
							i--
						}
					}
				} else if (char == 'a' || char == 'A') && (i == 0 || fileContent[i-1] == ' ') && fileContent[i+1] == ' ' {
					if isVowel(rune(fileContent[i+2])) {
						runes = append([]rune{'n'}, runes...)
					}
				}
				runes = append([]rune{char}, runes...)
			}
		}
	}
	return string(runes)
}

// Create the output File with formated content
func createOutputFile(fileName string, fileContent string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Output file creation error")
	}
	file.WriteString(fileContent)
	file.Close()
}

// Launch the go reloaded process
func Run() {
	inputFileName, outputFileName, error := getArgs()
	if error == "" {
		fileContent := getInputFile(inputFileName)
		result := Format(fileContent)
		createOutputFile(outputFileName, result)
	} else {
		fmt.Println(error)
	}
}
