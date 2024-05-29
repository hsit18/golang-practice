package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Title string `json:"title"`
}

func (todo Todo) Display() {
	fmt.Printf("Your note titled %v", todo.Title)
}

func (todo Todo) Save() error {
	fileName := strings.ReplaceAll(todo.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title string) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Title: title,
	}, nil
}
