package repository

import (
	"context"
	"fmt"
	"math/rand"

	"capm-exam-system/internal/database"
	"capm-exam-system/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	pmpExamName  = "PMP Mock Exam"
	hardExamName = "Hard Question Drill"
)

type Repository struct {
	db *database.DB
}

func New(db *database.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, email, name string) (*models.User, error) {
	var user models.User
	err := r.db.Pool.QueryRow(ctx,
		"INSERT INTO users (email, name) VALUES ($1, $2) RETURNING id, email, name",
		email, name).Scan(&user.ID, &user.Email, &user.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return &user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.Pool.QueryRow(ctx,
		"SELECT id, email, name FROM users WHERE email = $1",
		email).Scan(&user.ID, &user.Email, &user.Name)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &user, nil
}

func (r *Repository) CreateQuestion(ctx context.Context, prompt, domain, explanation string, popularityScore float64, isMultiSelect bool) (*models.Question, error) {
	var question models.Question
	err := r.db.Pool.QueryRow(ctx,
		"INSERT INTO questions (prompt, domain, explanation, popularity_score, is_multi_select) VALUES ($1, $2, $3, $4, $5) RETURNING id, prompt, domain, explanation, popularity_score, is_multi_select",
		prompt, domain, explanation, popularityScore, isMultiSelect).Scan(
		&question.ID, &question.Prompt, &question.Domain, &question.Explanation, &question.PopularityScore, &question.IsMultiSelect)

	if err != nil {
		return nil, fmt.Errorf("failed to create question: %v", err)
	}
	return &question, nil
}

func (r *Repository) CreateChoice(ctx context.Context, questionID int, text, label string, isCorrect bool) (*models.Choice, error) {
	var choice models.Choice
	err := r.db.Pool.QueryRow(ctx,
		"INSERT INTO choices (question_id, text, label, is_correct) VALUES ($1, $2, $3, $4) RETURNING id, question_id, text, label, is_correct",
		questionID, text, label, isCorrect).Scan(
		&choice.ID, &choice.QuestionID, &choice.Text, &choice.Label, &choice.IsCorrect)

	if err != nil {
		return nil, fmt.Errorf("failed to create choice: %v", err)
	}
	return &choice, nil
}

func (r *Repository) GetQuestionsWithChoices(ctx context.Context, questionIDs []int) ([]models.QuestionWithChoices, error) {
	if len(questionIDs) == 0 {
		return []models.QuestionWithChoices{}, nil
	}

	// Create placeholder string for IN clause
	placeholders := make([]interface{}, len(questionIDs))
	for i, id := range questionIDs {
		placeholders[i] = id
	}

	query := `
		SELECT q.id, q.prompt, q.domain, q.popularity_score, q.explanation, q.is_multi_select,
		       c.id, c.text, c.label, c.is_correct
		FROM questions q
		JOIN choices c ON q.id = c.question_id
		WHERE q.id = ANY($1)
		ORDER BY q.id, c.label`

	rows, err := r.db.Pool.Query(ctx, query, questionIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions with choices: %v", err)
	}
	defer rows.Close()

	questionMap := make(map[int]*models.QuestionWithChoices)

	for rows.Next() {
		var qID int
		var q models.Question
		var c models.Choice

		err := rows.Scan(&qID, &q.Prompt, &q.Domain, &q.PopularityScore, &q.Explanation, &q.IsMultiSelect,
			&c.ID, &c.Text, &c.Label, &c.IsCorrect)
		if err != nil {
			return nil, fmt.Errorf("failed to scan question row: %v", err)
		}

		q.ID = qID
		c.QuestionID = qID

		if questionMap[qID] == nil {
			questionMap[qID] = &models.QuestionWithChoices{
				Question: q,
				Choices:  []models.Choice{},
			}
		}

		questionMap[qID].Choices = append(questionMap[qID].Choices, c)
	}

	result := make([]models.QuestionWithChoices, 0, len(questionMap))
	for _, qID := range questionIDs {
		if q, exists := questionMap[qID]; exists {
			result = append(result, *q)
		}
	}

	return result, nil
}

func (r *Repository) GetEarnedValueDrillQuestions(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 10
	}

	query := `
		SELECT id
		FROM questions
		WHERE domain = $1
		ORDER BY random()
		LIMIT $2`

	rows, err := r.db.Pool.Query(ctx, query, "Earned Value Drill", count)
	if err != nil {
		return nil, fmt.Errorf("failed to select earned value drill questions: %v", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan earned value drill question id: %v", err)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return []models.QuestionWithChoices{}, nil
	}

	return r.GetQuestionsWithChoices(ctx, ids)
}

func (r *Repository) GetPertDrillQuestions(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 10
	}

	query := `
		SELECT id
		FROM questions
		WHERE domain = $1
		ORDER BY random()
		LIMIT $2`

	rows, err := r.db.Pool.Query(ctx, query, "PERT Drill", count)
	if err != nil {
		return nil, fmt.Errorf("failed to select PERT drill questions: %v", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan PERT drill question id: %v", err)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return []models.QuestionWithChoices{}, nil
	}

	return r.GetQuestionsWithChoices(ctx, ids)
}

func (r *Repository) GetStakeholderSalienceDrillQuestions(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 10
	}

	query := `
		SELECT id
		FROM questions
		WHERE domain = $1
		ORDER BY random()
		LIMIT $2`

	rows, err := r.db.Pool.Query(ctx, query, "Stakeholder Salience Drill", count)
	if err != nil {
		return nil, fmt.Errorf("failed to select stakeholder salience questions: %v", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan stakeholder salience question id: %v", err)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return []models.QuestionWithChoices{}, nil
	}

	return r.GetQuestionsWithChoices(ctx, ids)
}

func (r *Repository) GetProjectOperationsDrillQuestions(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 15
	}
	if count > 20 {
		count = 20
	}

	query := `
		SELECT id
		FROM questions
		WHERE domain = $1
		ORDER BY random()
		LIMIT $2`

	rows, err := r.db.Pool.Query(ctx, query, "Project Operations Classification Drill", count)
	if err != nil {
		return nil, fmt.Errorf("failed to select project vs operations questions: %v", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan project vs operations question id: %v", err)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return []models.QuestionWithChoices{}, nil
	}

	return r.GetQuestionsWithChoices(ctx, ids)
}

func (r *Repository) GetTeamMotivationDrillQuestions(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 20
	}
	if count > 20 {
		count = 20
	}

	query := `
		SELECT id
		FROM questions
		WHERE domain = $1
		ORDER BY random()
		LIMIT $2`

	rows, err := r.db.Pool.Query(ctx, query, "Team Motivation Drill", count)
	if err != nil {
		return nil, fmt.Errorf("failed to select team motivation drill questions: %v", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan team motivation drill question id: %v", err)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return []models.QuestionWithChoices{}, nil
	}

	return r.GetQuestionsWithChoices(ctx, ids)
}

func (r *Repository) GetRandomQuestions(ctx context.Context, count int, seed int64) ([]int, error) {
	query := `
		SELECT id, popularity_score
		FROM questions
		ORDER BY popularity_score DESC`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions: %v", err)
	}
	defer rows.Close()

	type weightedQuestion struct {
		ID     int
		Weight float64
	}

	var questions []weightedQuestion
	for rows.Next() {
		var wq weightedQuestion
		if err := rows.Scan(&wq.ID, &wq.Weight); err != nil {
			return nil, fmt.Errorf("failed to scan question: %v", err)
		}
		questions = append(questions, wq)
	}

	if len(questions) < count {
		return nil, fmt.Errorf("not enough questions available: have %d, need %d", len(questions), count)
	}

	// Use seed for reproducible randomness
	rng := rand.New(rand.NewSource(seed))

	// Weighted random selection
	selected := make([]int, 0, count)
	used := make(map[int]bool)

	for len(selected) < count {
		totalWeight := 0.0
		for _, q := range questions {
			if !used[q.ID] {
				totalWeight += q.Weight
			}
		}

		if totalWeight == 0 {
			break
		}

		target := rng.Float64() * totalWeight
		current := 0.0

		for _, q := range questions {
			if used[q.ID] {
				continue
			}
			current += q.Weight
			if current >= target {
				selected = append(selected, q.ID)
				used[q.ID] = true
				break
			}
		}
	}

	// Fill remaining with random unused questions if needed
	for len(selected) < count {
		for _, q := range questions {
			if !used[q.ID] {
				selected = append(selected, q.ID)
				used[q.ID] = true
				break
			}
		}
	}

	return selected, nil
}

