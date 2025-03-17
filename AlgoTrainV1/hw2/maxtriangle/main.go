package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type tuple struct {
	freq     float64
	feedback string
}

func createTuple(freq, feedback string) tuple {
	freqInt, err := strconv.ParseFloat(strings.TrimSpace(freq), 64)
	if err != nil {
		panic(err)
	}

	return tuple{
		freq:     float64(freqInt),
		feedback: feedback,
	}
}

// долбоёбы не могут дать мне провалидированный инпут,
// поэтому пришлось наговнокодить
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

func getFreqsAndFeedback(in []string) []tuple {
	tuples := make([]tuple, 0, len(in)/2)
	tuples = append(tuples, createTuple(in[0], ""))

	for i := 1; i < len(in)-1; i += 2 {
		tuples = append(tuples, createTuple(in[i], in[i+1]))
	}

	return tuples
}

func calcTriangleRange(tuningInfo []tuple) (minF, maxF float64) {
	minF, maxF = 30, 4000
	lastFreq := tuningInfo[0].freq
	for _, t := range tuningInfo[1:] {
		if t.freq < lastFreq {
			if t.feedback == "closer" {
				maxF = min(maxF, lastFreq-((lastFreq-t.freq)/2))
			} else if t.feedback == "further" {
				minF = max(minF, t.freq+((lastFreq-t.freq)/2))
			}
		} else if t.freq > lastFreq {
			if t.feedback == "closer" {
				minF = max(minF, lastFreq+((t.freq-lastFreq)/2))
			} else if t.feedback == "further" {
				maxF = min(maxF, t.freq-((t.freq-lastFreq)/2))
			}
		}
		lastFreq = t.freq
	}
	return minF, maxF
}

func main() {
	input := getInputFromFileInLines("input.txt")[1:]
	tuning := getFreqsAndFeedback(input)

	minF, maxF := calcTriangleRange(tuning)

	fmt.Printf("%.6f %.6f\n", minF, maxF)
}
