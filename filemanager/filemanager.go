package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadLines() ([]string, error) {
	// Open file
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("could not open file")
	}

	// defer: excute when function ends
	defer file.Close()

	// Read file
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("could not read file")
	}

	return lines, nil
}

func (fm Filemanager) WriteResult(data interface{}) error {
	// Create file
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("could not create file")
	}

	// defer: excute when function ends
	defer file.Close()

	// Simulate long operation
	time.Sleep(3 * time.Second)

	// Write data to file as JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("could not convert data to JSON")
	}
	return nil
}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
