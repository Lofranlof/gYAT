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

func IsSymmetric(in []int) bool {
	i, j := 0, len(in)-1
	for i <= j {
		if in[i] != in[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func ReverseArray(in []int) []int {
	i, j := 0, len(in)-1
	for i <= j {
		in[i], in[j] = in[j], in[i]
		i++
		j--
	}
	return in
}

// возвращает индекс, до которого нужно разслайсить исходный массив и развернуть его
// чтобы получить необходимую добавку до симметричного массива
func GetSymmetricEnding(in []int) []int {
	endings := make([][]int, 0, len(in))
	for l := len(in) - 2; l > 0; l-- {
		ending := make([]int, 0, l)
		orig := make([]int, 0, len(in))
		orig = append(orig, in...)

		for i := l; i >= 0; i-- {
			ending = append(ending, orig[i])
		}

		orig = append(orig, ending...)

		if IsSymmetric(orig) {
			endings = append(endings, ending)
		}
	}
	return endings[len(endings)-1]
}

// Попробовать сделать без доп памяти используя реслайсы
func GetSymmetricEndingSmart(in []int) []int {
	return []int{}
}

func printAns(ans []int) {
	fmt.Println(len(ans))
	for i, el := range ans {
		if i < len(ans)-1 {
			fmt.Printf("%d ", el)
		} else {
			fmt.Print(el)
		}
	}
	fmt.Println()
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))[1:]
	if IsSymmetric(input) {
		fmt.Println(0)
		return
	}
	minEnding := GetSymmetricEnding(input)
	printAns(minEnding)
}
