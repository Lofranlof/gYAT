package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getInputFromFileInLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var b strings.Builder
	b.Grow(len(data))

	strData := string(data)
	for _, el := range strData {
		if unicode.IsSpace(el) {
			b.WriteString(" ")
		} else {
			b.WriteRune(el)
		}
	}

	dirtyStr := strings.Split(b.String(), " ")
	cleanStr := make([]string, 0, len(dirtyStr))

	for _, el := range dirtyStr {
		if el != " " && el != "" {
			cleanStr = append(cleanStr, el)
		}
	}
	return cleanStr
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

func findMaxNumbers(in []int) (first int, second int) {
	first, second = in[0], in[1]
	for i := 2; i < len(in); i++ {
		if in[i] >= first {
			second = first
			first = in[i]
		} else if in[i] > second {
			second = in[i]
		}
	}
	return first, second
}

func findMinNumbers(in []int) (first int, second int) {
	first, second = in[0], in[1]
	for i := 2; i < len(in); i++ {
		if in[i] <= first {
			second = first
			first = in[i]
		} else if in[i] < second {
			second = in[i]
		}
	}
	return first, second
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	firstMax, secondMax := findMaxNumbers(input)
	firstMin, secondMin := findMinNumbers(input)

	if len(input) == 2 {
		if input[0] < input[1] {
			fmt.Printf("%d %d\n", input[0], input[1])
		} else {
			fmt.Printf("%d %d\n", input[1], input[0])
		}
		return
	}

	if firstMax*secondMax > firstMin*secondMin {
		fmt.Printf("%d %d\n", secondMax, firstMax)
	} else {
		fmt.Printf("%d %d\n", firstMin, secondMin)
	}
}
