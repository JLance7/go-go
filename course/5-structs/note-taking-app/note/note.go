package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
)

type Note struct{
	Title string `json:"title"` // struct tag (metadata for json package)
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Must enter input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) PrintNoteData(){
	fmt.Println(note.Title)
	fmt.Println(note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}