package pdf

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"capm-exam-system/internal/models"

	"github.com/jung-kurt/gofpdf/v2"
)

type PDFService struct{}

func New() *PDFService {
	return &PDFService{}
}

func (p *PDFService) GenerateExamReport(result *models.ExamResult) (*bytes.Buffer, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 15, 10)
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(190, 10, "CAPM Mock Exam Report")
	pdf.Ln(8)
	addDivider(pdf)

	// Score summary
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "Score Summary")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)
	percentage := float64(result.Score) / float64(result.MaxScore) * 100
	scoreText := fmt.Sprintf("Score: %d/%d (%.1f%%)", result.Score, result.MaxScore, percentage)
	pdf.Cell(190, 7, scoreText)
	pdf.Ln(7)

	status := "FAIL"
	if percentage >= 70 {
		status = "PASS"
	}
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 6, fmt.Sprintf("Status: %s", status))
	pdf.Ln(8)
	addDivider(pdf)

	// Domain breakdown
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 8, "Performance by Domain")
	pdf.Ln(6)

	domains := map[string]struct{ correct, total int }{}
	for _, r := range result.Results {
		stats := domains[r.Question.Domain]
		stats.total++
		if r.IsCorrect {
			stats.correct++
		}
		domains[r.Question.Domain] = stats
	}

	pdf.SetFont("Arial", "", 11)
	pdf.SetFillColor(242, 242, 242)
	pdf.CellFormat(120, 7, "Domain", "1", 0, "L", true, 0, "")
	pdf.CellFormat(70, 7, "Score", "1", 1, "R", true, 0, "")
	pdf.SetFillColor(255, 255, 255)

	for domain, stats := range domains {
		if stats.total == 0 {
			continue
		}
		pct := float64(stats.correct) / float64(stats.total) * 100
		pdf.CellFormat(120, 7, domain, "1", 0, "L", false, 0, "")
		pdf.CellFormat(70, 7, fmt.Sprintf("%d/%d (%.1f%%)", stats.correct, stats.total, pct), "1", 1, "R", false, 0, "")
	}

	pdf.Ln(8)
	addDivider(pdf)

	// Question review
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 8, "Question Review")
	pdf.Ln(6)

	for i, r := range result.Results {
		if pdf.GetY() > 260 {
			pdf.AddPage()
		}

		statusText := "Incorrect"
		headerColor := struct{ r, g, b int }{232, 76, 61}
		if r.IsCorrect {
			statusText = "Correct"
			headerColor = struct{ r, g, b int }{46, 204, 113}
		}

		sectionTop := pdf.GetY()
		pdf.SetXY(12, sectionTop)
		pdf.SetFillColor(headerColor.r, headerColor.g, headerColor.b)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(80, 6, fmt.Sprintf("Question %d", i+1), "", 0, "L", true, 0, "")
		pdf.CellFormat(50, 6, statusText, "", 0, "C", true, 0, "")
		pdf.SetFillColor(41, 128, 185)
		pdf.CellFormat(50, 6, r.Question.Domain, "", 1, "R", true, 0, "")

		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(12)
		questionText := r.Question.Prompt
		if len(questionText) > 400 {
			questionText = questionText[:400] + "..."
		}
		pdf.MultiCell(186, 4.5, questionText, "", "L", false)

		userAnswer := choiceListDisplay(r.Question.Choices, r.UserChoiceIDs, ", ")
		correctAnswer := choiceListDisplay(r.Question.Choices, r.CorrectChoiceIDs, ", ")
		pdf.SetX(12)
		pdf.CellFormat(93, 5, fmt.Sprintf("Your Answer: %s", userAnswer), "", 0, "L", false, 0, "")
		pdf.CellFormat(93, 5, fmt.Sprintf("Correct Answer: %s", correctAnswer), "", 1, "L", false, 0, "")

		if fb := detailedFeedbackText(r); fb != "" {
			pdf.SetFont("Arial", "I", 9)
			pdf.SetX(12)
			pdf.MultiCell(186, 4.5, "Feedback: "+fb, "", "L", false)
		}

		pdf.SetFont("Arial", "", 9)
		userSet := make(map[int]struct{})
		for _, id := range r.UserChoiceIDs {
			userSet[id] = struct{}{}
		}

		for _, choice := range r.Question.Choices {
			choiceText := singleChoiceDisplay(choice)
			prefix := "  "
			if _, ok := userSet[choice.ID]; ok {
				prefix = "> You: "
			}
			if choice.IsCorrect {
				prefix = "* Correct: "
			}
			line := fmt.Sprintf("%s%s", prefix, choiceText)
			if len(line) > 150 {
				line = line[:150] + "..."
			}
			pdf.SetX(12)
			pdf.Cell(186, 4, line)
			pdf.Ln(3.5)
		}

		if r.Question.Explanation != "" {
			expl := r.Question.Explanation
			if len(expl) > 400 {
				expl = expl[:400] + "..."
			}
			pdf.SetFont("Arial", "I", 9)
			pdf.SetX(12)
			pdf.MultiCell(186, 4.5, "Explanation: "+expl, "", "L", false)
		}

		if tip := domainTipText(r.Question.Domain, r.IsCorrect); tip != "" {
			pdf.SetFont("Arial", "I", 9)
			pdf.SetX(12)
			pdf.MultiCell(186, 4.5, "Tip: "+tip, "", "L", false)
		}

		sectionBottom := pdf.GetY()
		if sectionBottom-sectionTop < 20 {
			sectionBottom = sectionTop + 20
		}
		pdf.SetDrawColor(210, 210, 210)
		pdf.Rect(10, sectionTop, 190, sectionBottom-sectionTop+4, "D")
		pdf.SetY(sectionBottom + 6)
		addDivider(pdf)
		pdf.Ln(4)
	}

	pdf.SetY(-15)
	pdf.SetFont("Arial", "I", 8)
	pdf.Cell(190, 8, fmt.Sprintf("Generated on %s", time.Now().Format("2006-01-02 15:04:05")))

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}

	return &buf, nil
}

