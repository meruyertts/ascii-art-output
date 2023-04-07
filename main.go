package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	input := os.Args[1:]
	if len(input) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --output=<fileName.txt>")
		return
	}
	myStr := input[0]
	fileName := fileNameCheck(input[1])
	re := regexp.MustCompile("--output=")
	if !re.MatchString(input[2]) {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\nEX: go run . something standard --output=<fileName.txt>")
		return
	}
	nameFile := re.Split(input[2], -1)

	outputFile := nameFile[1]
	if !isASCII(myStr) {
		fmt.Println("non-ASCII character was entered")
		return
	}
	if len(myStr) == 0 {
		return
	}
	if fileName == "" {
		return
	}
	if lineCounter(fileName) != nil {
		return
	}
	splitWord(myStr, fileName, outputFile)
}

func splitWord(myStr, myFile, outputFile string) {
	re := regexp.MustCompile(`\\n`)
	newStr := re.Split(myStr, -1)
	var err error
	var myArr [8]string
	var finalStr string
	for i := 0; i < len(newStr); i++ {
		if len(newStr[i]) > 0 {
			myArr, err = printWord(newStr[i], myFile)
			if err != nil {
				return
			}
			for _, i := range myArr {
				finalStr += i + "\n"
			}

		}
		if newStr[i] == "" {
			fmt.Println("")
		}
	}
	writeFile(finalStr, outputFile)
}

func fileNameCheck(fName string) string {
	myFile := ""
	switch fName {
	case "standard":
		myFile = "standard.txt"
	case "shadow":
		myFile = "shadow.txt"
	case "thinkertoy":
		myFile = "thinkertoy.txt"
	}

	return myFile
}

func lineCounter(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount < 855 {
		return errors.New("file does not contain all characters")
	}
	return nil
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func printWord(s, fileName string) ([8]string, error) {
	var err error
	myLine := ""
	myarray := [8]string{}
	for _, char := range s {
		for line := 2; line <= 9; line++ {
			myrune := int(char)
			for i := ' '; i <= '~'; i++ {
				j := (int(i) - ' ') * 9
				if myrune == int(i) {
					firstline, err := readExactLine(fileName, line+j)
					if err != nil {
						log.Print(err)
						return [8]string{}, err
					}
					myLine += firstline
				}
			}
		}
		temp := strings.Split(myLine, "\n")
		for index, s := range temp[:len(temp)-1] {
			myarray[index] += s
		}
		myLine = ""
	}
	return myarray, err
}

func writeFile(arr string, outFile string) {
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	arr = arr + "\n"
	_, err2 := f.WriteString(arr)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func readExactLine(fileName string, lineNumber int) (string, error) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	br := bufio.NewReader(inputFile)
	for i := 1; i < lineNumber; i++ {
		_, _ = br.ReadString('\n')
	}
	str, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}

	return str, nil
}
