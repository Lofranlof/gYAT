package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rectangle struct {
	height int
	width  int
}

func (r rectangle) area() int {
	return r.height * r.width
}

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

func calculateMinTableDimensions(l1, l2 rectangle) rectangle {
	tables := make([]rectangle, 0, 4)
	area := -1
	var minTable rectangle
	tables = append(tables, rectangle{height: max(l1.height, l2.height), width: l1.width + l2.width})
	tables = append(tables, rectangle{height: max(l1.height, l2.width), width: l1.width + l2.height})
	tables = append(tables, rectangle{height: max(l1.width, l2.height), width: l1.height + l2.width})
	tables = append(tables, rectangle{height: max(l1.width, l2.width), width: l1.height + l2.height})

	for i, table := range tables {
		if i == 0 || table.area() < area {
			area = table.area()
			minTable = table
		}
	}
	return minTable
}

func main() {
	input := convertStrArrToIntArr(strings.Split(getInputFromFileInLines("input.txt")[0], " "))
	table := calculateMinTableDimensions(rectangle{height: input[0], width: input[1]}, rectangle{height: input[2], width: input[3]})
	fmt.Printf("%d %d\n", table.height, table.width)
}
