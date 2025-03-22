package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func getInputFromFileInLines(path string) [][]string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	linesAr := make([]string, 0)
	var line strings.Builder
	for _, el := range strData {
		if el == '\n' {
			linesAr = append(linesAr, line.String())
			line = strings.Builder{}
			continue
		}
		if !unicode.IsSpace(el) {
			line.WriteRune(el)
		} else {
			line.WriteString(" ")
		}
	}

	ans := make([][]string, 0)
	for _, l := range linesAr {
		dirtyStr := strings.Split(l, " ")
		cleanStr := make([]string, 0, len(dirtyStr))

		for _, el := range dirtyStr {
			if el != " " && el != "" {
				cleanStr = append(cleanStr, el)
			}
		}
		ans = append(ans, cleanStr)
	}

	return ans
}

func getInputFromStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for range 2 {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
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

func getNumsIntersection(nums1, nums2 []int) []int {
	set1 := make(map[int]struct{}, len(nums1))
	set2 := make(map[int]struct{}, len(nums2))

	for _, n := range nums1 {
		set1[n] = struct{}{}
	}

	for _, n := range nums2 {
		set2[n] = struct{}{}
	}

	minSet := set1

	if len(set2) < len(set1) {
		minSet = set2
	}

	ans := make([]int, 0, len(minSet))
	for k := range minSet {
		if _, ok := set1[k]; ok {
			if _, ok := set2[k]; ok {
				ans = append(ans, k)
			}
		}
	}
	slices.Sort(ans)
	return ans
}

func main() {
	input := getInputFromFileInLines("input.txt")

	if len(input) < 2 {
		return
	}

	nums1, nums2 := convertStrArrToIntArr(input[0]), convertStrArrToIntArr(input[1])

	ans := getNumsIntersection(nums1, nums2)
	for i, n := range ans {
		if i < len(ans)-1 {
			fmt.Printf("%d ", n)
		} else {
			fmt.Printf("%d\n", n)
		}
	}
}
