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
