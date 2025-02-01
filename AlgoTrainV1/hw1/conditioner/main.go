package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	var output int
	temp, command := input[0], input[1]
	tempAr := strings.Split(temp, " ")

	tRoom, err := strconv.Atoi(tempAr[0])
	if err != nil {
		log.Fatal(err)
	}

	tCond, err := strconv.Atoi(tempAr[1])
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "freeze":
		if tRoom <= tCond {
			output = tRoom
		} else {
			output = tCond
		}
	case "heat":
		if tRoom <= tCond {
			output = tCond
		} else {
			output = tRoom
		}
	case "auto":
		output = tCond
	case "fan":
		output = tRoom
	}

	writeAnsToFile("output.txt", strconv.Itoa(output))
}
