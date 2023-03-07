#   Go reloaded

##  Tasks
+   [ ]  Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")
+   [ ]  Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")
+   [x]  Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")
+   [x]  Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")
+   [x]  Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
    +   [x]  For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")
+   [ ]  Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
    +   [ ]  Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".
+   [ ]  The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
    +   [ ]  If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")
+   [ ]  Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!")

##  Algorithm
### Main Function
```go
package main

import "goreloaded"

func main() {
    inputFileName, outputFileName := goreloaded.GetArgs();
    fileContent := goreloaded.GetFile(inputFileName);
    result := fileContent;

    cmdList := goreloaded.ParseCmd(fileContent);

    for i, cmd ::= range cmdList {
        result = goreloaded.Format(result, cmd);
    }

    result := goreloaded.CreateResultFile(outputFileName);
}
```

### Get Args
```go
package goreloaded

import "os"

func GetArgs(): (string, string) {
    if len(os.Args) == 3 {
        return (os.Args[1], os.Args[2])
    }
}
```

### Get File
```go
package goreloaded

import (
    "fmt"
    "os"
)

func GetFile(file string) string {
    content, err = os.ReadFile(file)
    if err != nil {
        fmt.Println("Not found")
        return
    }
}
```

### Parse Cmd
```go
package goreloaded

import (
    "fmt"
    "os"
)

type CMD struct {
	Type     string
	Start    int
	End      int
}

func ParseCmd(fileContent string) Cmd[] {
    runesBuffer := []rune{}
    wordStartIndexes := []int{0}
    currentCmd := Cmd{}
    cmd := []Cmd
    coefficientFound := false
    cmdFound := false
    numberOfWords := 1
    for i, char := range fileContent {
        if char == '(' {
            cmdFound = true
            currentCmd.End = i
        } else if char == ')' {
            cmdFound = false
            if string(runesBuffer) == 'hex' ||
               string(runesBuffer) == 'bin' ||
               string(runesBuffer) == 'up' ||
               string(runesBuffer) == 'low' ||
               string(runesBuffer) == 'cap' {
               currentCmd.End = len(wordStartIndexes) - numberOfWords
               cmd = append(cmd, string(runesBuffer))
               numberOfWords = 1
               runesBuffer = []rune{}
            }
        } else if char == ',' && cmdFound {
            coefficientFound = true
        } else if char == ' ' && i != len(fileContent) - 1 && IsAlpha(fileContent[i+1]) {
            wordStartIndexes := append(wordStartIndexes, i)
        } else if coefficientFound && IsNumeric(char) {
            runesBuffer = append(runesBuffer, char)
        } else {
            runesBuffer = append(runesBuffer, char)
            if string(runesBuffer) == 'hex' ||
               string(runesBuffer) == 'bin' ||
               string(runesBuffer) == 'up' ||
               string(runesBuffer) == 'low' ||
               string(runesBuffer) == 'cap' {
                currentCmd.Type = string(runesBuffer)
                runesBuffer = []rune{}
            }
        }
    }
    return cmd
}
```

### Format
```go
package goreloaded

func Format(fileContent string, cmd CMD) string {
    runes := []rune(fileContent)
    for i, char := range fileContent {
        if cmd.Type == 'hex' {
            runes = hex(runes, cmd.Start, cmd.End)
        } else if cmd.Type == 'bin' {
            runes = bin(runes, cmd.Start, cmd.End)
        } else if cmd.Type == 'up' {
            runes = up(runes, cmd.Start, cmd.End)
        } else if cmd.Type == 'low' {
            runes = low(runes, cmd.Start, cmd.End)
        } else if cmd.Type == 'cap' {
            runes = cap(runes, cmd.Start, cmd.End)

        }
    }
}