func addDivider(pdf *gofpdf.Fpdf) {
	pdf.SetDrawColor(210, 210, 210)
	pdf.CellFormat(190, 0, "", "B", 1, "", false, 0, "")
}

func singleChoiceDisplay(choice models.Choice) string {
	if choice.Text == "" || choice.Text == choice.Label {
		if choice.Label != "" {
			return choice.Label
		}
		return fmt.Sprintf("Choice %d", choice.ID)
	}
	return fmt.Sprintf("%s. %s", choice.Label, choice.Text)
}

func choiceListDisplay(choices []models.Choice, choiceIDs []int, joiner string) string {
	if len(choiceIDs) == 0 {
		return "None"
	}

	choiceMap := make(map[int]models.Choice, len(choices))
	for _, choice := range choices {
		choiceMap[choice.ID] = choice
	}

	labels := make([]string, 0, len(choiceIDs))
	for _, id := range choiceIDs {
		if choice, ok := choiceMap[id]; ok {
			labels = append(labels, singleChoiceDisplay(choice))
		}
	}

	if len(labels) == 0 {
		return "None"
	}

	return strings.Join(labels, joiner)
}

func detailedFeedbackText(r models.QuestionResult) string {
	if len(r.UserChoiceIDs) == 0 {
		return "Question was left unanswered. Revisit the scenario and map it to PMI guidance before your next attempt."
	}

	if r.IsCorrect {
		return "Strong alignment with PMI expectations - keep reinforcing the principle demonstrated here."
	}

	user := choiceListDisplay(r.Question.Choices, r.UserChoiceIDs, ", ")
	correct := choiceListDisplay(r.Question.Choices, r.CorrectChoiceIDs, ", ")
	return fmt.Sprintf("You chose %s, but the better answer is %s. Review the scenario details and PMI references that support the correct option.", user, correct)
}

func domainTipText(domain string, correct bool) string {
	switch domain {
	case "Project Management Fundamentals":
		if correct {
			return "Keep scanning for stakeholder cues, governance triggers, and ethical obligations; they often differentiate top answers."
		}
		return "Revisit PMBOK 7th domains and the PMI Code of Ethics to sharpen decision patterns."
	case "Predictive Methodologies":
		if correct {
			return "Great predictive instincts - continue drilling earned value and change control scenarios."
		}
		return "Refresh schedule, cost, and baseline management techniques to strengthen predictive reasoning."
	case "Agile Frameworks":
		if correct {
			return "Your agile mindset is sharp - keep practicing servant leadership and flow optimization."
		}
		return "Review Scrum roles, Kanban signals, and Lean principles to clarify adaptive responses."
	case "Business Analysis":
		if correct {
			return "Excellent BA coverage - continue mapping stakeholder needs to value techniques."
		}
		return "Refresh elicitation, prioritization, and change enablement tools to raise accuracy."
	case "Hard Question":
		if correct {
			return "You mastered an exam-level scenario; compare this reasoning with other PMI ECO tasks for stronger recall."
		}
		return "Dissect the scenario carefully: isolate the value driver, governance hook, and PMI ECO domain it reflects."
	default:
		if correct {
			return "Solid work - keep tying outcomes back to PMI guidance."
		}
		return "Re-examine the scenario through PMI best practices to sharpen your intuition."
	}
}
