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

func sortInitialElements(in []int) []int {
	if in[0] < in[1] {
		in[0], in[1] = in[1], in[0]
	}
	if in[2] > in[0] {
		in[0], in[2] = in[2], in[0]
		in[1], in[2] = in[2], in[1]
	} else if in[2] > in[1] {
		in[2], in[1] = in[1], in[2]
	}
	return in
}

func findMaxNumbers(in []int) (first int, second int, third int) {
	initElements := sortInitialElements(in)
	first, second, third = initElements[0], initElements[1], initElements[2]
	for i := 3; i < len(in); i++ {
		if in[i] >= first {
			third = second
			second = first
			first = in[i]
		} else if in[i] >= second {
			third = second
			second = in[i]
		} else if in[i] > third {
			third = in[i]
		}
	}
	return first, second, third
}

func findMinNumbers(in []int) (first int, second int) {
	first, second = min(in[0], in[1]), max(in[0], in[1])
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
	fstMax, sndMax, trdMax := findMaxNumbers(input)
	fstMin, sndMin := findMinNumbers(input)

	if len(input) == 3 {
		fmt.Printf("%d %d %d\n", input[0], input[1], input[2])
		return
	}
	multiplier := max(fstMax, sndMax, trdMax)

	if float64(fstMax)*float64(sndMax)*float64(trdMax) >= float64(multiplier)*float64(fstMin)*float64(sndMin) {
		fmt.Printf("%d %d %d\n", fstMax, sndMax, trdMax)
	} else {
		fmt.Printf("%d %d %d\n", multiplier, fstMin, sndMin)
	}

}
