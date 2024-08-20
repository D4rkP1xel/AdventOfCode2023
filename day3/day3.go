package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var lines []string = utils.GetLines("input.txt")
	part2(lines)
}

func part1(lines []string) {
	var sum uint32 = 0
	for y, line := range lines {
		var numXStartPos uint8
		var numXEndPos uint8
		var isReadingNumber bool = false
		for x, char := range line {
			if !unicode.IsDigit(char) {
				isReadingNumber = false
				continue

			}
			// ------reading a number on current position------

			if !isReadingNumber {
				//start of a number
				isReadingNumber = true
				numXStartPos = uint8(x)
			}
			if x+1 > len(line)-1 || !unicode.IsDigit(rune(line[x+1])) {
				isReadingNumber = false
				numXEndPos = uint8(x)
				xToCheckAux := utils.Clamp(int(numXStartPos)-1, 0, len(line)-1)
				xToCheckAuxMax := utils.Clamp(int(numXEndPos)+1, 0, len(line)-1)
				fmt.Printf("%s %s", xToCheckAux, xToCheckAuxMax)
				for yToCheck := utils.Clamp(int(y)-1, 0, len(lines)-1); yToCheck <= utils.Clamp(int(y)+1, 0, len(lines)-1); yToCheck++ {
					for xToCheck := utils.Clamp(int(numXStartPos)-1, 0, len(line)-1); xToCheck <= utils.Clamp(int(numXEndPos)+1, 0, len(line)-1); xToCheck++ {
						var charToCheck = []rune(lines[yToCheck])[xToCheck]
						if unicode.IsSymbol(charToCheck) || unicode.IsPunct(charToCheck) && charToCheck != '.' {
							valueToSum, err := strconv.Atoi(line[numXStartPos : numXEndPos+1])
							fmt.Printf("LINE: %s NUM: %d\n", line, valueToSum)
							if err != nil {
								fmt.Printf("ERROR: %s\n", err)
								return
							}
							sum += uint32(valueToSum)
						}
					}
				}

			}

		}
	}
	fmt.Printf("%d\n", sum)
}

type gearPos struct {
	X uint16
	Y uint16
}

func part2(lines []string) {
	var sum uint32 = 0
	var gearList = make(map[gearPos][]uint16)

	for y, line := range lines {
		var numXStartPos uint8
		var numXEndPos uint8
		var isReadingNumber bool = false
		for x, char := range line {
			if !unicode.IsDigit(char) {
				isReadingNumber = false
				continue

			}
			// ------reading a number on current position------

			if !isReadingNumber {
				//start of a number
				isReadingNumber = true
				numXStartPos = uint8(x)
			}
			if x+1 > len(line)-1 || !unicode.IsDigit(rune(line[x+1])) {
				isReadingNumber = false
				numXEndPos = uint8(x)

				for yToCheck := utils.Clamp(int(y)-1, 0, len(lines)-1); yToCheck <= utils.Clamp(int(y)+1, 0, len(lines)-1); yToCheck++ {
					for xToCheck := utils.Clamp(int(numXStartPos)-1, 0, len(line)-1); xToCheck <= utils.Clamp(int(numXEndPos)+1, 0, len(line)-1); xToCheck++ {
						var charToCheck = []rune(lines[yToCheck])[xToCheck]
						if charToCheck == '*' {
							valueToSum, err := strconv.Atoi(line[numXStartPos : numXEndPos+1])
							if err != nil {
								fmt.Printf("ERROR: %s\n", err)
								return
							}
							key := gearPos{X: uint16(xToCheck), Y: uint16(yToCheck)}
							gearsValues, exists := gearList[key]
							if !exists {
								gearList[key] = []uint16{uint16(valueToSum)}
							} else {
								gearList[key] = append(gearsValues, uint16(valueToSum))
							}

						}
					}
				}

			}

		}
	}

	for _, gearValues := range gearList {
		if len(gearValues) == 2 {
			sum += uint32(gearValues[0]) * uint32(gearValues[1])
		}
	}
	fmt.Printf("%d\n", sum)
}
