package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPidValid(pid string) bool {
	/*
	 * A product ID is said to be invalid if it meets ANY of the following crieria:
	 *	- Leading zero (e.g. 0 and 0123 are invalid, but 123 is valid)
	 *	- Any repeated sequence of numbers (e.g. 11, 1212, 123123, ...)
	 * An interesting observation is that all repeated sequences will be identical starting from the half way point
	 */
	// Case 1: Leading Zero
	if pid[0] == '0' {
		return false
	}

	// Case 2 not possible if odd
	if len(pid)%2 != 0 {
		return true
	}
	halfLen := len(pid) / 2

	for i := range halfLen {
		if pid[i] != pid[halfLen+i] {
			return true
		}

	}
	return false
}

func atoiWithPanic(intStr string) int {
	i, err := strconv.Atoi(intStr)
	if err != nil {
		panic(err)
	}
	return i
}

func getInvalidPids(pidList []string) []string {
	invalidPidList := []string{}
	for i, pid := range pidList {
		fmt.Println(i, " : ", pid)
		if !isPidValid(pid) {
			invalidPidList = append(invalidPidList, pid)
			// fmt.Println(invalidPidList)
		}
	}
	return invalidPidList
}

func parsePidStrPairFromStr(pidPairStr string, delimitter string) (string, string) {
	pidStrArr := strings.Split(pidPairStr, delimitter)
	if len(pidStrArr) != 2 {
		panic("Error: Product ID range != product IDs")
	}
	return pidStrArr[0], pidStrArr[1]
}

func getPidListFromPidData(pidData string) []string {
	pidPairsArr := strings.Split(string(pidData), ",")
	pidList := make([]string, 2*len(pidPairsArr))
	for i, pair := range pidPairsArr {
		pidList[i], pidList[i+1] = parsePidStrPairFromStr(pair, "-")

	}
	return pidList
}

func sumPids(pidStrList []string) int {
	total := 0
	for _, pid := range pidStrList {
		total += atoiWithPanic(pid)
	}
	return total

}

func main() {
	pidData, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	pidList := getPidListFromPidData(string(pidData))
	invalidPidList := getInvalidPids(pidList)
	fmt.Println("Part 1: Sum of invalid pids: ", sumPids(invalidPidList))
}
