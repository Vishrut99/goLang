package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

func (n Note) Display() {
	println("Title:", n.Title)
	println("Content:", n.Content)
	println("Created At:", n.CreatedAt.String())
}

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return errors.New("Json error")
	}

	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("error")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
