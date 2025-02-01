package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type phoneBook struct {
	phoneNumbers []*phoneNumber
}

type phoneNumber struct {
	fullNumber []int
}

func buildPhoneBook() *phoneBook {
	pb := &phoneBook{}
	pb.phoneNumbers = make([]*phoneNumber, 0, 3)
	return pb
}

func (pb *phoneBook) addPhoneNumber(phNum *phoneNumber) {
	pb.phoneNumbers = append(pb.phoneNumbers, phNum)
}

func parsePhoneNumber(input string) *phoneNumber {
	pn := &phoneNumber{}
	pn.fullNumber = make([]int, 0, 11)

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] < 48 || input[i] > 57 {
			continue
		}
		pn.fullNumber = append(pn.fullNumber, int(byte(input[i])-48))
	}

	if len(pn.fullNumber) < 11 {
		pn.fullNumber = append(pn.fullNumber, []int{5, 9, 4, 7}...)
	}
	return pn
}

func areEqualNumbers(pn1, pn2 *phoneNumber) bool {
	for i := 0; i < min(len(pn1.fullNumber), len(pn2.fullNumber))-1; i++ {
		if pn1.fullNumber[i] != pn2.fullNumber[i] {
			return false
		}
	}
	return true
}

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
	pb := buildPhoneBook()
	input, err := getInputFromFileInLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pnToAdd := parsePhoneNumber(input[0])

	for i := 1; i < len(input); i++ {
		pn := parsePhoneNumber(input[i])
		pb.addPhoneNumber(pn)
	}

	for _, pn := range pb.phoneNumbers {
		if areEqualNumbers(pn, pnToAdd) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
