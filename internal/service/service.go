package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"capm-exam-system/internal/models"
	"capm-exam-system/internal/repository"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

var (
	ErrAttemptNotFound      = errors.New("attempt not found")
	ErrAttemptAlreadyClosed = errors.New("attempt already completed")
	ErrAttemptForbidden     = errors.New("user cannot modify this attempt")
)

const (
	defaultExamName     = "CAPM Mock Exam"
	pmpExamName         = "PMP Mock Exam"
	hardExamName        = "Hard Question Drill"
	hardDomainName      = "Hard Question"
	shortQuizHardTarget = 3
)

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetOrCreateUser(ctx context.Context, email, name string) (*models.User, error) {
	// Try to get existing user
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Create new user if doesn't exist
	if user == nil {
		user, err = s.repo.CreateUser(ctx, email, name)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *Service) StartExam(ctx context.Context, userID uuid.UUID) (*models.Attempt, error) {
	return s.StartExamWithQuestionCount(ctx, userID, 150)
}

func (s *Service) StartShortQuiz(ctx context.Context, userID uuid.UUID) (*models.Attempt, error) {
	return s.StartExamWithQuestionCount(ctx, userID, 15)
}

func (s *Service) StartPMPExam(ctx context.Context, userID uuid.UUID) (*models.Attempt, error) {
	exam, err := s.repo.GetExamByName(ctx, pmpExamName)
	if err != nil {
		return nil, err
	}

	if exam == nil {
		exam, err = s.repo.CreateExam(ctx, pmpExamName, "150-question PMP scenario exam")
		if err != nil {
			return nil, err
		}
	}

	return s.createAttemptForExam(ctx, userID, exam, 150)
}

func (s *Service) StartExamWithQuestionCount(ctx context.Context, userID uuid.UUID, questionCount int) (*models.Attempt, error) {
	// Get or create default exam
	exam, err := s.repo.GetExamByName(ctx, defaultExamName)
	if err != nil {
		return nil, err
	}

	if exam == nil {
		exam, err = s.repo.CreateExam(ctx, defaultExamName, "150-question CAPM certification practice exam")
		if err != nil {
			return nil, err
		}
	}

	return s.createAttemptForExam(ctx, userID, exam, questionCount)
}

func (s *Service) StartHardDrill(ctx context.Context, userID uuid.UUID) (*models.Attempt, error) {
	exam, err := s.repo.GetExamByName(ctx, hardExamName)
	if err != nil {
		return nil, err
	}

	if exam == nil {
		exam, err = s.repo.CreateExam(ctx, hardExamName, "20-question advanced CAPM scenario drill")
		if err != nil {
			return nil, err
		}
	}

	return s.createAttemptForExam(ctx, userID, exam, 20)
}

func (s *Service) GetExamQuestions(ctx context.Context, attemptID uuid.UUID) ([]models.QuestionWithChoices, error) {
	// Get attempt to verify it exists and get seed
	attempt, err := s.repo.GetAttempt(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	if attempt == nil {
		return nil, fmt.Errorf("attempt not found")
	}

	exam, err := s.repo.GetExamByID(ctx, attempt.ExamID)
	if err != nil {
		return nil, err
	}
	if exam == nil {
		return nil, fmt.Errorf("exam not found")
	}

	questionIDs, err := s.pickQuestionIDs(ctx, attempt, exam)
	if err != nil {
		return nil, err
	}

	// Get questions with choices
	questions, err := s.repo.GetQuestionsWithChoices(ctx, questionIDs)
	if err != nil {
		return nil, err
	}

	// Remove explanation and correct answers from the response
	for i := range questions {
		questions[i].Explanation = "" // Hide explanation until submission
		for j := range questions[i].Choices {
			questions[i].Choices[j].IsCorrect = false // Hide correct answers
		}
	}

	return questions, nil
}

func (s *Service) SubmitExam(ctx context.Context, attemptID uuid.UUID, submission models.ExamSubmission) (*models.ExamResult, error) {
	// Get attempt
	attempt, err := s.repo.GetAttempt(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	if attempt == nil {
		return nil, fmt.Errorf("attempt not found")
	}

	// Check if already submitted
	if attempt.EndedAt != nil {
		return nil, fmt.Errorf("exam already submitted")
	}

	exam, err := s.repo.GetExamByID(ctx, attempt.ExamID)
	if err != nil {
		return nil, err
	}
	if exam == nil {
		return nil, fmt.Errorf("exam not found")
	}

	questionIDs, err := s.pickQuestionIDs(ctx, attempt, exam)
	if err != nil {
		return nil, err
	}

	questions, err := s.repo.GetQuestionsWithChoices(ctx, questionIDs)
	if err != nil {
		return nil, err
	}

	// Create a map for quick lookup
	questionMap := make(map[int]models.QuestionWithChoices)
	for _, q := range questions {
		questionMap[q.ID] = q
	}

	// Normalize submitted answers by question
	answersByQuestion := make(map[int][]int)
	for _, answer := range submission.Answers {
		question, exists := questionMap[answer.QuestionID]
		if !exists {
			continue
		}

		if len(answer.ChoiceIDs) == 0 {
			continue
		}

		selectedSet := make(map[int]struct{})
		for _, choiceID := range answer.ChoiceIDs {
			if choiceID <= 0 {
				continue
			}
			selectedSet[choiceID] = struct{}{}
		}

		if len(selectedSet) == 0 {
			continue
		}

		orderedSelection := make([]int, 0, len(selectedSet))
		for _, choice := range question.Choices {
			if _, ok := selectedSet[choice.ID]; ok {
				orderedSelection = append(orderedSelection, choice.ID)
			}
		}

		if len(orderedSelection) == 0 {
			continue
		}

		answersByQuestion[answer.QuestionID] = orderedSelection
	}

	score := 0
	results := make([]models.QuestionResult, 0, len(questionIDs))

	for _, qID := range questionIDs {
		question := questionMap[qID]

		correctIDs := make([]int, 0)
		correctSet := make(map[int]struct{})
		for _, choice := range question.Choices {
			if choice.IsCorrect {
				correctIDs = append(correctIDs, choice.ID)
				correctSet[choice.ID] = struct{}{}
			}
		}

		selectedIDs, answered := answersByQuestion[qID]
		if !answered {
			selectedIDs = nil
		}

		isCorrect := false
		if len(selectedIDs) > 0 && len(selectedIDs) == len(correctIDs) {
			isCorrect = true
			for _, sel := range selectedIDs {
				if _, ok := correctSet[sel]; !ok {
					isCorrect = false
					break
				}
			}
		}

		if isCorrect {
			score++
		}

		if len(selectedIDs) > 0 {
			for _, choiceID := range selectedIDs {
				if err := s.repo.CreateAttemptAnswer(ctx, attemptID, question.ID, choiceID, isCorrect); err != nil {
					return nil, err
				}
			}
		}

		result := models.QuestionResult{
			Question:         question,
			UserChoiceIDs:    selectedIDs,
			CorrectChoiceIDs: correctIDs,
			IsCorrect:        isCorrect,
		}
		results = append(results, result)
	}

	// Update attempt with score
	err = s.repo.UpdateAttemptScore(ctx, attemptID, score)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	attempt.Score = &score
	attempt.EndedAt = &now

	examResult := &models.ExamResult{
		AttemptID: attemptID,
		UserID:    attempt.UserID,
		ExamID:    attempt.ExamID,
		Score:     score,
		MaxScore:  attempt.MaxScore,
		StartedAt: attempt.StartedAt,
		EndedAt:   attempt.EndedAt,
		Results:   results,
	}

	return examResult, nil
}

func (s *Service) GetExamResult(ctx context.Context, attemptID uuid.UUID) (*models.ExamResult, error) {
	// Get attempt
	attempt, err := s.repo.GetAttempt(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	if attempt == nil {
		return nil, fmt.Errorf("attempt not found")
	}

	// Check if exam was submitted
	if attempt.EndedAt == nil || attempt.Score == nil {
		return nil, fmt.Errorf("exam not yet submitted")
	}

	exam, err := s.repo.GetExamByID(ctx, attempt.ExamID)
	if err != nil {
		return nil, err
	}
	if exam == nil {
		return nil, fmt.Errorf("exam not found")
	}

	questionIDs, err := s.pickQuestionIDs(ctx, attempt, exam)
	if err != nil {
		return nil, err
	}

	questions, err := s.repo.GetQuestionsWithChoices(ctx, questionIDs)
	if err != nil {
		return nil, err
	}

	// Get user answers
	answers, err := s.repo.GetAttemptAnswers(ctx, attemptID)
	if err != nil {
		return nil, err
	}

	// Create maps for quick lookup
	questionMap := make(map[int]models.QuestionWithChoices)
	for _, q := range questions {
		questionMap[q.ID] = q
	}

	answersByQuestion := make(map[int][]models.AttemptAnswer)
	for _, a := range answers {
		answersByQuestion[a.QuestionID] = append(answersByQuestion[a.QuestionID], a)
	}

	// Build results
	results := make([]models.QuestionResult, 0, len(questions))

	for _, question := range questions {
		correctIDs := make([]int, 0)
		correctSet := make(map[int]struct{})
		for _, choice := range question.Choices {
			if choice.IsCorrect {
				correctIDs = append(correctIDs, choice.ID)
				correctSet[choice.ID] = struct{}{}
			}
		}

		var userChoiceIDs []int
		var recordedCorrect *bool
		if submitted := answersByQuestion[question.ID]; len(submitted) > 0 {
			selectedSet := make(map[int]struct{})
			for _, ans := range submitted {
				if ans.ChoiceID != nil {
					selectedSet[*ans.ChoiceID] = struct{}{}
				}
				if ans.IsCorrect != nil {
					recordedCorrect = ans.IsCorrect
				}
			}

			userChoiceIDs = make([]int, 0, len(selectedSet))
			for _, choice := range question.Choices {
				if _, ok := selectedSet[choice.ID]; ok {
					userChoiceIDs = append(userChoiceIDs, choice.ID)
				}
			}
		}

		isCorrect := false
		if recordedCorrect != nil {
			isCorrect = *recordedCorrect
		} else if len(userChoiceIDs) > 0 && len(userChoiceIDs) == len(correctIDs) {
			isCorrect = true
			for _, sel := range userChoiceIDs {
				if _, ok := correctSet[sel]; !ok {
					isCorrect = false
					break
				}
			}
		}

		result := models.QuestionResult{
			Question:         question,
			UserChoiceIDs:    userChoiceIDs,
			CorrectChoiceIDs: correctIDs,
			IsCorrect:        isCorrect,
		}

		results = append(results, result)
	}

	examResult := &models.ExamResult{
		AttemptID: attemptID,
		UserID:    attempt.UserID,
		ExamID:    attempt.ExamID,
		Score:     *attempt.Score,
		MaxScore:  attempt.MaxScore,
		StartedAt: attempt.StartedAt,
		EndedAt:   attempt.EndedAt,
		Results:   results,
	}

	return examResult, nil
}

func (s *Service) GetUserAttemptHistory(ctx context.Context, userID uuid.UUID) ([]models.AttemptHistory, error) {
	return s.repo.GetAttemptsByUser(ctx, userID)
}

func (s *Service) GetEarnedValueDrill(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 10
	}
	if count > 50 {
		count = 50
	}

	return s.repo.GetEarnedValueDrillQuestions(ctx, count)
}

func (s *Service) GetPertDrill(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 10
	}
	if count > 50 {
		count = 50
	}

	return s.repo.GetPertDrillQuestions(ctx, count)
}

func (s *Service) GetStakeholderSalienceDrill(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 10
	}
	if count > 50 {
		count = 50
	}

	return s.repo.GetStakeholderSalienceDrillQuestions(ctx, count)
}

func (s *Service) GetProjectOperationsDrill(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 15
	}
	if count > 20 {
		count = 20
	}

	return s.repo.GetProjectOperationsDrillQuestions(ctx, count)
}

