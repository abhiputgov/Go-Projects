package file_manager

import (
	"bufio"
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
