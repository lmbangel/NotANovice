# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a **Daily Discipleship Quiz** application built in Go, designed as a Christian discipleship tool inspired by Hebrews 5:12-14. The application helps believers grow spiritually through daily Bible-based quiz questions, tracking progress via a leaderboard system.

## Architecture

- **HTTP API**: Chi router-based REST API (`main.go`)
- **Database**: SQLite with goose migrations (`quiz.db`)
- **Code Generation**: Uses `sqlc` for type-safe database access
- **AI Integration**: Local LLM agents for quiz generation (`pkg/agents/`)

### Key Components

- `main.go` - HTTP server with REST endpoints
- `internal/db/` - Generated database models and queries via sqlc
- `pkg/agents/` - AI agent interface for LLM integration (quiz generation)
- `cmd/quiz_agent.go` - CLI tool for AI-powered quiz generation
- `migrations/` - Database schema migrations managed by goose
- `queries/quiz.sql` - SQL queries for sqlc code generation

## Development Commands

### Running the Application
```bash
make serve
# or
go run main.go
```

### Database Management
```bash
# Run pending migrations
make goose-up

# Rollback last migration  
make goose-down

# Check migration status
make goose-status
```

### Code Generation
```bash
# Generate Go code from SQL queries (after modifying queries/quiz.sql)
sqlc generate
```

## Database Schema

Core entities:
- **users** - User accounts with username/email
- **questions** - Quiz questions with answers
- **attempts** - User quiz attempts with correctness tracking
- **leader_board** - User scores and leaderboard rankings
- **answers** - Answer options/choices

## API Endpoints

- `GET /health` - Health check
- `GET /questions` - Retrieve quiz questions
- Additional endpoints implemented in `main.go`

## Development Workflow

1. Modify SQL queries in `queries/quiz.sql`
2. Run `sqlc generate` to update Go models
3. Create database migrations in `migrations/` if schema changes
4. Run `make goose-up` to apply migrations
5. Test with `make serve`

## Dependencies

- Chi router for HTTP handling
- SQLite with mattn/go-sqlite3 driver
- goose for database migrations
- sqlc for type-safe database access
- Local LLM integration for quiz generation