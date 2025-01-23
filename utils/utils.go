package utils

import (
	"bufio"
	"io"
	"os"
)

func GetInputFromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

func GetInputFromFileInLines(path string) ([]string, error) {
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

func WriteAnsToFile(path, ans string) error {
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
