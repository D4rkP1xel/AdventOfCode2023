package main

import (
	"adventofcode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.GetLines("input.txt")
	part2(lines)
}

func part1(lines []string) {
	timeSplit := strings.Fields(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	distanceSplit := strings.Fields(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	racesWon := make(map[uint8]uint32)
	for raceIndex := range len(timeSplit) {
		timeValue, timeErr := strconv.Atoi(timeSplit[raceIndex])
		distanceValue, distanceErr := strconv.Atoi(distanceSplit[raceIndex])
		if timeErr != nil {
			fmt.Printf("TIME ERROR: %s", timeSplit[raceIndex])
		}

		if distanceErr != nil {
			fmt.Printf("DISTANCE ERROR: %s", distanceSplit[raceIndex])
		}

		for velocity := 0; velocity <= timeValue; velocity++ {
			timeTravelled := (uint32(timeValue) - uint32(velocity))
			var distanceTravelled uint32 = timeTravelled * uint32(velocity)
			if distanceTravelled > uint32(distanceValue) {
				racesWon[uint8(raceIndex)]++
			}
		}
	}
	sum := racesWon[0]
	for i := 1; i < len(racesWon); i++ {
		sum *= racesWon[uint8(i)]
	}
	fmt.Println(sum)
}
func part2(lines []string) {
	timeSplit := strings.Fields(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	distanceSplit := strings.Fields(strings.TrimSpace(strings.Split(lines[1], ":")[1]))

	time, timeErr := strconv.Atoi(timeSplit[0] + timeSplit[1] + timeSplit[2] + timeSplit[3])
	if timeErr != nil {
		fmt.Printf("TIME ERROR: %s", timeErr)
	}

	distance, distanceErr := strconv.Atoi(distanceSplit[0] + distanceSplit[1] + distanceSplit[2] + distanceSplit[3])
	if distanceErr != nil {
		fmt.Printf("DISTANCE ERROR: %s", distanceErr)
	}

	var numRacesWon uint64
	var velocity uint64
	for velocity = 0; velocity <= uint64(time); velocity++ {
		timeTravelled := (uint64(time) - uint64(velocity))
		var distanceTravelled uint64 = timeTravelled * uint64(velocity)
		if distanceTravelled > uint64(distance) {
			numRacesWon++
		}
	}

	fmt.Println(numRacesWon)
}
