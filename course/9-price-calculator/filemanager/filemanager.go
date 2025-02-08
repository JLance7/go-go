package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"encoding/json"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("An error occured")
		file.Close()
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("An error occured")
		file.Close()
		return nil, err
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error{
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert json")
	}
	file.Close()
	return nil
}

func New(inputPath string, outputPath string) *FileManager{
	return &FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}