func (r *Repository) GetRandomQuestionsByDomain(ctx context.Context, count int, seed int64, domain string, excludeIDs []int) ([]int, error) {
	query := `
		SELECT id, popularity_score
		FROM questions
		WHERE domain = $1`

	args := []interface{}{domain}

	if len(excludeIDs) > 0 {
		query += ` AND NOT (id = ANY($2))`
		exclusion := make([]int32, len(excludeIDs))
		for i, id := range excludeIDs {
			exclusion[i] = int32(id)
		}
		args = append(args, exclusion)
	}

	query += `
		ORDER BY popularity_score DESC`

	rows, err := r.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions for domain %s: %v", domain, err)
	}
	defer rows.Close()

	type weightedQuestion struct {
		ID     int
		Weight float64
	}

	var questions []weightedQuestion
	for rows.Next() {
		var wq weightedQuestion
		if err := rows.Scan(&wq.ID, &wq.Weight); err != nil {
			return nil, fmt.Errorf("failed to scan question: %v", err)
		}
		questions = append(questions, wq)
	}

	if len(questions) < count {
		return nil, fmt.Errorf("not enough questions available in domain %s: have %d, need %d", domain, len(questions), count)
	}

	rng := rand.New(rand.NewSource(seed))
	selected := make([]int, 0, count)
	used := make(map[int]bool)

	for len(selected) < count {
		totalWeight := 0.0
		for _, q := range questions {
			if !used[q.ID] {
				totalWeight += q.Weight
			}
		}

		if totalWeight == 0 {
			break
		}

		target := rng.Float64() * totalWeight
		current := 0.0

		for _, q := range questions {
			if used[q.ID] {
				continue
			}
			current += q.Weight
			if current >= target {
				selected = append(selected, q.ID)
				used[q.ID] = true
				break
			}
		}
	}

	for len(selected) < count {
		for _, q := range questions {
			if !used[q.ID] {
				selected = append(selected, q.ID)
				used[q.ID] = true
				break
			}
		}
	}

	return selected, nil
}

