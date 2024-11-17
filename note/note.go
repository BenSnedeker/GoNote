package note

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func SaveNote(note Note, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed to craete file: %v", err)
	}
	defer file.Close()

	// Serialize to JSON and write to file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(note)
	if err != nil {
		return fmt.Errorf("Failed to encode to JSON: %v", err)
	}

	fmt.Println("Note saved to file successfully.")
	return nil
}

func LoadNote(filename string) (Note, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return Note{}, fmt.Errorf("Failed to read file: %v", err)
	}

	// Deserialize JSON
	var note Note
	err = json.Unmarshal(fileContent, &note)
	if err != nil {
		return Note{}, fmt.Errorf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Println("Note loaded successfully.")
	return note, nil
}
