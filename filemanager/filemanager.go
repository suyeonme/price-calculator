package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("could not open file")
	}

	// Read file
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Could not read file")
	}

	file.Close()
	return lines, nil
}