func (r *Repository) GetAttemptsByUser(ctx context.Context, userID uuid.UUID) ([]models.AttemptHistory, error) {
	query := `
		SELECT a.id, a.exam_id, e.name, a.max_score, a.score, a.started_at, a.ended_at
		FROM attempts a
		JOIN exams e ON e.id = a.exam_id
		WHERE a.user_id = $1
		ORDER BY a.started_at DESC`

	rows, err := r.db.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempts: %v", err)
	}
	defer rows.Close()

	var history []models.AttemptHistory
	for rows.Next() {
		var record models.AttemptHistory
		var score pgtype.Int4

		if err := rows.Scan(&record.AttemptID, &record.ExamID, &record.ExamName, &record.MaxScore, &score, &record.StartedAt, &record.EndedAt); err != nil {
			return nil, fmt.Errorf("failed to scan attempt history: %v", err)
		}

		if score.Valid {
			val := int(score.Int32)
			record.Score = &val
		}

		record.QuestionCount = record.MaxScore

		switch {
		case record.ExamName == pmpExamName:
			record.AttemptType = "PMP Mock Exam"
		case record.ExamName == hardExamName:
			record.AttemptType = "Hard Drill"
		case record.MaxScore <= 20:
			record.AttemptType = "Short Quiz"
		default:
			record.AttemptType = "Mock Exam"
		}

		history = append(history, record)
	}

	return history, nil
}

