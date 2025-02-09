package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// normalized to contain values starting from 0
type mine struct {
	x int
	y int
}

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

func parseMinesCoords(in []int) []mine {
	mines := make([]mine, 0, len(in)/2)
	for i := 0; i+1 < len(in); i += 2 {
		// adjust coords for array (start at 0)
		mines = append(mines, mine{x: in[i] - 1, y: in[i+1] - 1})
	}
	return mines
}

func buildMineField(rows, cols int) [][]byte {
	mf := make([][]byte, rows, rows)
	for r := 0; r < rows; r++ {
		col := make([]byte, cols, cols)
		mf[r] = col
		for c := 0; c < cols; c++ {
			mf[r][c] = '0'
		}
	}
	return mf
}

func incByte(el byte) byte {
	return el - '0' + '1'
}

func calcScore(mf *[][]byte, r, c int) {
	// по часовой стрелке начинаем проверять, есть ли рядом бомбы
	if r-1 >= 0 && (*mf)[r-1][c] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if r-1 >= 0 && c+1 < len((*mf)[r-1]) && (*mf)[r-1][c+1] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if c+1 < len((*mf)[r]) && (*mf)[r][c+1] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if r+1 < len(*mf) && c+1 < len((*mf)[r+1]) && (*mf)[r+1][c+1] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if r+1 < len(*mf) && (*mf)[r+1][c] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if r+1 < len(*mf) && c-1 >= 0 && (*mf)[r+1][c-1] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if c-1 >= 0 && (*mf)[r][c-1] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
	if r-1 >= 0 && c-1 >= 0 && (*mf)[r-1][c-1] == '*' {
		(*mf)[r][c] = incByte((*mf)[r][c])
	}
}

func enrichMineField(mf *[][]byte, mines []mine) *[][]byte {
	for _, mine := range mines {
		(*mf)[mine.x][mine.y] = '*'
	}
	// calc non-bombs fields
	for r := 0; r < len(*mf); r++ {
		for c := 0; c < len((*mf)[r]); c++ {
			if (*mf)[r][c] == '*' {
				continue
			}
			calcScore(mf, r, c)
		}

	}
	return mf
}

func printMineField(mf *[][]byte) {
	for _, rows := range *mf {
		for j, col := range rows {
			if j < len(rows)-1 {
				fmt.Printf("%s ", string(col))
			} else {
				fmt.Print(string(col))
			}
		}
		fmt.Println()
	}
}

func main() {
	input := convertStrArrToIntArr(getInputFromFileInLines("input.txt"))
	rows, columns := input[0], input[1]
	mines := parseMinesCoords(input[3:])
	mf := buildMineField(rows, columns)
	printMineField(enrichMineField(&mf, mines))
}
