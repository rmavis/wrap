package main

type CharacterKey int
const (
	SetSplit = iota
	SetJoin
	SetWrapOpen
	SetWrapClose
	RecordSplit
	RecordJoin
	RecordWrapOpen
	RecordWrapClose
	FieldSplit
	FieldJoin
	FieldWrapOpen
	FieldWrapClose
)

