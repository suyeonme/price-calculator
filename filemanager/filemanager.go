package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
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

	// Read file
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("could not read file")
	}

	file.Close()
	return lines, nil
}

func (fm Filemanager) WriteResult(data interface{}) error {
	// Create file
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("could not create file")
	}

	// Write data to file as JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("could not convert data to JSON")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