func (r *Repository) GetRandomQuestionsExcludingDomain(ctx context.Context, count int, seed int64, domain string) ([]int, error) {
	query := `
		SELECT id, popularity_score
		FROM questions
		WHERE domain <> $1
		ORDER BY popularity_score DESC`

	rows, err := r.db.Pool.Query(ctx, query, domain)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions excluding %s: %v", domain, err)
	}
	defer rows.Close()

	type weightedQuestion struct {
		ID     int
		Weight float64
	}

	var questions []weightedQuestion
	for rows.Next() {
		var wq weightedQuestion
		if err := rows.Scan(&wq.ID, &wq.Weight); err != nil {
			return nil, fmt.Errorf("failed to scan question: %v", err)
		}
		questions = append(questions, wq)
	}

	if len(questions) < count {
		return nil, fmt.Errorf("not enough questions outside domain %s: have %d, need %d", domain, len(questions), count)
	}

	rng := rand.New(rand.NewSource(seed))
	selected := make([]int, 0, count)
	used := make(map[int]bool)

	for len(selected) < count {
		totalWeight := 0.0
		for _, q := range questions {
			if !used[q.ID] {
				totalWeight += q.Weight
			}
		}

		if totalWeight == 0 {
			break
		}

		target := rng.Float64() * totalWeight
		current := 0.0

		for _, q := range questions {
			if used[q.ID] {
				continue
			}
			current += q.Weight
			if current >= target {
				selected = append(selected, q.ID)
				used[q.ID] = true
				break
			}
		}
	}

	for len(selected) < count {
		for _, q := range questions {
			if !used[q.ID] {
				selected = append(selected, q.ID)
				used[q.ID] = true
				break
			}
		}
	}

	return selected, nil
}

func (r *Repository) CreateAttempt(ctx context.Context, userID uuid.UUID, examID uuid.UUID, seed int64, maxScore int) (*models.Attempt, error) {
	var attempt models.Attempt
	err := r.db.Pool.QueryRow(ctx,
		"INSERT INTO attempts (user_id, exam_id, seed, max_score) VALUES ($1, $2, $3, $4) RETURNING id, exam_id, user_id, seed, score, max_score, started_at, ended_at",
		userID, examID, seed, maxScore).Scan(
		&attempt.ID, &attempt.ExamID, &attempt.UserID, &attempt.Seed, &attempt.Score, &attempt.MaxScore, &attempt.StartedAt, &attempt.EndedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create attempt: %v", err)
	}
	return &attempt, nil
}

func (r *Repository) GetAttempt(ctx context.Context, attemptID uuid.UUID) (*models.Attempt, error) {
	var attempt models.Attempt
	err := r.db.Pool.QueryRow(ctx,
		"SELECT id, exam_id, user_id, seed, score, max_score, started_at, ended_at FROM attempts WHERE id = $1",
		attemptID).Scan(
		&attempt.ID, &attempt.ExamID, &attempt.UserID, &attempt.Seed, &attempt.Score, &attempt.MaxScore, &attempt.StartedAt, &attempt.EndedAt)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt: %v", err)
	}
	return &attempt, nil
}

func (r *Repository) CreateAttemptAnswer(ctx context.Context, attemptID uuid.UUID, questionID, choiceID int, isCorrect bool) error {
	_, err := r.db.Pool.Exec(ctx,
		"INSERT INTO attempt_answers (attempt_id, question_id, choice_id, is_correct) VALUES ($1, $2, $3, $4)",
		attemptID, questionID, choiceID, isCorrect)

	if err != nil {
		return fmt.Errorf("failed to create attempt answer: %v", err)
	}
	return nil
}

