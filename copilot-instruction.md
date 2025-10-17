# Copilot Instructions — CAPM Mock Exam System (Go + Postgres)

These instructions help GitHub Copilot (and Copilot Chat) generate consistent code, tests, and docs for this project.

> **Stack**: Go (net/http, pgx), Postgres, (planned) Vue 3 + Vite frontend, gofpdf for PDFs.
> **Domain**: CAPM exam simulator (150 items), weighted random selection by popularity, explanations shown on results, PDF report, full persistence.

---

## 1) Project Context (Tell Copilot)

Paste this when starting a session or pin it in a `.copilot/context.md` file:

* We’re building a CAPM mock exam backend in **Go**, with **Postgres** for persistence.
* Core entities: `users`, `questions`, `choices`, `exams`, `exam_questions`, `attempts`, `attempt_answers`.
* Flows:

    1. **Start exam** → create attempt, pick 150 questions with **weighted random** by `popularity_score`.
    2. **Take exam** → fetch ordered questions with choices (A–D).
    3. **Submit** → store answers, compute score, return per-question **explanations** (always visible).
    4. **Download** → generate **PDF report** with score + per-item details.
* All DB writes must be **transactional** when seeding/upserting a question + its choices.
* Idempotent seeding uses a stable `questions.code` (e.g., `CAPM-0001`).

---

## 2) Coding Conventions (Go)

* Use `context.Context` on all repository and handler methods.
* SQL access via `pgx` stdlib (`github.com/jackc/pgx/v5/stdlib`).
* Errors: return wrapped errors, log at the edge (handlers / main), prefer `http.Error` for 4xx/5xx.
* JSON: lower_snake_case or lowerCamelCase consistently; this project uses **lowerCamelCase** in payloads.
* Keep handlers small; push logic into `repository` or small `service` functions.
* For randomness, use a **seeded** `rand.Rand` stored per attempt.

---

## 3) Database Reminders

**Tables** (columns abbreviated):

* `questions(id, code UNIQUE, prompt, explanation, domain, difficulty, popularity_score, timestamps)`
* `choices(id, question_id, label CHAR(1), text, is_correct)`
* `exams(id, title, total_items)`
* `exam_questions(exam_id, question_id, position)`
* `attempts(id, exam_id, user_id, seed, score, started_at, submitted_at)`
* `attempt_answers(attempt_id, question_id, choice_label, is_correct)`

**Indexes**: `questions_popularity_idx`, `exam_questions_exam_pos_idx`, `attempt_answers_attempt_idx`.

**Upsert rules**:

* `questions`: upsert by `code`.
* `choices`: upsert by `(question_id, label)`.

---

## 4) Public API (Backend)

* `POST /api/exams/start`  → `{ title?, total_items? }` → `{ attempt_id, seed }`
* `GET  /api/exams/{attemptId}/questions` → `{ questions: [...] }`
* `POST /api/exams/{attemptId}/submit` → `{ answers: [{question_id, choice_label}] }` → `Result`
* `GET  /api/exams/{attemptId}/report.pdf` → PDF attachment

**Result shape**:

```json
{
  "attempt_id": "...",
  "score": 120,
  "total": 150,
  "items": [
    {
      "question_id": "...",
      "code": "CAPM-0001",
      "prompt": "...",
      "your_answer": "C",
      "correct": true,
      "correct_label": "C",
      "explanation": "Crashing adds resources..."
    }
  ]
}
```

---

## 5) High-Quality Prompt Templates

Copy‑paste these to Copilot Chat as needed.

### 5.1 Handlers & Routing

> **Goal**: Add an endpoint to shuffle choices per attempt and return in stable order.

```
Write a Go handler `GET /api/exams/{attemptId}/questions` that returns questions with choices
shuffled per attempt using the attempt's seed. The shuffle must be stable: same attempt → same order.
Use repository methods to load questions and choices.
Return JSON `{ questions: [{ id, code, prompt, choices: [{ label, text }] }] }`.
```

### 5.2 Weighted Selection Service

> **Goal**: Extract roulette wheel selection into a service with tests.

```
Create `internal/service/selector.go` exposing `PickQuestions(ctx, db, total, seed)`.
Implement roulette-wheel selection without replacement using weight = 0.1 + popularity_score.
Add unit tests covering: stable picks given same seed; popularity bias; capping when pool < total.
```

