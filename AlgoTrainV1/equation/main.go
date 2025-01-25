package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	a, b, c := input[0], input[1], input[2]
	// fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	if a == 0 && b == 0 && c == 0 {
		fmt.Println("MANY SOLUTIONS")
		return
	}
	if a == 0 && b < 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	if a == 0 && b >= 0 {
		ans := int(math.Pow(float64(c), 2))
		if ans == b {
			fmt.Println("MANY SOLUTIONS")
		} else {
			fmt.Println("NO SOLUTION")
		}
		return
	}
	// check if there is a remainder after doing a divison, we don't want any
	if rem := (int(math.Pow(float64(c), 2)) - b) % a; rem != 0 {
		fmt.Println("NO SOLUTION")
	} else {
		fmt.Println((int(math.Pow(float64(c), 2)) - b) / a)
	}
}
