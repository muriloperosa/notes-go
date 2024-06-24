package note

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/muriloperosa/notes-go/output"
	"github.com/olekukonko/tablewriter"
)

const STORAGE_PATH = "storage/"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	fileName  string
}

func (note Note) Show() {
	output.TextLn("Title: " + note.Title)
	output.TextLn("Content: " + note.Content)
}

func (note Note) ShortTitle(maxLen int) string {
	if len(note.Title) > maxLen {
		return note.Title[:maxLen] + "..."
	}

	return note.Title
}

func (note Note) Save() error {
	fileName := STORAGE_PATH + note.fileName

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
		fileName:  time.Now().Format("20060102150405") + ".json",
	}, nil
}

func NewFromFile(fileName string) (Note, error) {

	jsonFile, err := os.Open(STORAGE_PATH + fileName)
	if err != nil {
		return Note{}, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Note{}, err
	}

	if len(byteValue) == 0 {
		return Note{}, errors.New("file is empty")
	}

	var note Note

	err = json.Unmarshal(byteValue, &note)
	if err != nil {
		return Note{}, err
	}

	note.fileName = fileName

	return note, nil
}

func GetAll() ([]Note, error) {
	files, err := os.ReadDir(STORAGE_PATH)
	if err != nil {
		return []Note{}, err
	}

	notes := []Note{}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			note, err := NewFromFile(file.Name())

			if err != nil {
				return []Note{}, err
			}

			notes = append(notes, note)
		}
	}

	return notes, nil
}

func ListAll(notes []Note) error {

	if len(notes) == 0 {
		return errors.New("no notes found")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Title", "File Name"})
	table.SetBorder(true)

	for i := 0; i < len(notes); i++ {
		table.Append([]string{strconv.Itoa(i + 1), notes[i].ShortTitle(20), notes[i].fileName})
	}

	table.Render()

	return nil
}
