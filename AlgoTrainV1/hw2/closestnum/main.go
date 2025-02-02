package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	CONSTANT = iota
	ASCENDING
	WEAKLY_ASCENDING
	DESCENDING
	WEAKLY_DESCENDING
	RANDOM
)

func getInputFromFileInLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
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

func prepareInput(input []string) (ar []int, numToFind int) {
	ar = convertStrArrToIntArr(strings.Split(input[1], " "))
	numToFind, err := strconv.Atoi(input[len(input)-1])
	if err != nil {
		log.Fatal()
	}
	return ar, numToFind
}

func findClosestNum(ar []int, numToFind int) int {
	var closestNum int
	for i := 0; i < len(ar); i++ {
		if i == 0 || math.Abs(float64(ar[i]-numToFind)) < math.Abs(float64(closestNum-numToFind)) {
			closestNum = ar[i]
		}
	}
	return closestNum
}

func main() {
	input := getInputFromFileInLines("input.txt")
	array, numToFind := prepareInput(input)
	fmt.Println(findClosestNum(array, numToFind))
}
