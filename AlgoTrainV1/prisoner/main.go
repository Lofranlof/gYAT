package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func doesBrickFit(a, b, c, d, e int) bool {
	if (e >= a && d >= b) || (e >= b && d >= a) {
		return true
	} else if (e >= b && d >= c) || (e >= c && d >= b) {
		return true
	} else if (e >= a && d >= c) || (e >= c && d >= a) {
		return true
	}
	return false
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	if doesBrickFit(input[0], input[1], input[2], input[3], input[4]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