func (s *Service) GetTeamMotivationDrill(ctx context.Context, count int) ([]models.QuestionWithChoices, error) {
	if count <= 0 {
		count = 20
	}
	if count > 20 {
		count = 20
	}

	return s.repo.GetTeamMotivationDrillQuestions(ctx, count)
}

func (s *Service) DeleteAttempt(ctx context.Context, userID, attemptID uuid.UUID) error {
	attempt, err := s.repo.GetAttempt(ctx, attemptID)
	if err != nil {
		return err
	}
	if attempt == nil {
		return ErrAttemptNotFound
	}

	if attempt.UserID != userID {
		return ErrAttemptForbidden
	}

	if attempt.Score != nil || attempt.EndedAt != nil {
		return ErrAttemptAlreadyClosed
	}

	return s.repo.DeleteAttempt(ctx, attemptID)
}

func (s *Service) createAttemptForExam(ctx context.Context, userID uuid.UUID, exam *models.Exam, questionCount int) (*models.Attempt, error) {
	if exam == nil {
		return nil, fmt.Errorf("exam reference is nil")
	}

	seed := time.Now().UnixNano()
	attempt, err := s.repo.CreateAttempt(ctx, userID, exam.ID, seed, questionCount)
	if err != nil {
		return nil, err
	}

	return attempt, nil
}

