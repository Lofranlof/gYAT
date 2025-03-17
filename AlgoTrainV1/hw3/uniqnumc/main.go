package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// долбоёбы не могут дать мне провалидированный инпут,
// поэтому пришлось наговнокодить
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
			panic(err)
		}
		out = append(out, elInt)
	}
	return out
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	m := make(map[int]struct{}, len(input))
	for _, num := range input {
		m[num] = struct{}{}
	}

	fmt.Println(len(m))
}
