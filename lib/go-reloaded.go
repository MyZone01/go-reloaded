package goreloaded

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Command struct {
	name          string
	numberOfWords int
	skippedWords  int
	overlap      *Command
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'A' || char == 'e' || char == 'E' || char == 'o' || char == 'O' || char == 'u' || char == 'U' || char == 'i' || char == 'I'
}

func isHeadWord(char rune, str string, i int) bool {
	return i == 0 || !isAlphaNumeric(rune(str[i-1]))
}

func isPunctuation(char rune) bool {
	return char == ',' || char == '.' || char == '!' || char == '?' || char == ':' || char == ';' || char == '\''
}

func isAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

// Format the input file content with specified rules
func Format(fileContent string) string {
	result := []rune{}
	buffer := []rune{}
	cmdFound := false
	cmd := Command{
		name:          "",
		numberOfWords: 1,
		skippedWords:  0,
	}
	isMark := false
	base := map[string]string{
		"hex": "0123456789ABCDEF",
		"dec": "0123456789",
		"bin": "01",
	}
	cmdRegEx := regexp.MustCompile(`\((cap|low|up|bin|hex)(\,\s\d+\)|\))`)
	for i := len(fileContent) - 1; i >= 0; i-- {
		char := rune(fileContent[i])
		if char == '(' {
			if cmd.name == "" {
				buffer = append([]rune{char}, buffer...)
				_cmd := string(buffer)
				isMatch := cmdRegEx.MatchString(_cmd)
				if isMatch {
					_cmd = _cmd[1 : len(_cmd)-1]
					cmdPart := strings.Split(_cmd, ", ")
					cmd.name = cmdPart[0]
					if len(cmdPart) == 2 {
						number := cmdPart[1]
						cmd.numberOfWords = AtoiBase(string(number), base["dec"])
					}
					if i > 0 && fileContent[i-1] == ' ' {
						i--
					}
				} else {
					result = append(buffer, result...)
				}
			} else if len(buffer) == 0 {
				buffer = append([]rune{char}, buffer...)
				_cmd := string(buffer)
				isMatch := cmdRegEx.MatchString(_cmd)
				if isMatch {
					_cmd = _cmd[1 : len(_cmd)-1]
					cmdPart := strings.Split(_cmd, ", ")
					cmd.name = cmdPart[0]
					if len(cmdPart) == 2 {
						number := cmdPart[1]
						cmd.numberOfWords = AtoiBase(string(number), base["dec"])
					}
					if i > 0 && fileContent[i-1] == ' ' {
						i--
					}
				} else {
					result = append([]rune{char}, result...)
				}
			}
			buffer = []rune{}
			cmdFound = false
		} else if char == ')' && !cmdFound {
			cmdFound = true
			buffer = append([]rune{char}, buffer...)
		} else {
			if cmdFound {
				if char == ')' {
					buffer = append([]rune{char}, []rune{}...)
					result = append(buffer, result...)
					cmdFound = true
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
					if cmd.name == "cap" {
						if char >= 'a' && char <= 'z' {
							char -= 32
						}
						result = append([]rune{char}, result...)
						cmd.numberOfWords--
						if cmd.numberOfWords == 0 {
							cmd.numberOfWords = 1
							cmd.name = ""
						}
						continue
					} else if cmd.name == "hex" || cmd.name == "bin" {
						if !isPunctuation(char) {
							buffer = append([]rune{char}, buffer...)
							number := ConvertBase(string(buffer), base[cmd.name], "0123456789")
							result = append([]rune(number), result...)
							buffer = []rune{}
							cmd.name = ""
							continue
						}
					}
				}
				if char == ' ' || isPunctuation(char) {
					if cmd.name != "cap" {
						cmd.numberOfWords--
						if cmd.numberOfWords == 0 {
							cmd.numberOfWords = 1
							cmd.name = ""
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
					if cmd.name == "hex" || cmd.name == "bin" {
						buffer = append([]rune{char}, buffer...)
						continue
					} else if cmd.name == "up" && char >= 'a' && char <= 'z' {
						char -= 32
					} else if cmd.name == "low" && char >= 'A' && char <= 'Z' {
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
