package todo

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

type Todo struct{
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Must enter input")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo *Todo) Display(){
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := strings.ReplaceAll(todo.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}