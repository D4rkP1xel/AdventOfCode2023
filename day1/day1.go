package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var lines []string = utils.GetLines("input.txt")
	//sumPart1(lines)
	sumPart2(lines)
	utils.WaitForInput()
}

func sumPart1(lines []string) {
	var sum uint32 = 0

	for _, line := range lines {
		var rawDigits []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				rawDigits = append(rawDigits, char)
			}
		}
		fmt.Printf("\nLine: %s, Raw Digits: %s\n", line, string(rawDigits))

		if len(rawDigits) > 0 {
			var finalDigits []rune = []rune{rawDigits[0], rawDigits[len(rawDigits)-1]}

			num, err := strconv.Atoi(string(finalDigits))
			fmt.Printf("Final Digits: %s %d \n", string(finalDigits), num)
			if err != nil {
				fmt.Println("Error converting digits to number:", err)
				continue
			}
			sum += uint32(num)
		}
	}
	fmt.Println(sum)
}

func sumPart2(lines []string) {
	var sum uint32 = 0
	numMap := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for _, line := range lines {
		var rawDigits []rune
		for i, char := range line {
			if unicode.IsDigit(char) {
				rawDigits = append(rawDigits, char)
			} else if unicode.IsLetter(char) {
				var digit rune = isLetterTheStartOfANumber(numMap, line[i:])
				if digit != 0 {
					rawDigits = append(rawDigits, digit)
				}
			}
		}
		fmt.Printf("\nLine: %s, Raw Digits: %s\n", line, string(rawDigits))

		if len(rawDigits) > 0 {
			var finalDigits []rune = []rune{rawDigits[0], rawDigits[len(rawDigits)-1]}

			num, err := strconv.Atoi(string(finalDigits))
			fmt.Printf("Final Digits: %s %d \n", string(finalDigits), num)
			if err != nil {
				fmt.Println("Error converting digits to number:", err)
				continue
			}
			sum += uint32(num)
		}
	}
	fmt.Println(sum)
}

func getNumberByPrefix(numMap map[string]rune, prefix string) map[string]rune {
	var numMapAux map[string]rune = make(map[string]rune)

	for key, value := range numMap {
		if strings.HasPrefix(key, prefix) {
			numMapAux[key] = value
		}
	}
	return numMapAux
}

func isLetterTheStartOfANumber(numMap map[string]rune, restOfLine string) rune {
	// 3 is the min number of letters in one possible digit, 5 is the max
	if len(restOfLine) < 3 {
		return 0
	}

	for i := 3; i <= 5; i++ {
		if i > len(restOfLine) {
			break
		}
		word := restOfLine[0:i]
		//fmt.Printf("%s \n", word)

		// no possible combination
		if len(getNumberByPrefix(numMap, word)) == 0 {
			return 0
		}

		value, exists := numMap[word]
		if exists {
			return value
		}
	}

	return 0
}
