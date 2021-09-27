package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Note struct {
	author string
	book string
	notetype string
	marked string
	note string
}

type NoteTypes struct {
	note string
	marking string
	bookmark string
}

func main() {
	notes, err := os.Open("test.txt")
	check(err)
	defer notes.Close()

	scanner := bufio.NewScanner(notes)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		switch checkNoteType(scanner.Text()) {
		case "note":
			fmt.Println("It is a note")
		case "marking":
			fmt.Println("It is a marking")
		case "bookmark":
			fmt.Println("it is a bookmark")
		case "delimeter":
			fmt.Println("Delimeter")
		case "empty":
			fmt.Println("Empty line")
		case "other":
			fmt.Println("String does not point to a note type, probably booktitle with author or a marking")
		}
	}
}

func checkNoteType(text string) string {
	notetypes := NoteTypes{
		note: "Notiz",
		marking: "Markierung",
		bookmark: "Lesezeichen",
	}
	delimeter := "-----------------------------------"

	if strings.Contains(text, notetypes.note) {
		return "note"
	} else if strings.Contains(text, notetypes.marking) {
		return "marking"
	} else if strings.Contains(text, notetypes.bookmark) {
		return "bookmark"
	} else if strings.Contains(text, delimeter) {
		return "delimeter"
	} else if len(text) == 0 {
		return "empty"
	}
	return "other"
}
