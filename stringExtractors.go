package main

import (
	"regexp"
	"strings"
)

func getMarking(text string) string {
	marking := text[strings.Index(text, ":"): len(text)]
	return marking
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
	isQuote := strings.Contains(text, `"`) || strings.Contains(text,"»") || strings.Contains(text, "«")
	title := ""
	author := ""

	if !isAddingInformation && !isQuote && checkAuthorString(text) {
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

func checkAuthorString(text string) bool {
	re := regexp.MustCompile(`(?m)\((.*?)\)`)
	str := text

	foundStrings := re.FindAllString(str, -1)

	isAuthor := len(foundStrings) > 0

	return isAuthor
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
	} else if strings.Contains(text, ADDED) || strings.Contains(text, CHANGED) {
		return "added"
	} else if len(text) == 0 {
		return "empty"
	}
	return "other"
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getNotesSortedByBooks(books []string, notes Notes) []Notes {
	sortedNotes := make([]Notes, 0)

	for _ = range books {
		emptyNote := make([]Note, 0)
		sortedNotes = append(sortedNotes, emptyNote)
	}

	for _, note := range notes {
		bookIndex := IndexOf(books, note.book)
		sortedNotes[bookIndex] = append(sortedNotes[bookIndex], note)
	}

	return sortedNotes
}

func getStringSeperated(text string, seperator string) string {
	return text[strings.Index(text, seperator) + 1: len(text)]
}
