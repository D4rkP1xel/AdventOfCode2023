package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part2(utils.GetLines("input.txt"))
}

func part1(lines []string) {
	var sum uint32

	for _, line := range lines {
		var cardSum uint32 = 0
		var cardSplit []string = strings.Split(line, ":")
		var numbersSplit []string = strings.Split(cardSplit[1], "|")
		var winningNumbersStringList []string = strings.Fields(strings.TrimSpace(numbersSplit[0]))
		var playerNumbersStringList []string = strings.Fields(strings.TrimSpace(numbersSplit[1]))
		var winningNumbers map[uint8]bool = make(map[uint8]bool)

		for _, winningNumber := range winningNumbersStringList {
			value, err := strconv.Atoi(winningNumber)
			if err != nil {
				fmt.Printf("ERROR 1: %s", err)
				return
			}
			winningNumbers[uint8(value)] = true
		}

		for _, playerNumber := range playerNumbersStringList {
			value, err := strconv.Atoi(playerNumber)
			if err != nil {
				fmt.Printf("ERROR 2: %s", err)
				return
			}

			if winningNumbers[uint8(value)] {
				if cardSum == 0 {
					cardSum = 1
				} else {
					cardSum *= 2
				}
			}

		}
		sum += uint32(cardSum)
	}
	fmt.Printf("%d", sum)
}

func part2(lines []string) {
	var sum int
	var cardInstances map[int]int = make(map[int]int)
	for i := range lines {
		cardInstances[int(i)] = 1
	}

	for cardIndex, line := range lines {
		for i := 0; i < int(cardInstances[int(cardIndex)]); i++ {
			sum++
			var numWinningNumbers int = 0
			var cardSplit []string = strings.Split(line, ":")
			var numbersSplit []string = strings.Split(cardSplit[1], "|")
			var winningNumbersStringList []string = strings.Fields(strings.TrimSpace(numbersSplit[0]))
			var playerNumbersStringList []string = strings.Fields(strings.TrimSpace(numbersSplit[1]))
			var winningNumbers map[int]bool = make(map[int]bool)
			for _, winningNumber := range winningNumbersStringList {
				value, err := strconv.Atoi(winningNumber)
				if err != nil {
					fmt.Printf("ERROR 1: %s", err)
					return
				}
				winningNumbers[int(value)] = true
			}

			for _, playerNumber := range playerNumbersStringList {
				value, err := strconv.Atoi(playerNumber)
				if err != nil {
					fmt.Printf("ERROR 2: %s", err)
					return
				}

				if winningNumbers[int(value)] {
					numWinningNumbers++
					cardInstances[int(cardIndex)+numWinningNumbers]++
				}
			}
		}
	}

	fmt.Printf("%d", sum)
}
