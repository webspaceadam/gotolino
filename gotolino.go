package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	notes, err := os.Open("test.txt")
	check(err)
	defer notes.Close()

	filteredNotes := make([]Note, 0)

	scanner := bufio.NewScanner(notes)

	currentlyCreatingNewNote := true

	// tmpSavings
	tmpAuthor := "tmp"
	tmpBook := "tmp"
	tmpNote := "tmp"
	tmpMarked := "tmp"
	tmpNoteType := "tmp"
	tmpSiteInformation := "tmp"

	noteCount := 0
	fmt.Println("Starting new note! 🏃🏼")

	for scanner.Scan() {
		switch checkNoteType(scanner.Text()) {
		case "note":
			tmpNote = getNote(scanner.Text())
			tmpSiteInformation = getSiteInformation(scanner.Text())
			tmpNoteType = "note"
		case "marking":
			tmpNoteType = "marking"
			tmpSiteInformation = getSiteInformation(scanner.Text())
			tmpMarked = getMarking(scanner.Text())
		case "bookmark":
			tmpNoteType = "bookmark"
			tmpSiteInformation = getSiteInformation(scanner.Text())
		case "delimeter":
			noteCount += 1
			currentlyCreatingNewNote = true
		case "empty":
		case "added":
			if currentlyCreatingNewNote {
				newNote := Note{
					author:          tmpAuthor,
					book:     	     tmpBook,
					notetype: 	     tmpNoteType,
					marked:   	     tmpMarked,
					note:     	     tmpNote,
					siteInformation: tmpSiteInformation,
				}

				filteredNotes = append(filteredNotes, newNote)
				currentlyCreatingNewNote = false

				fmt.Println("Ending Note 🙅🏼‍♂️")

				// Reset Tmp Fields
				tmpNote = "tmp"
				tmpAuthor = "tmp"
				tmpBook = "tmp"
				tmpAuthor = "tmp"
				tmpMarked = "tmp"
				tmpNoteType = "tmp"
				tmpSiteInformation = "tmp"
			}
		case "other":
			titleOrMarking, authorOrMarkingFlag := getTitleAndAuthorOrMarking(scanner.Text())

			if authorOrMarkingFlag != "marking" {
				tmpBook = titleOrMarking
				tmpAuthor = authorOrMarkingFlag
			} else {
				tmpMarked = titleOrMarking
			}
		}
	}

	for i:=0;i<len(filteredNotes);i++{
		fmt.Println(createNoteMd(filteredNotes[i]))
	}
}
