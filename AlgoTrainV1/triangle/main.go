package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getInputFromFileInLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
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

func main() {
	input, err := getInputFromFileInLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	output := "NO"

	a, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}
	c, err := strconv.Atoi(input[2])
	if err != nil {
		log.Fatal(err)
	}

	if a+b > c && c+b > a && c+a > b {
		output = "YES"
	}

	writeAnsToFile("output.txt", output)
}
