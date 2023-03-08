package goreloaded

import (
	"fmt"
	"os"
)

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

func GetArgs() (string, string) {
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2]
	}
	return "", ""
}

func getFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Not found")
		return ""
	}
	return string(content)
}

func Format(fileContent string) string {
	runes := []rune{}
	buffer := []rune{}
	cmdFound := false
	cmd := ""
	isMark := false
	numberOfWords := 1
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
					if char == ' ' {
						numberOfWords--
						if cmd == "hex" {
							number := " " + ConvertBase(string(buffer), "0123456789ABCDEF", "0123456789")
							runes = append([]rune(number), runes...)
							buffer = []rune{}
							cmd = ""
							continue
						} else if cmd == "bin" {
							number := " " + ConvertBase(string(buffer), "01", "0123456789")
							runes = append([]rune(number), runes...)
							buffer = []rune{}
							cmd = ""
							continue
						}
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

func createResultFile(fileName string, fileContent string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Creation Error")
	}
	file.WriteString(fileContent)
	file.Close()
}

func Run() {
	inputFileName, outputFileName := GetArgs()
	fileContent := getFile(inputFileName)
	result := Format(fileContent)
	createResultFile(outputFileName, result)
}
