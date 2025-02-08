package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func New(inputPath string, outputPath string) *FileManager{
	return &FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("An error occured")
		// file.Close()
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("An error occured")
		// file.Close()
		return nil, err
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error{
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file")
	}
	defer file.Close()

	time.Sleep(time.Second * 3)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// file.Close()
		return errors.New("Failed to convert json")
	}
	// file.Close()
	return nil
}
