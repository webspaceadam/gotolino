package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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
		case "bookmark":
			tmpNoteType = "bookmark"
			tmpSiteInformation = getSiteInformation(scanner.Text())
		case "delimeter":
			//fmt.Println("Starting new note! 🏃🏼")
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

func createNoteMd(note Note) string {
	writtenNote := ""

	if note.note != "tmp" {
		writtenNote = note.note
	}

	noteMD := fmt.Sprintf(`
- %s %s
  type:: note
  book:: [[Book - %s]]
  author:: [[%s]] 
  relates:: #Za:area, #Zb:book, #Zc:Chapter
  - > %s
`, note.siteInformation, writtenNote, note.book, note.author, note.marked)

	return noteMD
}

func getNote(text string) string {
	note := text[strings.Index(text, ":") + 2:]
	return note
}

func getSiteInformation(text string) string {
	siteInformation := text[0:strings.Index(text, ":")]
	return siteInformation
}

func getTitleAndAuthorOrMarking(text string) (string, string) {
	isAddingInformation := strings.Contains(text, "Hinzugefügt am")
	isQuote := strings.Contains(text, `"`)
	title := ""
	author := ""

	if !isAddingInformation && !isQuote {
		re := regexp.MustCompile(`(?m)\((.*?)\)`)
		str := text

		foundStrings := re.FindAllString(str, -1)

		author = foundStrings[len(foundStrings) - 1]
		author = author[1 : len(author) - 1]
		authorSplitted := strings.Split(author, ", ")
		author = authorSplitted[1] + " " + authorSplitted[0]
		firstParentheses := strings.Index(text, "(")
		title = text[0:firstParentheses]
	} else {
		return text, "marking"
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
	} else if strings.Contains(text, ADDED) {
		return "added"
	} else if len(text) == 0 {
		return "empty"
	}
	return "other"
}
