package goreloaded

import (
	"fmt"
	"os"
)

func isVowel(char rune) bool {
	return char == 'a' || char == 'A' || char == 'e' || char == 'E' || char == 'o' || char == 'O' || char == 'u' || char == 'U' || char == 'i' || char == 'I'
}

func isHeadWord(char rune, str string, i int) bool {
	return i == 0 || str[i-1] == ' ' || str[i-1] == '(' || str[i-1] == ')' || isPunctuation(rune(str[i-1]))
}

func isPunctuation(char rune) bool {
	return char == ',' || char == '.' || char == '!' || char == '?' || char == ':' || char == ';' || char == '\''
}

// Format the input file content with specified rules
func Format(fileContent string) string {
	result := []rune{}
	buffer := []rune{}
	cmdFound := false
	cmd := ""
	// overideCmd := ""
	isMark := false
	numberOfWords := 1
	base := map[string]string{
		"hex": "0123456789ABCDEF",
		"bin": "01",
	}
	for i := len(fileContent) - 1; i >= 0; i-- {
		char := rune(fileContent[i])
		if char == '(' {
			if cmd == "" {
				cmd = string(buffer)
			}
			buffer = []rune{}
			if isHeadWord(char, fileContent, i) {
				i--
			}
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
				if isHeadWord(char, fileContent, i) {
					if (char == 'a' || char == 'A') && (i == len(fileContent)-1 || fileContent[i+1] == ' ') {
						if i != len(fileContent)-1 && isVowel(rune(fileContent[i+2])) {
							result = append([]rune{'n'}, result...)
						}
					}
					if cmd == "cap" && char >= 'a' && char <= 'z' {
						char -= 32
						result = append([]rune{char}, result...)
						numberOfWords--
						if numberOfWords == 0 {
							numberOfWords = 1
							cmd = ""
						}
						continue
					} else if cmd == "hex" || cmd == "bin" {
						buffer = append([]rune{char}, buffer...)
						number := ConvertBase(string(buffer), base[cmd], "0123456789")
						result = append([]rune(number), result...)
						buffer = []rune{}
						cmd = ""
						continue
					}
				}
				if char == ' ' || isPunctuation(char) {
					if cmd != "cap" {
						numberOfWords--
						if numberOfWords == 0 {
							numberOfWords = 1
							cmd = ""
						}
					}
					if isPunctuation(char) {
						if char == '\'' {
							if (i != len(fileContent)-1 && i >= 3 && fileContent[i-3:i+2] != "don't") || i == len(fileContent)-1 || i < 3 {
								if !isMark {
									if i != len(fileContent)-1 && fileContent[i+1] != ' ' {
										result = append([]rune{' '}, result...)
									} else if i != 0 && fileContent[i-1] == ' ' {
										i--
									}
								} else {
									if i != len(fileContent)-1 && fileContent[i+1] == ' ' {
										result = result[1:]
									} else if i != 0 && fileContent[i-1] != ' ' {
										result = append([]rune{' '}, result...)
									}
								}
								isMark = !isMark
							}
						} else {
							if i != len(fileContent)-1 && fileContent[i+1] != '.' && fileContent[i+1] != '?' {
								if fileContent[i+1] != ' ' && i != len(fileContent)-2 && !isPunctuation(rune(fileContent[i+2])) {
									result = append([]rune{' '}, result...)
								}
							}
							if i != 0 && fileContent[i-1] == ' ' {
								i--
							}
						}
					}
				} else {
					if cmd == "hex" || cmd == "bin" {
						buffer = append([]rune{char}, buffer...)
						continue
					} else if cmd == "up" && char >= 'a' && char <= 'z' {
						char -= 32
					} else if cmd == "low" && char >= 'A' && char <= 'Z' {
						char += 32
					}
				}
				result = append([]rune{char}, result...)
			}
		}
	}
	return string(result)
}

// Launch the go reloaded process
func Run() {
	if len(os.Args) == 3 {
		inputFileName, outputFileName := os.Args[1], os.Args[2]
		fileContent, err := os.ReadFile(inputFileName)
		if err != nil {
			fmt.Println("Input file not found")
			os.Exit(1)
		}
		result := Format(string(fileContent))
		file, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println("Output file creation error")
		}
		file.WriteString(result)
		file.Close()
	} else {
		fmt.Println("You should have 3 arguments")
		os.Exit(1)
	}
}
