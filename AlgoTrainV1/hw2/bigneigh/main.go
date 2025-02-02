package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func bigNeighCount(in []int) int {
	count := 0
	for i := 1; i < len(in)-1; i++ {
		if in[i-1] < in[i] && in[i+1] < in[i] {
			count++
		}
	}
	return count
}

func main() {
	input := convertStrArrToIntArr(strings.Split(getInputFromFileInLines("input.txt")[0], " "))
	fmt.Println(bigNeighCount(input))
}
