package main

import (
	"GoNote/note"
	"fmt"
)

func testWrite() {
	testNote := note.Note{
		Title:   "Cat",
		Content: "Cats are a type of cat.",
	}
	err := note.SaveNote(testNote, fmt.Sprintf("./storage/%s.json", testNote.Title))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func testRead() {
	testNote, err := note.LoadNote("./storage/Cat.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Loaded Note:\n%+v\n", testNote)
}

func main() {
	testWrite()
	testRead()
}
