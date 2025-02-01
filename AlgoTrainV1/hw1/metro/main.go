package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const MAX_WAIT_TIME = 1000000

func getInputFromFileInLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func writeAnsToFile(path, ans string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(ans)
	if err != nil {
		return err
	}
	return nil
}

func convertStrArrToIntArr(in []string) []int {
	out := make([]int, 0, len(in))
	for _, el := range in {
		elInt, err := strconv.Atoi(el)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, elInt)
	}
	return out
}

func calcMinAmountOfTrains(interval, waitTime int) int {
	trainCount := 0
	trainCount = waitTime / (interval + 1)
	return trainCount
}

func calcMaxAmountOfTrains(interval, waitTime int) int {
	newWaitTime := waitTime - 1
	trainCount := (newWaitTime / (interval + 1)) + 1
	return trainCount
}

func calcMinMaxTime(interval1, interval2, trainCount1, trainCount2, maxTime int) (minWaitTime, maxWaitTime int, err error) {
	for i := 0; i <= maxTime; i++ {
		minAmountOfTrainsForInterval1, maxAmountOfTrainsForInterval1 := calcMinAmountOfTrains(interval1, i), calcMaxAmountOfTrains(interval1, i)
		minAmountOfTrainsForInterval2, maxAmountOfTrainsForInterval2 := calcMinAmountOfTrains(interval2, i), calcMaxAmountOfTrains(interval2, i)
		if (trainCount1 == minAmountOfTrainsForInterval1 || trainCount1 == maxAmountOfTrainsForInterval1) && (trainCount2 == minAmountOfTrainsForInterval2 || trainCount2 == maxAmountOfTrainsForInterval2) {
			if minWaitTime == 0 {
				minWaitTime = i
			} else {
				maxWaitTime = i
			}
		}
	}
	if minWaitTime != 0 && maxWaitTime == 0 {
		maxWaitTime = minWaitTime
	}

	return minWaitTime, maxWaitTime, nil
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	// t1 := time.Now()
	// defer func() {
	// 	fmt.Printf("Time spent: %v\n", time.Since(t1))
	// }()
	min, max, _ := calcMinMaxTime(input[0], input[1], input[2], input[3], MAX_WAIT_TIME)
	if min == 0 && max == 0 {
		fmt.Printf("%d\n", -1)
		return
	}
	fmt.Printf("%d %d\n", min, max)
}
