package main

// QuestionData represents the structure for question data
type QuestionData struct {
	prompt          string
	domain          string
	explanation     string
	popularityScore float64
	isMultiSelect   bool
	choices         []ChoiceData
}

type ChoiceData struct {
	text      string
	label     string
	isCorrect bool
}

// QuestionData and ChoiceData are shared between seed data files.
