package goreloaded

import (
	"fmt"
	"os"
)

type CMD struct {
	Type      string
	WordStart int
	WordEnd   int
	CmdStart  int
	CmdEnd    int
}

func GetArgs() (string, string) {
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2]
	}
	return "", ""
}

func GetFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Not found")
		return ""
	}
	return string(content)
}

func ParseCmd(fileContent string) []CMD {
	runesBuffer := []rune{}
	wordStartIndexes := []int{0}
	currentCmd := CMD{}
	cmd := []CMD{}
	coefficientFound := false
	cmdFound := false
	numberOfWords := 1
	for i, char := range fileContent {
		if char == '(' {
			cmdFound = true
			currentCmd.WordEnd = i - 1 // i - 1
			currentCmd.CmdStart = i
		} else if char == ')' {
			cmdFound = false
			if coefficientFound {
				numberOfWords = AtoiBase(string(runesBuffer[1:]), "0123456789")
			}
			if currentCmd.Type == "hex" ||
				currentCmd.Type == "bin" ||
				currentCmd.Type == "up" ||
				currentCmd.Type == "low" ||
				currentCmd.Type == "cap" {
				currentCmd.CmdEnd = i
				currentCmd.WordStart = wordStartIndexes[len(wordStartIndexes)-numberOfWords]
				cmd = append(cmd, currentCmd)
				numberOfWords = 1
				runesBuffer = []rune{}
			}
		} else if char == ',' && cmdFound {
			coefficientFound = true
		} else if (char == ' ' || char == ',' || char == '.' || char == '!' || char == '?' || char == ':' || char == ';') && !cmdFound && i != len(fileContent)-1 && isAlpha(rune(fileContent[i+1])) {
			wordStartIndexes = append(wordStartIndexes, i)
		} else if coefficientFound && isNumeric(char) {
			runesBuffer = append(runesBuffer, char)
		} else if cmdFound {
			runesBuffer = append(runesBuffer, char)
			if string(runesBuffer) == "hex" ||
				string(runesBuffer) == "bin" ||
				string(runesBuffer) == "up" ||
				string(runesBuffer) == "low" ||
				string(runesBuffer) == "cap" {
				currentCmd.Type = string(runesBuffer)
				runesBuffer = []rune{}
			}
		}
	}
	return cmd
}

func Format(fileContent string, cmd CMD) string {
	runes := []rune(fileContent)
	if cmd.Type == "hex" {
		runes = hex(runes, cmd.WordStart, cmd.WordEnd)
	} else if cmd.Type == "bin" {
		runes = bin(runes, cmd.WordStart, cmd.WordEnd)
	} else if cmd.Type == "up" {
		runes = up(runes, cmd.WordStart, cmd.WordEnd)
	} else if cmd.Type == "low" {
		runes = low(runes, cmd.WordStart, cmd.WordEnd)
	} else if cmd.Type == "cap" {
		runes = cap(runes, cmd.WordStart, cmd.WordEnd)
	}
	return string(runes)
}

func CreateResultFile(fileName string, fileContent string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Creation Error")
	}
	file.WriteString(fileContent)
	file.Close()
}

func Clean(fileContent string) string {
	runes := []rune(fileContent)
	// forbiddenPunctuation := []string{
	// 	" ...",
	// 	" !?",
	// 	" ,",
	// 	" .",
	// 	" !",
	// 	" ?",
	// 	" :",
	// 	" ;",
	// }
	i := 0
	for {
		if i == len(runes) {
			break
		}
		char := runes[i]
		// str := string(char)
		// fmt.Println(str)
		if char == ',' || char == '.' || char == '!' || char == '?' || char == ':' || char == ';' {
			if runes[i-1] == ' ' {
				runes = append(runes[:i-1], runes[i:]...)
				i--
			}
			if i < len(runes)-1 && runes[i+1] != ' ' {
				if string(runes[i:i+4]) == "..." && i+3 <= len(runes)-1 {
					runes = append(runes[:i+4], append([]rune{' '}, runes[i+4:]...)...)
					i += 4
					continue
				} else if string(runes[i:i+3]) == "!?" && i+2 <= len(runes)-1 {
					runes = append(runes[:i+3], append([]rune{' '}, runes[i+3:]...)...)
					i += 3
					continue
				} else {
					runes = append(runes[:i+1], append([]rune{' '}, runes[i+1:]...)...)
					i++
					continue
				}
			}
		}
		i++
	}
	return string(runes)
}

func isLetter(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func isNumeric(char rune) bool {
	return char >= '0' && char <= '9'
}

func isLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func isUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func cap(runes []rune, start, end int) []rune {
	first := true
	for i := start; i < end; i++ {
		if isLetter(runes[i]) && first {
			if isLower(runes[i]) {
				runes[i] -= 32
			}
			first = false
		} else if isUpper(runes[i]) {
			runes[i] += 32
		} else if !isLetter(runes[i]) {
			first = true
		}
	}
	return runes
}

func low(runes []rune, start, end int) []rune {
	for i := start; i < end; i++ {
		if isUpper(runes[i]) {
			runes[i] = rune(runes[i] + 32)
		}
	}
	return runes
}

func up(runes []rune, start, end int) []rune {
	for i := start; i < end; i++ {
		if isLower(runes[i]) {
			runes[i] = rune(runes[i] - 32)
		}
	}
	return runes
}

func bin(runes []rune, start, end int) []rune {
	number := ConvertBase(string(runes[start+1:end-1]), "01", "0123456789")
	return []rune(string(runes[:start+1]) + number + string(runes[end-1:]))
}

func hex(runes []rune, start, end int) []rune {
	number := ConvertBase(string(runes[start+1:end-1]), "0123456789ABCDEF", "0123456789")
	return []rune(string(runes[:start+1]) + number + string(runes[end-1:]))
}
