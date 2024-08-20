package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(data))

	var lines []string = strings.Split(string(data), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	if strings.Compare(lines[len(lines)-1], "") == 0 {
		lines = lines[0 : len(lines)-1]
	}
	return lines
}

func WaitForInput() {
	// Wait for user input to close
	fmt.Println("Press Enter to exit...")
	var input string
	fmt.Scanln(&input)
}
