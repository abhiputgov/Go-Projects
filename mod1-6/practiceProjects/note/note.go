package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}

func (note Note) DisplayNote() {
	fmt.Printf("your note titled '%v' has the following content:\n %v", note.Title, note.Content)
}

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileNameFinal := strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileNameFinal, json, 0644) // will be nil if there is no error. and err is generated if something goes wrong
}
