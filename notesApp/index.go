package notesApp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hsit18/golang-practice/notesApp/todo"
)

type outputable interface {
	Display()
	Save() error
}

func Execute() {
	title, content := getNoteData()

	userNote, err := New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputter(userNote)
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")

	todoTitle := getTodoData()

	userTodo, err := todo.New(todoTitle)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputter(userTodo)
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Saving the todo succeeded!")
}

func getTodoData() string {
	title := getUserInput("Note title:")
	return title
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func outputter(o outputable) error {
	fmt.Println("This is the outputter function")

	o.Display()
	err := o.Save()
	return err
}
