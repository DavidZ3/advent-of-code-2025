package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INIT_POS = 50
var N_DIAL_ENTRIES = 100

func calcNewPos(oldPos int, nRotations int, left bool) int {
	var newPos int
	if left {
		newPos = (oldPos - nRotations) % N_DIAL_ENTRIES
	} else {
		newPos = (oldPos + nRotations) % N_DIAL_ENTRIES
	}
	if newPos < 0 {
		newPos += N_DIAL_ENTRIES
	}
	return newPos
}

func countZeros(rotationStrArr []string, initialPos int) int {
	nZeros := 0
	currPos := initialPos
	for _, rotationStr := range rotationStrArr {
		left := false
		if rotationStr[0] == 'L' {
			left = true
		}
		nRotations, err := strconv.Atoi(rotationStr[1:])
		if err != nil {
			panic(err)
		}
		currPos = calcNewPos(currPos, nRotations, left)
		if currPos == 0 {
			nZeros++
		}
	}
	return nZeros
}

func calcNumZeroCrossings(oldPos int, nRotations int, left bool) int {
	nZeroCrossings := 0

	/* We know that (a + k*N) % N = a. So we can calculate the number of times we pass zero and exclude it */
	nZeroCrossings += nRotations / N_DIAL_ENTRIES

	/* We now want to work within the region [0, N_DIAL_ENTIRES) - residue of modulo N_DIAL_ENTRIES */
	nRotations = nRotations % N_DIAL_ENTRIES

	/* If there's no more residule rotations then finish */
	if nRotations == 0 {
		return nZeroCrossings
	}

	/* Catch the case where we start at 0 but go left to avoid double counting zeros */
	if left && oldPos == 0 {
		return nZeroCrossings
	}

	var newPos = 0
	if left {
		newPos = oldPos - nRotations
	} else {
		newPos = oldPos + nRotations
	}
	if (newPos <= 0) || (newPos >= N_DIAL_ENTRIES) {
		nZeroCrossings++
	}
	return nZeroCrossings
}

func countZeroCrossings(rotationStrArr []string, initialPos int) int {
	nZeroCrossings := 0
	currPos := initialPos
	for _, rotationStr := range rotationStrArr {
		left := false
		if rotationStr[0] == 'L' {
			left = true
		}
		nRotations, err := strconv.Atoi(rotationStr[1:])
		if err != nil {
			panic(err)
		}
		nZeroCrossings += calcNumZeroCrossings(currPos, nRotations, left)
		currPos = calcNewPos(currPos, nRotations, left)
	}
	return nZeroCrossings
}

func main() {
	fmt.Println("Hello World! This is my first Go Lang program :D")
	rotationsData, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	rotationStrArr := strings.Split(string(rotationsData), "\r\n")
	fmt.Println("Part 1: Number of times we land on zero: ", countZeros(rotationStrArr, INIT_POS))
	fmt.Println("Part 2: Number of times we pass zero: ", countZeroCrossings(rotationStrArr, INIT_POS))
}