func (r *Repository) UpdateAttemptScore(ctx context.Context, attemptID uuid.UUID, score int) error {
	_, err := r.db.Pool.Exec(ctx,
		"UPDATE attempts SET score = $1, ended_at = NOW() WHERE id = $2",
		score, attemptID)

	if err != nil {
		return fmt.Errorf("failed to update attempt score: %v", err)
	}
	return nil
}

func (r *Repository) DeleteAttempt(ctx context.Context, attemptID uuid.UUID) error {
	if _, err := r.db.Pool.Exec(ctx, "DELETE FROM attempt_answers WHERE attempt_id = $1", attemptID); err != nil {
		return fmt.Errorf("failed to delete attempt answers: %v", err)
	}

	commandTag, err := r.db.Pool.Exec(ctx, "DELETE FROM attempts WHERE id = $1", attemptID)
	if err != nil {
		return fmt.Errorf("failed to delete attempt: %v", err)
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("attempt not found")
	}
	return nil
}

func (r *Repository) GetAttemptAnswers(ctx context.Context, attemptID uuid.UUID) ([]models.AttemptAnswer, error) {
	query := `
		SELECT id, attempt_id, question_id, choice_id, is_correct
		FROM attempt_answers
		WHERE attempt_id = $1
		ORDER BY question_id`

	rows, err := r.db.Pool.Query(ctx, query, attemptID)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt answers: %v", err)
	}
	defer rows.Close()

	var answers []models.AttemptAnswer
	for rows.Next() {
		var answer models.AttemptAnswer
		err := rows.Scan(&answer.ID, &answer.AttemptID, &answer.QuestionID, &answer.ChoiceID, &answer.IsCorrect)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attempt answer: %v", err)
		}
		answers = append(answers, answer)
	}

	return answers, nil
}

func (r *Repository) UpdateAttemptAnswerCorrectness(ctx context.Context, answerID uuid.UUID, isCorrect bool) error {
	_, err := r.db.Pool.Exec(ctx,
		"UPDATE attempt_answers SET is_correct = $1 WHERE id = $2",
		isCorrect, answerID)

	if err != nil {
		return fmt.Errorf("failed to update attempt answer correctness: %v", err)
	}
	return nil
}

func (r *Repository) CreateExam(ctx context.Context, name, description string) (*models.Exam, error) {
	var exam models.Exam
	err := r.db.Pool.QueryRow(ctx,
		"INSERT INTO exams (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at",
		name, description).Scan(
		&exam.ID, &exam.Name, &exam.Description, &exam.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create exam: %v", err)
	}
	return &exam, nil
}

func (r *Repository) GetExamByName(ctx context.Context, name string) (*models.Exam, error) {
	var exam models.Exam
	err := r.db.Pool.QueryRow(ctx,
		"SELECT id, name, description, created_at FROM exams WHERE name = $1 LIMIT 1",
		name).Scan(
		&exam.ID, &exam.Name, &exam.Description, &exam.CreatedAt)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get exam: %v", err)
	}
	return &exam, nil
}

func (r *Repository) GetExamByID(ctx context.Context, examID uuid.UUID) (*models.Exam, error) {
	var exam models.Exam
	err := r.db.Pool.QueryRow(ctx,
		"SELECT id, name, description, created_at FROM exams WHERE id = $1",
		examID).Scan(&exam.ID, &exam.Name, &exam.Description, &exam.CreatedAt)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get exam by id: %v", err)
	}
	return &exam, nil
}

func (r *Repository) CreateDefaultExam(ctx context.Context) (*models.Exam, error) {
	return r.CreateExam(ctx, "CAPM Mock Exam", "150-question CAPM certification practice exam")
}

func (r *Repository) GetDefaultExam(ctx context.Context) (*models.Exam, error) {
	return r.GetExamByName(ctx, "CAPM Mock Exam")
}
