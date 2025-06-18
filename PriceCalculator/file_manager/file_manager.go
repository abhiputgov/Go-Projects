package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath string, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLinesFromFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errFromScanner := scanner.Err()
	if errFromScanner != nil {
		return nil, errFromScanner
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	createdFile, errCreatingFile := os.Create(fm.OutputFilePath)
	if errCreatingFile != nil {
		return errors.New("Failed to create file.")
	}
	defer createdFile.Close()
	encoder := json.NewEncoder(createdFile)
	errorEncoding := encoder.Encode(data)
	if errorEncoding != nil {
		//createdFile.Close()
		return errors.New("Failed to convert data to JSON.")
	}
	//createdFile.Close()
	return nil
}
