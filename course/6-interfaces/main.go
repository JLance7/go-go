package main

import (
	"bufio"
	"fmt"
	"interfaces-example/note"
	"interfaces-example/todo"
	"os"
	"strings"
)

type saver interface { // struct that has method Save()
	Save() error
}

type displayer interface {
	Display()
	saver
}

func main(){
	// note stuff
	title, content := getNoteData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(&note)
	if err != nil {
		return
	}
	

	// todo stuff
	todoText := getUserInput("Todo Text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		return
	}

	err = outputData(&todo)
	if err != nil {
		return
	}

	addedVal := add(1, 2)
	fmt.Println(addedVal + 1)
}

// func printSomething(value interface{}){
// 	fmt.Println(value)
// }

func printSomething(value any){
	fmt.Println(value)

	// typedVal, ok := value.(int)
	// if !ok {

	// }

	switch value.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}

func getTodoData(prompt string) string {
	text := getUserInput(prompt)
	return text
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

func outputData(data displayer) error {
	data.Display()
	err := saveData(data)
	return err
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving note failed")
		return err
	}
	fmt.Println("Saving note succeeded")
	return nil
}

// generic
func add[T int | float64 | string](a, b T) T {
	return a + b
}