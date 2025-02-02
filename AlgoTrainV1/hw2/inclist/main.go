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

func isIncrList(in []int) bool {
	isInc := true
	if len(in) == 0 {
		return true
	} else if len(in) == 1 {
		return true // ???
	}

	for i := range in {
		if i == 0 || in[i] > in[i-1] {
			continue
		} else {
			isInc = false
		}
	}
	return isInc
}

func main() {
	inStr := getInputFromFileInLines("input.txt")
	if len(inStr) == 0 {
		fmt.Println("YES")
		return
	}
	input := convertStrArrToIntArr(strings.Split(inStr[0], " "))
	if isIncrList(input) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
