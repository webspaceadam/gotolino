package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	settingNotes, settingsPath := readSettings()

	fmt.Println("Settings Notes", settingNotes, "settingsPath", settingsPath)

	notes, err := os.Open(settingNotes)
	check(err)
	defer notes.Close()

	filteredNotes := make(Notes, 0)
	books := make([]string, 0)

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
	fmt.Println("Starting to collect the notes 🏃🏼")

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
			fmt.Println(scanner.Text())
			titleOrMarking, authorOrMarkingFlag := getTitleAndAuthorOrMarking(scanner.Text())

			if authorOrMarkingFlag != "marking" {
				tmpBook = titleOrMarking
				tmpAuthor = authorOrMarkingFlag

				// Add Book to booklist
				if(!contains(books, tmpBook)) {
					books = append(books, tmpBook)
				}
			} else {
				tmpMarked = titleOrMarking
			}
		}
	}

	sortedNotesByBook := getNotesSortedByBooks(books, filteredNotes)

	for _, book := range sortedNotesByBook {
		// create folder code
		bookNamePath := settingsPath + "/" + book[0].author + " - " + book[0].book
		fmt.Println("Book being looped:",  bookNamePath)


		if _, err := os.Stat(bookNamePath); os.IsNotExist(err) {
			fmt.Println("Creating directory for:", bookNamePath)
			err := os.Mkdir(bookNamePath, 0755)
			check(err)
		} else {
			fmt.Println("Path already exists for: ", bookNamePath)
		}

		for _, note := range book {
			noteFileName := note.siteInformation + ".md"
			noteString := createNoteMd(note)
			notePath := bookNamePath + "/" + noteFileName

			if _, err := os.Stat(notePath); os.IsNotExist(err) {
				file, err := os.Create(notePath)
				defer file.Close()
				check(err)

				_, err2 := file.WriteString(noteString)
				check(err2)
			} else {
				fmt.Println("Path already exists for: ", notePath)
			}
		}

		fmt.Println("✅ Done creating notes in: ", settingsPath)
	}
}

// readSettings reads the settings.txt in the current directory and returns the two locations needed for the programm
func readSettings() (string, string) {
	currentScriptDirectory := GetCurrentScriptDirectory();
	var notesTxtLocation, notesSavingPath string

	settings, err := os.Open(currentScriptDirectory + "/settings.txt")
	check(err)
	defer settings.Close()

	scanner := bufio.NewScanner(settings)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "notesTxtLocation") {
			notesTxtLocation = getStringSeperated(scanner.Text(), "=")
		} else if strings.Contains(scanner.Text(), "notesSavingPath") {
			notesSavingPath = getStringSeperated(scanner.Text(), "=")
		}
	}

	return notesTxtLocation, notesSavingPath
}

// GetCurrentScriptDirectory returns the directory of the currently running go script file.
func GetCurrentScriptDirectory() string {
	// NOTE: Replace the 1 with a 0 if you use this code directly, instead of wrapping it in a function.
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "../")
}
