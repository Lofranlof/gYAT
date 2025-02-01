package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	NoSolution = iota
	OneSolution
	ManySolOneFreeVar
	ManySolTwoFreeVar
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
		line := strings.Trim(scanner.Text(), " ")
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func convertStrArrToFloatArr(in []string) []float64 {
	out := make([]float64, 0, len(in))
	for _, el := range in {
		elFloat, err := strconv.ParseFloat(strings.Trim(el, " "), 64)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, elFloat)
	}
	return out
}

// matrix of a form:
// a b
// c d
func findRank(aTopLeft, bTopRight, cBottomLeft, dBottomRight float64) int {
	if aTopLeft == 0 && bTopRight == 0 && cBottomLeft == 0 && dBottomRight == 0 {
		return 0
	}
	det := (aTopLeft * dBottomRight) - (bTopRight * cBottomLeft)
	if det == 0 {
		return 1
	}
	return 2
}

// 0 - inconsistent, no solution
// 1 - consistent, 1 solution
// 2 - consistent, 1 free variable
// 3 - consistent, 2 free variables
func isConsistent(a, b, c, d, e, f float64) int {
	rankCoeffMatrix := findRank(a, b, c, d)
	rankAugMatrix := max(findRank(a, e, c, f), findRank(b, e, d, f))
	if rankCoeffMatrix != rankAugMatrix {
		return NoSolution
	}
	switch rankCoeffMatrix {
	case 2:
		return OneSolution
	case 1:
		return ManySolOneFreeVar
	}
	return ManySolTwoFreeVar
}

// Augmented matrix of a form:
// a b | e
// c d | f
func KramerMethod(a, b, c, d, e, f float64) []float64 {
	ans := make([]float64, 0, 2)
	genDet := a*d - b*c
	detX := e*d - b*f
	detY := a*f - e*c
	ans = append(ans, detX/genDet, detY/genDet)
	return ans
}

// we get only matrices with rank 1
func GaussMethod(a, b, c, d, e, f float64) (int, []float64) {
	ans := make([]float64, 0, 2)
	var mode int = 1
	if a == 0 && c == 0 {
		mode = 4
		if b != 0 {
			ans = append(ans, e/b)
		} else {
			ans = append(ans, f/d)
		}
	} else if b == 0 && d == 0 {
		mode = 3
		if a != 0 {
			ans = append(ans, e/a)
		} else {
			ans = append(ans, f/c)
		}
	} else {
		if b != 0 {
			ans = append(ans, -a/b, e/b)
		} else {
			ans = append(ans, -c/d, f/d)
		}
	}
	return mode, ans

}

//	{
//		ax + by = e
//		cx + dy = f
//	}
func solveSLAU(a, b, c, d, e, f float64) (int, []float64) {
	consistency := isConsistent(a, b, c, d, e, f)
	switch consistency {
	case NoSolution:
		return 0, []float64{}
	case ManySolTwoFreeVar:
		return 5, []float64{}
	case OneSolution:
		return 2, KramerMethod(a, b, c, d, e, f)
	}
	return GaussMethod(a, b, c, d, e, f)
}

func main() {
	input := convertStrArrToFloatArr(getInputFromFileInLines("input.txt"))
	printAns(solveSLAU(input[0], input[1], input[2], input[3], input[4], input[5]))
}

func printAns(mode int, ans []float64) {
	fmt.Printf("%d", mode)
	for _, num := range ans {
		fmt.Printf(" %.10f", num)

	}
	fmt.Println()
}
