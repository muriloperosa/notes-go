package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/muriloperosa/notes-go/output"
)

const STORAGE_PATH = "storage/"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Show() {
	output.TextLn("Title: " + note.Title)
	output.TextLn("Content: " + note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = STORAGE_PATH + strings.ToLower(fileName) + ".json"

	content, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, content, 0644)
}

func New(title string, content string) (*Note, error) {

	if title == "" {
		return &Note{}, errors.New("invalid title for the note")
	}

	if content == "" {
		return &Note{}, errors.New("invalid content for the note")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
