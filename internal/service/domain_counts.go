package service

import "capm-exam-system/internal/models"

type domainQuota struct {
	domain string
	count  int
}

type attemptBlueprint struct {
	domainCounts []domainQuota
	hardCount    int
}

var examBlueprint = attemptBlueprint{
	domainCounts: []domainQuota{
		{domain: "Project Management Fundamentals", count: 47},
		{domain: "Predictive Methodologies", count: 22},
		{domain: "Agile Frameworks", count: 26},
		{domain: "Business Analysis", count: 35},
	},
	hardCount: 20,
}

var pmpBlueprint = attemptBlueprint{
	domainCounts: []domainQuota{
		{domain: "People", count: 63},
		{domain: "Process", count: 75},
		{domain: "Business Environment", count: 12},
	},
}

var quizBlueprint = attemptBlueprint{
	domainCounts: []domainQuota{
		{domain: "Project Management Fundamentals", count: 4},
		{domain: "Predictive Methodologies", count: 3},
		{domain: "Agile Frameworks", count: 3},
		{domain: "Business Analysis", count: 3},
	},
	hardCount: 2,
}

func blueprintForExam(examName string, questionCount int) attemptBlueprint {
	switch examName {
	case defaultExamName:
		if questionCount <= 20 {
			return quizBlueprint
		}
		return examBlueprint
	case pmpExamName:
		return pmpBlueprint
	case hardExamName:
		return attemptBlueprint{hardCount: questionCount}
	default:
		return attemptBlueprint{domainCounts: []domainQuota{{domain: "Project Management Fundamentals", count: questionCount}}}
	}
}

func sumQuota(quota attemptBlueprint) int {
	total := quota.hardCount
	for _, dq := range quota.domainCounts {
		total += dq.count
	}
	return total
}

func questionDomains(question models.QuestionWithChoices) string {
	return question.Question.Domain
}
