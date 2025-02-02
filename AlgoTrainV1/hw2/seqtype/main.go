package main

import (
	"bufio"
	"fmt"
	"log"
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
		if elInt == -2000000000 {
			break
		}
		out = append(out, elInt)
	}
	return out
}

func isAsc(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] >= in[i+1] {
			return false
		}
	}
	return true
}

func isDesc(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] <= in[i+1] {
			return false
		}
	}
	return true
}

func isConst(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] != in[i+1] {
			return false
		}
	}
	return true
}

func isWeakAsc(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] > in[i+1] {
			return false
		}
	}
	return true
}

func isWeakDesc(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] < in[i+1] {
			return false
		}
	}
	return true
}

func getListType(in []int) int {
	if isConst(in) {
		return CONSTANT
	} else if isAsc(in) {
		return ASCENDING
	} else if isDesc(in) {
		return DESCENDING
	} else if isWeakAsc(in) {
		return WEAKLY_ASCENDING
	} else if isWeakDesc(in) {
		return WEAKLY_DESCENDING
	} else {
		return RANDOM
	}
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	switch getListType(input) {
	case CONSTANT:
		fmt.Println("CONSTANT")
	case ASCENDING:
		fmt.Println("ASCENDING")
	case WEAKLY_ASCENDING:
		fmt.Println("WEAKLY ASCENDING")
	case DESCENDING:
		fmt.Println("DESCENDING")
	case WEAKLY_DESCENDING:
		fmt.Println("WEAKLY DESCENDING")
	case RANDOM:
		fmt.Println("RANDOM")
	}
}
