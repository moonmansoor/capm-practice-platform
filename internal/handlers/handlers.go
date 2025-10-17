package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"capm-exam-system/internal/models"
	"capm-exam-system/internal/pdf"
	"capm-exam-system/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handlers struct {
	service    *service.Service
	pdfService *pdf.PDFService
}

func New(service *service.Service, pdfService *pdf.PDFService) *Handlers {
	return &Handlers{
		service:    service,
		pdfService: pdfService,
	}
}

func (h *Handlers) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/exams/start", h.StartExam).Methods("POST")
	api.HandleFunc("/quiz/start", h.StartShortQuiz).Methods("POST")
	api.HandleFunc("/hard/start", h.StartHardDrill).Methods("POST")
	api.HandleFunc("/pmp/start", h.StartPmpExam).Methods("POST")
	api.HandleFunc("/exams/{attemptId}/questions", h.GetExamQuestions).Methods("GET")
	api.HandleFunc("/exams/{attemptId}/submit", h.SubmitExam).Methods("POST")
	api.HandleFunc("/exams/{attemptId}/results", h.GetExamResults).Methods("GET")
	api.HandleFunc("/exams/{attemptId}/report.pdf", h.DownloadReport).Methods("GET")
	api.HandleFunc("/users/{userId}/attempts", h.GetUserAttempts).Methods("GET")
	api.HandleFunc("/users/login", h.LoginUser).Methods("POST")
	api.HandleFunc("/attempts/{attemptId}", h.DeleteAttempt).Methods("DELETE")
	api.HandleFunc("/earned-value/questions", h.GetEarnedValueQuestions).Methods("GET")
	api.HandleFunc("/pert/questions", h.GetPertQuestions).Methods("GET")
	api.HandleFunc("/stakeholder-salience/questions", h.GetStakeholderSalienceQuestions).Methods("GET")
	api.HandleFunc("/project-operations/questions", h.GetProjectOperationsQuestions).Methods("GET")
	api.HandleFunc("/team-motivation/questions", h.GetTeamMotivationQuestions).Methods("GET")

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	// Frontend routes
	r.HandleFunc("/", h.HomePage).Methods("GET")
	r.HandleFunc("/exam/{attemptId}", h.ExamPage).Methods("GET")
	r.HandleFunc("/quiz/{attemptId}", h.QuizPage).Methods("GET")
	r.HandleFunc("/results/{attemptId}", h.ResultsPage).Methods("GET")
	r.HandleFunc("/formula", h.FormulaPage).Methods("GET")
	r.HandleFunc("/earned-value-drill", h.EarnedValueDrillPage).Methods("GET")
	r.HandleFunc("/pert-drill", h.PertDrillPage).Methods("GET")
	r.HandleFunc("/practice", h.PracticePage).Methods("GET")
	r.HandleFunc("/stakeholder-salience", h.StakeholderSaliencePage).Methods("GET")
	r.HandleFunc("/project-operations", h.ProjectOperationsPage).Methods("GET")
	r.HandleFunc("/team-motivation", h.TeamMotivationPage).Methods("GET")

	return r
}

func (h *Handlers) StartExam(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Name == "" {
		http.Error(w, "Email and name are required", http.StatusBadRequest)
		return
	}

	// Get or create user
	user, err := h.service.GetOrCreateUser(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Start exam
	attempt, err := h.service.StartExam(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Failed to start exam", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"attempt_id": attempt.ID,
		"user_id":    user.ID,
		"started_at": attempt.StartedAt,
	})
}

func (h *Handlers) StartShortQuiz(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Name == "" {
		http.Error(w, "Email and name are required", http.StatusBadRequest)
		return
	}

	// Get or create user
	user, err := h.service.GetOrCreateUser(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Start short quiz
	attempt, err := h.service.StartShortQuiz(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Failed to start short quiz", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"attempt_id": attempt.ID,
		"user_id":    user.ID,
		"started_at": attempt.StartedAt,
	})
}

func (h *Handlers) StartPmpExam(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Name == "" {
		http.Error(w, "Email and name are required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetOrCreateUser(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	attempt, err := h.service.StartPMPExam(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Failed to start PMP exam", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"attempt_id": attempt.ID,
		"user_id":    user.ID,
		"started_at": attempt.StartedAt,
	})
}

func (h *Handlers) StartHardDrill(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Name == "" {
		http.Error(w, "Email and name are required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetOrCreateUser(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	attempt, err := h.service.StartHardDrill(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Failed to start hard drill", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"attempt_id": attempt.ID,
		"user_id":    user.ID,
		"started_at": attempt.StartedAt,
	})
}

