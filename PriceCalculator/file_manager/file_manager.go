package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
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
		file.Close()
		return nil, errFromScanner
	}
	return lines, nil
}

func WriteJSON(path string, data any) error {
	createdFile, errCreatingFile := os.Create(path)
	if errCreatingFile != nil {
		return errors.New("Failed to create file.")
	}
	encoder := json.NewEncoder(createdFile)
	errorEncoding := encoder.Encode(data)
	if errorEncoding != nil {
		createdFile.Close()
		return errors.New("Failed to convert data to JSON.")
	}
	createdFile.Close()
	return nil
}
