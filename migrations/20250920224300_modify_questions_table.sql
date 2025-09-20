-- +goose Up
-- 1. Rebuild questions table with new schema (d_answer is NOT NULL)
CREATE TABLE questions_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1,
    a_answer TEXT NOT NULL,
    b_answer TEXT NOT NULL,
    c_answer TEXT NOT NULL,
    d_answer TEXT NOT NULL
);

-- Copy old data into new table
-- If some rows had NULL d_answer, set a default like '' or 'N/A'
INSERT INTO questions_new (id, question, correct_answer, timestamp, is_active, a_answer, b_answer, c_answer, d_answer)
SELECT 
    id,
    question,
    correct_answer,
    timestamp,
    is_active,
    a_answer,
    b_answer,
    c_answer,
    COALESCE(d_answer, '') -- or 'N/A'
FROM questions;

-- Replace old table with new one
DROP TABLE questions;
ALTER TABLE questions_new RENAME TO questions;


-- +goose Down
-- Roll back to previous version (d_answer nullable)
CREATE TABLE questions_old (
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

-- Copy back data
INSERT INTO questions_old (id, question, correct_answer, timestamp, is_active, a_answer, b_answer, c_answer, d_answer)
SELECT id, question, correct_answer, timestamp, is_active, a_answer, b_answer, c_answer, d_answer
FROM questions;

-- Replace table
DROP TABLE questions;
ALTER TABLE questions_old RENAME TO questions;