func (h *Handlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Name == "" {
		http.Error(w, "Email and name are required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetOrCreateUser(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	attempts, err := h.service.GetUserAttemptHistory(r.Context(), user.ID)
	if err != nil {
		http.Error(w, "Failed to load attempt history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"attempts": attempts,
	})
}

func (h *Handlers) GetExamQuestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptIDStr := vars["attemptId"]

	attemptID, err := uuid.Parse(attemptIDStr)
	if err != nil {
		http.Error(w, "Invalid attempt ID", http.StatusBadRequest)
		return
	}

	questions, err := h.service.GetExamQuestions(r.Context(), attemptID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *Handlers) SubmitExam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptIDStr := vars["attemptId"]

	attemptID, err := uuid.Parse(attemptIDStr)
	if err != nil {
		http.Error(w, "Invalid attempt ID", http.StatusBadRequest)
		return
	}

	var submission models.ExamSubmission
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.SubmitExam(r.Context(), attemptID, submission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) DownloadReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptIDStr := vars["attemptId"]

	attemptID, err := uuid.Parse(attemptIDStr)
	if err != nil {
		http.Error(w, "Invalid attempt ID", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetExamResult(r.Context(), attemptID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pdfBuffer, err := h.pdfService.GenerateExamReport(result)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=capm_exam_report.pdf")
	w.Header().Set("Content-Length", strconv.Itoa(pdfBuffer.Len()))

	w.Write(pdfBuffer.Bytes())
}

func (h *Handlers) DeleteAttempt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptIDStr := vars["attemptId"]

	attemptID, err := uuid.Parse(attemptIDStr)
	if err != nil {
		http.Error(w, "Invalid attempt ID", http.StatusBadRequest)
		return
	}

	var req struct {
		UserID string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteAttempt(r.Context(), userID, attemptID); err != nil {
		switch {
		case errors.Is(err, service.ErrAttemptNotFound):
			http.Error(w, "Attempt not found", http.StatusNotFound)
		case errors.Is(err, service.ErrAttemptForbidden):
			http.Error(w, "Attempt does not belong to user", http.StatusForbidden)
		case errors.Is(err, service.ErrAttemptAlreadyClosed):
			http.Error(w, "Attempt already completed", http.StatusBadRequest)
		default:
			http.Error(w, "Failed to delete attempt", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) GetEarnedValueQuestions(w http.ResponseWriter, r *http.Request) {
	count := 10
	if raw := r.URL.Query().Get("count"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			count = parsed
		}
	}
	if count > 50 {
		count = 50
	}

	questions, err := h.service.GetEarnedValueDrill(r.Context(), count)
	if err != nil {
		http.Error(w, "Failed to load earned value questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *Handlers) GetPertQuestions(w http.ResponseWriter, r *http.Request) {
	count := 10
	if raw := r.URL.Query().Get("count"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			count = parsed
		}
	}
	if count > 50 {
		count = 50
	}

	questions, err := h.service.GetPertDrill(r.Context(), count)
	if err != nil {
		http.Error(w, "Failed to load PERT questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *Handlers) GetStakeholderSalienceQuestions(w http.ResponseWriter, r *http.Request) {
	count := 10
	if raw := r.URL.Query().Get("count"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			count = parsed
		}
	}
	if count > 50 {
		count = 50
	}

	questions, err := h.service.GetStakeholderSalienceDrill(r.Context(), count)
	if err != nil {
		http.Error(w, "Failed to load stakeholder salience questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *Handlers) GetProjectOperationsQuestions(w http.ResponseWriter, r *http.Request) {
	count := 15
	if raw := r.URL.Query().Get("count"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			count = parsed
		}
	}
	if count > 20 {
		count = 20
	}

	questions, err := h.service.GetProjectOperationsDrill(r.Context(), count)
	if err != nil {
		http.Error(w, "Failed to load project vs operations questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *Handlers) GetTeamMotivationQuestions(w http.ResponseWriter, r *http.Request) {
	count := 20
	if raw := r.URL.Query().Get("count"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			count = parsed
		}
	}
	if count > 20 {
		count = 20
	}

	questions, err := h.service.GetTeamMotivationDrill(r.Context(), count)
	if err != nil {
		http.Error(w, "Failed to load team motivation questions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *Handlers) GetExamResults(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptIDStr := vars["attemptId"]

	attemptID, err := uuid.Parse(attemptIDStr)
	if err != nil {
		http.Error(w, "Invalid attempt ID", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetExamResult(r.Context(), attemptID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handlers) GetUserAttempts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userId"]

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	attempts, err := h.service.GetUserAttemptHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attempts)
}

// Frontend handlers
func (h *Handlers) HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/index.html")
}

func (h *Handlers) PracticePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/practice.html")
}

func (h *Handlers) StakeholderSaliencePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/stakeholder_salience.html")
}

func (h *Handlers) ProjectOperationsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/project_operations.html")
}

func (h *Handlers) TeamMotivationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/team_motivation.html")
}

func (h *Handlers) ExamPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/exam.html")
}

func (h *Handlers) QuizPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/quiz.html")
}

func (h *Handlers) ResultsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/results.html")
}

func (h *Handlers) FormulaPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/formula.html")
}

func (h *Handlers) EarnedValueDrillPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/earned_value_drill.html")
}

func (h *Handlers) PertDrillPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/pert_drill.html")
}
