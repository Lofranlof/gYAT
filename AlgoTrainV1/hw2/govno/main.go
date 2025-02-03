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
			log.Fatal(err)
		}
		out = append(out, elInt)
	}
	return out
}

func findWinnerMinIndex(in []int) int {
	winnerScore := -1
	// find max score
	for i, score := range in {
		if i == 0 || score > winnerScore {
			winnerScore = score
		}
	}

	// find indexes of contestants with max score a.k.a winners indexes
	for i := range in {
		if in[i] == winnerScore {
			return i
		}
	}
	return 1000000
}

func findVasyaPlace(in []int) int {
	vasyaScore := -1
	k := 0
	winnerMinInd := findWinnerMinIndex(in)
	for i := 1; i < len(in)-1; i++ {
		if in[i]%10 == 5 && in[i] > in[i+1] && winnerMinInd < i {
			vasyaScore = max(in[i], vasyaScore)
		}
	}

	if vasyaScore == -1 {
		return 0
	}

	// сколько человек метнуло лепёшку дальше Васи
	for _, score := range in {
		if score > vasyaScore {
			k++
		}
	}

	return k + 1
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt")[1:])
	fmt.Println(findVasyaPlace(input))
}
