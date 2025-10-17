package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

type Question struct {
	ID              int     `json:"id"`
	Prompt          string  `json:"prompt"`
	Domain          string  `json:"domain"`
	PopularityScore float64 `json:"popularity_score"`
	Explanation     string  `json:"explanation"`
	IsMultiSelect   bool    `json:"is_multi_select"`
}

type Choice struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
	Label      string `json:"label"` // A, B, C, D
}

type QuestionWithChoices struct {
	Question
	Choices []Choice `json:"choices"`
}

type Exam struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Attempt struct {
	ID        uuid.UUID  `json:"id"`
	ExamID    uuid.UUID  `json:"exam_id"`
	UserID    uuid.UUID  `json:"user_id"`
	Seed      int64      `json:"seed"`
	Score     *int       `json:"score,omitempty"`
	MaxScore  int        `json:"max_score"`
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}

type AttemptAnswer struct {
	ID         uuid.UUID `json:"id"`
	AttemptID  uuid.UUID `json:"attempt_id"`
	QuestionID int       `json:"question_id"`
	ChoiceID   *int      `json:"choice_id,omitempty"`
	IsCorrect  *bool     `json:"is_correct,omitempty"`
}

type ExamQuestion struct {
	ID         uuid.UUID `json:"id"`
	ExamID     uuid.UUID `json:"exam_id"`
	QuestionID int       `json:"question_id"`
	Position   int       `json:"position"`
}

type ExamSubmission struct {
	Answers []AnswerSubmission `json:"answers"`
}

type AnswerSubmission struct {
	QuestionID int   `json:"question_id"`
	ChoiceIDs  []int `json:"choice_ids"`
}

type ExamResult struct {
	AttemptID uuid.UUID        `json:"attempt_id"`
	UserID    uuid.UUID        `json:"user_id"`
	ExamID    uuid.UUID        `json:"exam_id"`
	Score     int              `json:"score"`
	MaxScore  int              `json:"max_score"`
	StartedAt time.Time        `json:"started_at"`
	EndedAt   *time.Time       `json:"ended_at,omitempty"`
	Results   []QuestionResult `json:"results"`
}

type AttemptHistory struct {
	AttemptID     uuid.UUID  `json:"attempt_id"`
	ExamID        uuid.UUID  `json:"exam_id"`
	ExamName      string     `json:"exam_name"`
	StartedAt     time.Time  `json:"started_at"`
	EndedAt       *time.Time `json:"ended_at,omitempty"`
	Score         *int       `json:"score,omitempty"`
	MaxScore      int        `json:"max_score"`
	QuestionCount int        `json:"question_count"`
	AttemptType   string     `json:"attempt_type"`
}

type QuestionResult struct {
	Question         QuestionWithChoices `json:"question"`
	UserChoiceIDs    []int               `json:"user_choice_ids"`
	CorrectChoiceIDs []int               `json:"correct_choice_ids"`
	IsCorrect        bool                `json:"is_correct"`
}
