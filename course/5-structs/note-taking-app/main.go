package main

import (
	"bufio"
	"fmt"
	"note-taker/note"
	"os"
	"strings"
)

func main(){
	title, content := getNoteData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.PrintNoteData()
	err = note.Save()
	if err != nil {
		fmt.Println("Saving note failed")
		return
	}
	fmt.Println("Saving note succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	// var value string
	// fmt.Scanln(&value) // doesn't work well with multiple spaces in string

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")


	return text
}

