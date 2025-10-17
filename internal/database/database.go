package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New() (*DB, error) {
	pgURL := os.Getenv("PG_URL")
	if pgURL == "" {
		pgURL = "postgres://postgres:postgres@localhost:5432/capm?sslmode=disable"
	}

	pool, err := pgxpool.New(context.Background(), pgURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) CreateTables(ctx context.Context) error {
	queries := []string{
		// Users table
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email VARCHAR(255) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Questions table
		`CREATE TABLE IF NOT EXISTS questions (
			id SERIAL PRIMARY KEY,
			prompt TEXT NOT NULL,
			domain VARCHAR(100) NOT NULL,
			popularity_score DECIMAL(3,2) DEFAULT 1.0,
			explanation TEXT NOT NULL,
			is_multi_select BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Choices table
		`CREATE TABLE IF NOT EXISTS choices (
			id SERIAL PRIMARY KEY,
			question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
			text TEXT NOT NULL,
			is_correct BOOLEAN DEFAULT FALSE,
			label CHAR(1) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		`ALTER TABLE questions ADD COLUMN IF NOT EXISTS is_multi_select BOOLEAN NOT NULL DEFAULT FALSE`,

		// Exams table
		`CREATE TABLE IF NOT EXISTS exams (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255) NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Exam questions (junction table)
		`CREATE TABLE IF NOT EXISTS exam_questions (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			exam_id UUID REFERENCES exams(id) ON DELETE CASCADE,
			question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
			position INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Attempts table
		`CREATE TABLE IF NOT EXISTS attempts (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			exam_id UUID REFERENCES exams(id) ON DELETE CASCADE,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			seed BIGINT NOT NULL,
			score INTEGER,
			max_score INTEGER NOT NULL DEFAULT 150,
			started_at TIMESTAMP DEFAULT NOW(),
			ended_at TIMESTAMP
		)`,

		// Attempt answers table
		`CREATE TABLE IF NOT EXISTS attempt_answers (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			attempt_id UUID REFERENCES attempts(id) ON DELETE CASCADE,
			question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
			choice_id INTEGER REFERENCES choices(id) ON DELETE CASCADE,
			is_correct BOOLEAN,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Indexes
		`CREATE INDEX IF NOT EXISTS idx_questions_domain ON questions(domain)`,
		`CREATE INDEX IF NOT EXISTS idx_questions_popularity ON questions(popularity_score DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_attempts_user_id ON attempts(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_attempt_answers_attempt_id ON attempt_answers(attempt_id)`,
	}

	for _, query := range queries {
		if _, err := db.Pool.Exec(ctx, query); err != nil {
			return fmt.Errorf("error executing query: %v", err)
		}
	}

	return nil
}
