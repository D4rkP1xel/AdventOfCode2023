package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var lines []string = utils.GetLines("input.txt")
	part2(lines)
}

func part1(lines []string) {
	var sum uint32 = 0
	const MAX_NUM_RED_CUBES = 12
	const MAX_NUM_GREEN_CUBES = 13
	const MAX_NUM_BLUE_CUBES = 14
	for _, line := range lines {
		isGamePossible := true

		var twoDotsSplit = strings.Split(line, ":")
		gameNum, err := strconv.Atoi(strings.Split(twoDotsSplit[0], " ")[1])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		sum += uint32(gameNum)
		var sets []string = strings.Split(twoDotsSplit[1], ";")
		for _, setString := range sets {
			if !isGamePossible {
				break
			}
			var cubesSplits []string = strings.Split(setString, ",")
			for _, cubeString := range cubesSplits {
				if !isGamePossible {
					break
				}
				cubeString = strings.TrimSpace(cubeString)
				var cubeSplit []string = strings.Split(cubeString, " ")
				var numOfCubesString string = cubeSplit[0]
				var cubeColor string = cubeSplit[1]
				numOfCubes, err := strconv.Atoi(numOfCubesString)
				if err != nil {
					fmt.Printf("ERROR: invalid number of cubes: %s %s \n", err, numOfCubesString)
					return
				}
				switch cubeColor {
				case "red":
					if numOfCubes > MAX_NUM_RED_CUBES {
						isGamePossible = false
						sum -= uint32(gameNum)
					}

				case "green":
					if numOfCubes > MAX_NUM_GREEN_CUBES {
						isGamePossible = false
						sum -= uint32(gameNum)
					}

				case "blue":
					if numOfCubes > MAX_NUM_BLUE_CUBES {
						isGamePossible = false
						sum -= uint32(gameNum)
					}
				}
			}
		}
	}

	fmt.Printf("%d \n", sum)
}

func part2(lines []string) {
	var sum uint32 = 0

	for _, line := range lines {
		var maxRedCubes uint8 = 0
		var maxGreenCubes uint8 = 0
		var maxBlueCubes uint8 = 0
		var twoDotsSplit = strings.Split(line, ":")
		var sets []string = strings.Split(twoDotsSplit[1], ";")
		for setIndex, setString := range sets {
			var cubesSplits []string = strings.Split(setString, ",")
			for cubeIndex, cubeString := range cubesSplits {
				cubeString = strings.TrimSpace(cubeString)
				var cubeSplit []string = strings.Split(cubeString, " ")
				var numOfCubesString string = cubeSplit[0]
				var cubeColor string = cubeSplit[1]
				numOfCubes, err := strconv.Atoi(numOfCubesString)
				if err != nil {
					fmt.Printf("ERROR: invalid number of cubes: %s %s \n", err, numOfCubesString)
					return
				}
				switch cubeColor {
				case "red":
					if uint8(numOfCubes) > maxRedCubes {
						maxRedCubes = uint8(numOfCubes)
					}

				case "green":
					if uint8(numOfCubes) > maxGreenCubes {
						maxGreenCubes = uint8(numOfCubes)
					}

				case "blue":
					if uint8(numOfCubes) > maxBlueCubes {
						maxBlueCubes = uint8(numOfCubes)
					}
				}
				if setIndex == len(sets)-1 && cubeIndex == len(cubesSplits)-1 {
					//we are in the last cube of the line
					fmt.Printf("red: %d green: %d blue: %d \n", maxRedCubes, maxGreenCubes, maxBlueCubes)
					sum += uint32(maxRedCubes) * uint32(maxGreenCubes) * uint32(maxBlueCubes)
				}
			}
		}
	}
	fmt.Printf("%d \n", sum)
}
