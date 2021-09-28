package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
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

func main() {
	notes, err := os.Open("test.txt")
	check(err)
	defer notes.Close()

	filteredNotes := make([]Note, 0)

	//fmt.Println("filtered notes", filteredNotes)

	scanner := bufio.NewScanner(notes)
	for scanner.Scan() {
		newNote := Note{}
		//fmt.Println(scanner.Text())
		switch checkNoteType(scanner.Text()) {
		case "note":
			//fmt.Println("It is a note")
		case "marking":
			//fmt.Println("It is a marking")
		case "bookmark":
			//fmt.Println("it is a bookmark")
		case "delimeter":
			//fmt.Println("Delimeter")
		case "empty":
			//fmt.Println("Empty line")
		case "other":
			//fmt.Println("String does not point to a note type, probably booktitle with author or a marking")
			title, author := getTitleAndAuthor(scanner.Text())

			fmt.Println("Title: ", title, "author", author)

			newNote.book = title
			newNote.author = author
		}

		filteredNotes = append(filteredNotes, newNote)
	}

	//fmt.Println(filteredNotes);
}

func getTitleAndAuthor(text string) (string, string) {
	isAddingInformation := strings.Contains(text, "Hinzugefügt am")
	isQuote := strings.Contains(text, `"`)
	title := ""
	author := ""

	if !isAddingInformation && !isQuote {
		re := regexp.MustCompile(`(?m)\((.*?)\)`)
		str := text

		foundStrings := re.FindAllString(str, -1)

		author = foundStrings[len(foundStrings) - 1]
		firstParentheses := strings.Index(text, "(")
		title = text[0:firstParentheses]
	}

	return title, author
}

func checkNoteType(text string) string {
	if strings.Contains(text, NOTE) {
		return "note"
	} else if strings.Contains(text, MARKING) {
		return "marking"
	} else if strings.Contains(text, BOOKMARK) {
		return "bookmark"
	} else if strings.Contains(text, DELIMETER) {
		return "delimeter"
	} else if len(text) == 0 {
		return "empty"
	}
	return "other"
}
