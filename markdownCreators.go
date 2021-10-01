package main

import "fmt"

func createNoteMd(note Note) string {
	writtenNote := ""

	if note.note != "tmp" {
		writtenNote = ": " + note.note
	}

	if note.notetype == "note" {
		noteMD := fmt.Sprintf(`
- %s %s
  type:: note
  book:: [[Book - %s]]
  author:: [[%s]] 
  relates:: #Za:area, #Zb:book, #Zc:Chapter
  - > %s
`, note.siteInformation, writtenNote, note.book, note.author, note.marked)

		return noteMD
	} else if note.notetype == "marking" {
		noteMD := fmt.Sprintf(`
- %s
  type:: note
  book:: [[Book - %s]]
  author:: [[%s]] 
  relates:: #Za:area, #Zb:book, #Zc:Chapter
  - > %s
`, note.siteInformation, note.book, note.author, note.marked)

		return noteMD
	} else {
		noteMD := fmt.Sprintf(`
- %s
  type:: note
  book:: [[Book - %s]]
  author:: [[%s]] 
  relates:: #Za:area, #Zb:book, #Zc:Chapter
`, note.siteInformation, note.book, note.author)

		return noteMD
	}
}
