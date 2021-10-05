package main

const NOTE = "Notiz"
const MARKING = "Markierung"
const BOOKMARK = "Lesezeichen"
const DELIMETER = "-----------------------------------"
const ADDED = "Hinzugefügt am"
const CHANGED = "Geändert am"

type Note struct {
	author string
	book string
	notetype string
	marked string
	note string
	siteInformation string
}

type Notes []Note

func IndexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
