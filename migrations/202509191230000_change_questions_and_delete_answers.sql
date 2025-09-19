-- +goose Up
-- 1. Rebuild questions table with new schema
CREATE TABLE questions_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1,
    a_answer TEXT NOT NULL,
    b_answer TEXT NOT NULL,
    c_answer TEXT NOT NULL,
    d_answer TEXT
);

-- Copy old data into new table (map old "answer" → "correct_answer")
INSERT INTO questions_new (id, question, correct_answer, timestamp, is_active)
SELECT id, question, answer, timestamp, is_active FROM questions;

-- Replace old table with new one
DROP TABLE questions;
ALTER TABLE questions_new RENAME TO questions;

-- 2. Drop answers table (no longer needed)
DROP TABLE IF EXISTS answers;

-- +goose Down
-- 1. Recreate old questions table
CREATE TABLE questions_old (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1
);

-- Copy back compatible data (map "correct_answer" → "answer")
INSERT INTO questions_old (id, question, answer, timestamp, is_active)
SELECT id, question, correct_answer, timestamp, is_active FROM questions;

-- Replace new table with old one
DROP TABLE questions;
ALTER TABLE questions_old RENAME TO questions;

-- 2. Recreate answers table
CREATE TABLE answers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    answer TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1
);
