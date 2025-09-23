-- name: CreateQuestion :one
INSERT INTO questions (question, correct_answer, a_answer, b_answer, c_answer, d_answer)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetQuestion :one
SELECT * FROM questions WHERE id = ?;

-- name: GetQuestions :many
SELECT * FROM questions;

-- name: CreateUser :one
INSERT INTO users (username, email) VALUES (?, ?)
RETURNING *;
-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users where username = ?;

-- name: UpdateUser :one
UPDATE users SET username = ?, email = ? where id = ?
RETURNING *;

-- name: RecordAttempt :one
INSERT INTO attempts (user_id, quiz_id, answer, is_correct)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetAttempts :many
SELECT * FROM attempts;

-- name: GetAttemptByID :one
SELECT * FROM attempts WHERE id = ?;

-- name: GetAttemptsByUserID :many
SELECT * FROM attempts WHERE user_id = ?;

-- name: GetQuizes :many
SELECT quiz.id as quiz_id , questions.* 
FROM quiz 
LEFT JOIN questions on questions.id = quiz.q_id
WHERE quiz.is_active = 1;

-- name: GetQuizByID :one
SELECT quiz.id as quiz_id , questions.* 
FROM quiz 
LEFT JOIN questions on questions.id = quiz.q_id 
WHERE quiz.id = ?;

-- name: GetQuizOfTheDay :one
SELECT quiz.id as quiz_id , questions.* 
FROM quiz 
LEFT JOIN questions on questions.id = quiz.q_id
WHERE DATE(quiz.date) = DATE('now');

-- name: UpdateLeaderboard :exec
INSERT INTO leader_board (user_id, total_score)
VALUES (?, ?)
ON CONFLICT(user_id) DO UPDATE
SET total_score = total_score + excluded.total_score,
    last_updated = CURRENT_TIMESTAMP;

-- name: GetLeaderBoard :many
SELECT u.username, u.email, l.* FROM 
leader_board l
Left Join users u ON u.id = l.user_id;

-- name: GetLeaderBoardByUserID :one
SELECT u.username, u.email, l.* FROM 
leader_board l
Left Join users u ON u.id = l.user_id
WHERE l.user_id = ?;