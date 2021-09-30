package main

const NOTE = "Notiz"
const MARKING = "Markierung"
const BOOKMARK = "Lesezeichen"
const DELIMETER = "-----------------------------------"
const ADDED = "Hinzugefügt am"

type Note struct {
	author string
	book string
	notetype string
	marked string
	note string
	siteInformation string
}
