package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"capm-exam-system/internal/database"
	"capm-exam-system/internal/repository"
)

func main() {
	exam := flag.String("exam", "capm", "Exam to seed: capm or pmp")
	flag.Parse()

	// Connect to database
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := repository.New(db)
	ctx := context.Background()

	var (
		questions  []QuestionData
		additional []QuestionData
		messages   []string
	)

	switch strings.ToLower(*exam) {
	case "capm":
		questions = GetQuestions()
		additional = GetAdditionalQuestions()
		messages = append(messages, "Starting to seed database with authentic CAPM questions based on PMBOK 7th Edition and PMI ECO 2024...")
	case "pmp":
		questions = GetPMPQuestions()
		additional = GetAdditionalPMPQuestions()
		messages = append(messages, "Starting to seed database with PMP-aligned questions based on PMI ECO 2021 and current standards...")
	default:
		log.Fatalf("Unsupported exam: %s", *exam)
	}

	for _, msg := range messages {
		log.Println(msg)
	}

	for i, q := range questions {
		// Create question
		question, err := repo.CreateQuestion(ctx, q.prompt, q.domain, q.explanation, q.popularityScore, q.isMultiSelect)
		if err != nil {
			log.Fatalf("Failed to create question %d: %v", i+1, err)
		}

		// Create choices
		for _, c := range q.choices {
			_, err := repo.CreateChoice(ctx, question.ID, c.text, c.label, c.isCorrect)
			if err != nil {
				log.Fatalf("Failed to create choice for question %d: %v", i+1, err)
			}
		}

		log.Printf("Created question %d: %s", i+1, question.Prompt[:min(80, len(question.Prompt))]+"...")
	}

	for i, q := range additional {
		question, err := repo.CreateQuestion(ctx, q.prompt, q.domain, q.explanation, q.popularityScore, q.isMultiSelect)
		if err != nil {
			log.Fatalf("Failed to create additional question %d: %v", i+1, err)
		}

		for _, c := range q.choices {
			_, err := repo.CreateChoice(ctx, question.ID, c.text, c.label, c.isCorrect)
			if err != nil {
				log.Fatalf("Failed to create choice for additional question %d: %v", i+1, err)
			}
		}
	}

	log.Printf("Database seeding completed successfully! Created %d questions for the %s exam.", len(questions)+len(additional), strings.ToUpper(*exam))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
