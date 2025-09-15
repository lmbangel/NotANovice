-- name: CreateQuestion :one
INSERT INTO questions (question, answer) VALUES (?, ?)
RETURNING *;

-- name: GetQuestion :one
SELECT * FROM questions WHERE id = ?;

-- name: CreateUser :one
INSERT INTO users (username, email) VALUES (?, ?)
RETURNING *;
-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users where username = ?;

-- name: RecordAttempt :one
INSERT INTO attempts (user_id, quiz_id, answer, is_correct)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetAttempts :many
SELECT * FROM attempts;

-- name: UpdateLeaderboard :exec
INSERT INTO leader_board (user_id, total_score)
VALUES (?, ?)
ON CONFLICT(user_id) DO UPDATE
SET total_score = total_score + excluded.total_score,
    last_updated = CURRENT_TIMESTAMP;
