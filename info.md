# CAPM Mock Exam System

This document outlines the current capabilities of the **CAPM Mock Exam System**. The application is written in **Go**, persists data in **PostgreSQL**, and delivers exam-style practice aligned with **PMBOK® Guide 7th Edition** and **PMI ECO 2024**.

---

## 🎯 Purpose
- Provide three practice modes:
  - **Full Mock Exam**: 150-question simulation.
  - **Short Quiz**: 15-question quick drills.
  - **Hard Question Drill**: 30 exam-level scenario questions.
- Bias question selection using popularity scores while maintaining coverage.
- Deliver instant feedback with detailed explanations and targeted study tips.
- Generate polished PDF reports for offline review.
- Persist exam attempts for history lookups and analytics.

---

## 🛠️ Tech Stack
- **Backend**: Go (net/http, pgx, gorilla/mux)
- **Database**: PostgreSQL
- **PDF Generation**: gofpdf
- **Frontend**: Server-rendered HTML + Bootstrap (Vue optional)

---

## ⚙️ Core Features

### Question Management
- Questions stored with:
  - `prompt`
  - `choices` (A/B/C/D with boolean flag)
  - `domain` (e.g., Project Fundamentals, Predictive, Agile, Business Analysis, Hard Question)
  - `explanation`
  - `popularity_score`

### Exam Modes
- **Mock Exam** (`/api/exams/start`): 150 weighted questions across all domains.
- **Short Quiz** (`/api/quiz/start`): 15-question mini session.
- **Hard Drill** (`/api/hard/start`): 20 scenario-heavy CAPM problems with narrative answer options.
- **Team & Motivation Theories** (`/api/team-motivation/questions`): 20-question drill covering Tuckman, Maslow, Herzberg, McGregor, McClelland, Vroom, and Theory Z scenarios.

### Exam Flow
1. **Launch**: User submits name/email; backend creates attempt and assigns seed.
2. **Retrieve Questions**: Weighted random selection (domain-filtered for hard drill).
3. **Submit**: System grades, stores answers, records score and completion timestamp.
4. **Results View**:
   - Detailed breakdown by question with explanations and custom tips.
   - Filter by All / Correct / Incorrect.
   - Shows your answer versus correct answer.
   - Provides domain-specific improvement guidance.
5. **PDF Report**:
   - Professionally formatted summary.
   - Domain table with % performance.
   - Per-question cards with feedback, explanations, and study tips.

### History Tracking
- Attempts stored per user; accessible via `/api/users/{userId}/attempts`.
- Landing page includes **View History** modal (Short Quiz / Mock Exam / Hard Drill badges).

---

## 🔑 API Endpoints
- `POST /api/exams/start` – Start mock exam.
- `POST /api/quiz/start` – Start short quiz.
- `POST /api/hard/start` – Start hard drill.
- `GET /api/exams/{attemptId}/questions`
- `POST /api/exams/{attemptId}/submit`
- `GET /api/exams/{attemptId}/results`
- `GET /api/exams/{attemptId}/report.pdf`
- `GET /api/users/{userId}/attempts`

---

## 🧪 Hard Question Drill
- 30 expert-level CAPM scenarios.
- Long, exam-style prompts with letter-only answer choices.
- Explanations and tips emphasize value delivery, governance, risk, and hybrid tailoring per PMI ECO.

---

## 📄 Result Experience
- In-app review highlights:
  - Question prompt and domain.
  - Your answer vs Correct answer.
  - Feedback message with improvement guidance.
  - Links explanations and domain-specific tips.
- PDF mirrors the in-app detail with enhanced layout and section dividers.

---

## 🗄️ Data Model Snapshot
- `users`
- `questions`, `choices`
- `exams` (Mock, Short Quiz, Hard Drill)
- `attempts`, `attempt_answers`

---

## 🛠️ Runbook
```
export PG_URL=postgres://postgres:postgres@localhost:5432/capm?sslmode=disable
make db-migrate
make db-seed
make restart   # stop + start server
```
Access the app at `http://localhost:8080`.

---

## 📈 Roadmap Ideas
- Timed exams and pause/resume logic.
- Advanced analytics dashboards (domain mastery trends, question difficulty).
- Admin UI for question curation and popularity tuning.
- Authentication for multi-user tracking.

---

## 📥 Sample PDF Output
- Score summary with pass/fail.
- Domain performance table.
- Question-by-question cards showing prompt, your answer, correct answer, feedback, explanation, and study tip.

---

This platform aims to **mirror real CAPM exam pressure** while providing rich feedback so every attempt becomes a targeted study session.