### 5.3 PDF Report (gofpdf)

> **Goal**: Replace stub with real PDF.

```
Implement `renderPDF(r *Result) ([]byte, error)` using gofpdf.
Layout: Title page (Attempt ID, date, Score), then per-question pages with:
Prompt, Your Answer vs Correct Answer, and Explanation.
Add page numbers and a simple header/footer.
```

### 5.4 Repository Upserts

> **Goal**: Idempotent seeding by `code` with transaction.

```
In `repository.go`, add `CreateOrUpdateQuestionTx(ctx, q, choices)` that upserts a question by `code`
then upserts choices by `(question_id, label)`, all in a transaction. Write tests using a test DB.
```

### 5.5 Frontend (Vue 3 + Vite)

> **Goal**: Thin SPA that calls the backend.

```
Scaffold Vue 3 + Vite app with Tailwind. Pages:
- Home: button “Take Mock Exam” → POST /api/exams/start → route to /exam/:attemptId
- Exam: fetch questions, render cards with radio choices. Local save to localStorage.
- Submit: POST /api/exams/{attemptId}/submit → show score and breakdown, with link to report.pdf
Use Pinia for attempt state. Keep components: QuestionCard.vue, ResultsTable.vue.
```

### 5.6 Test Data Generator

> **Goal**: Generate 200 synthetic questions for dev.

```
Write a small Go CLI `cmd/devseed` that inserts 200 demo questions with choices and random popularity.
Ensure codes are stable (e.g., DEMO-0001..DEMO-0200). Use the same repository upsert API.
```

### 5.7 ECO Quotas

> **Goal**: Ensure the exam matches PMI ECO proportions by domain.

```
Add domain quotas to `PickQuestions`: Fundamentals 36%, Predictive 17%, Agile 20%, Business Analysis 27%.
If a bucket lacks enough items, borrow from the next most-populated bucket.
Write tests for quota satisfaction and graceful fallback.
```

### 5.8 Analytics

> **Goal**: Basic item stats.

```
Add a nightly job that computes per-question difficulty (p-value = % correct) and stores it.
Expose `GET /api/admin/items/stats` for a small admin table.
```

---

## 6) Copilot Chat Tips

* **Pin context**: Paste Section 1 (Project Context) into the chat and pin it.
* **Be specific**: Ask for file paths and function names. E.g., “Create `internal/service/selector.go` with `PickQuestions`.”
* **Iterate**: If output is close but not exact, say “Refactor to use repository, not direct SQL,” or “Return camelCase JSON.”
* **Ask for tests**: “Generate table-driven tests with realistic edge cases.”
* **Review**: Request a diff-style patch when modifying existing files.

---

## 7) File Layout (desired)

```
/cmd/server/main.go
/cmd/devseed/main.go
/internal/database/database.go
/internal/repository/repository.go
/internal/service/selector.go
/internal/http/handlers.go
/internal/pdf/report.go
/web (Vue app)
```

---

## 8) PR Template (paste into `.github/pull_request_template.md`)

```
## What
- Short summary of the change.

## Why
- Link to issue / motivation.

## How
- Key implementation notes (handlers, repo, SQL, tests).

## Screenshots / PDFs
- If applicable.

## Risks & Rollback
- Potential impacts and how to revert.

## Checklist
- [ ] Unit tests pass
- [ ] API schema updated (if needed)
- [ ] Migrations applied locally
```

---

## 9) Common Copilot Chat Requests (copy/paste)

* “Write a Bruno collection for the 4 endpoints with sample bodies.”
* “Write a migration to add `UNIQUE(code)` to `questions` and backfill codes.”
* “Write unit tests for `renderPDF` using a golden file.”
* “Add `OPTIONS` and CORS headers for local Vue dev at `http://localhost:5173`.”
* “Implement pagination for an admin `GET /api/questions?limit=50&cursor=...`.”

---

## 10) Troubleshooting Prompts

* “Fix pq: relation does not exist — generate migration order and guard in code.”
* “Handle Postgres timeouts and add context deadlines on long queries.”
* “Add graceful shutdown with context cancellation and `server.Shutdown`.”

---

## 11) Security & Integrity Hints

* Validate `choice_label` in `A..D`.
* Don’t expose `is_correct` in question payloads.
* Shuffle choices per attempt; store the mapping if you later display letters.
* Rate-limit start/submit endpoints if exposing publicly.

---