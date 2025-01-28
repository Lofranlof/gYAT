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
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func writeAnsToFile(path, ans string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(ans)
	if err != nil {
		return err
	}
	return nil
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

func calculateNumberOfNuts(alloyWeight, barWeight, nutWeight int) int {
	nuts := 0
	if nutWeight > barWeight || barWeight > alloyWeight {
		return 0
	}
	for ; alloyWeight >= barWeight; alloyWeight -= barWeight {
		alloyWeight += barWeight % nutWeight
		nuts += barWeight / nutWeight
	}
	return nuts
}

func main() {
	input := convertStrArrToIntArr(strings.Split(getInputFromFileInLines("input.txt")[0], " "))
	fmt.Printf("%d\n", calculateNumberOfNuts(input[0], input[1], input[2]))
}