func (s *Service) pickQuestionIDs(ctx context.Context, attempt *models.Attempt, exam *models.Exam) ([]int, error) {
	blueprint := blueprintForExam(exam.Name, attempt.MaxScore)

	if exam.Name == hardExamName {
		return s.repo.GetRandomQuestionsByDomain(ctx, attempt.MaxScore, attempt.Seed, hardDomainName, nil)
	}

	if sumQuota(blueprint) != attempt.MaxScore {
		return nil, fmt.Errorf("allocator mismatch: expected %d, got quota %d", attempt.MaxScore, sumQuota(blueprint))
	}

	selected := make([]int, 0, attempt.MaxScore)
	selectedSet := make(map[int]struct{}, attempt.MaxScore)
	seed := attempt.Seed

	// Hard question allocation first
	if blueprint.hardCount > 0 {
		hardIDs, err := s.repo.GetRandomQuestionsByDomain(ctx, blueprint.hardCount, seed, hardDomainName, nil)
		if err != nil {
			return nil, err
		}
		for _, id := range hardIDs {
			if _, exists := selectedSet[id]; exists {
				continue
			}
			selected = append(selected, id)
			selectedSet[id] = struct{}{}
		}
		seed++
	}

	for _, dq := range blueprint.domainCounts {
		if dq.count == 0 {
			continue
		}

		// If this domain is Hard Question, skip because already allocated
		if dq.domain == hardDomainName {
			continue
		}

		exclude := make([]int, 0, len(selectedSet))
		for id := range selectedSet {
			exclude = append(exclude, id)
		}

		domainIDs, err := s.repo.GetRandomQuestionsByDomain(ctx, dq.count, seed, dq.domain, exclude)
		if err != nil {
			return nil, err
		}
		for _, id := range domainIDs {
			if _, exists := selectedSet[id]; exists {
				continue
			}
			selected = append(selected, id)
			selectedSet[id] = struct{}{}
		}
		seed++
	}

	if len(selected) != attempt.MaxScore {
		return nil, fmt.Errorf("selected %d unique questions, expected %d", len(selected), attempt.MaxScore)
	}

	// Shuffle combined slice for randomness
	rng := rand.New(rand.NewSource(attempt.Seed))
	rng.Shuffle(len(selected), func(i, j int) {
		selected[i], selected[j] = selected[j], selected[i]
	})

	return selected, nil
}
