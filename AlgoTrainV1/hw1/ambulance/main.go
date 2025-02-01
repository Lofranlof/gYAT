package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const Z_LEN = 1000001

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

// калибрует количество квартир на этаже
// с помощью входных данных
func findZs(m, k2, p2, n2 int) []int {
	ans := make([]int, 0, Z_LEN)
	for z := 1; z < Z_LEN; z++ {
		// строим небоскрёб
		n2New := m*(p2-1) + n2
		maxApartNumber := n2New * z
		minApartNumber := (n2New-1)*z + 1
		if k2 <= maxApartNumber && k2 >= minApartNumber {
			ans = append(ans, z)
		}
	}
	return ans
}

func findP(k, m, z int) int {
	return int(math.Ceil(math.Ceil(float64(k)/float64(z)) / float64(m)))
}

func findN(k, m, z int) int {
	return (((k - 1) / z) % m) + 1
}

func findAmbigousParameter(zs []int, k1, m int) (p1 int, n1 int) {
	n1s := make(map[int]struct{}, len(zs))
	p1s := make(map[int]struct{}, len(zs))

	for _, z := range zs {
		n := findN(k1, m, z)
		p := findP(k1, m, z)
		n1s[n], p1s[p] = struct{}{}, struct{}{}
	}

	if len(n1s) == 1 {
		for n := range n1s {
			n1 = n
		}
	}

	if len(p1s) == 1 {
		for p := range p1s {
			p1 = p
		}
	}
	return p1, n1
}

func main() {
	input := convertStrArrToIntArr(strings.Split(getInputFromFileInLines("input.txt")[0], " "))
	k1, m, k2, p2, n2 := input[0], input[1], input[2], input[3], input[4]
	if n2 > m {
		fmt.Printf("%d %d\n", -1, -1)
		return
	}
	zs := findZs(m, k2, p2, n2)

	if len(zs) == 0 {
		fmt.Printf("%d %d\n", -1, -1)
	} else if len(zs) == 1 {
		p1 := findP(k1, m, zs[0])
		n1 := findN(k1, m, zs[0])
		fmt.Printf("%d %d\n", p1, n1)
		return
	} else if len(zs) > 1 {
		p1, n1 := findAmbigousParameter(zs, k1, m)
		fmt.Printf("%d %d\n", p1, n1)
	}
}
