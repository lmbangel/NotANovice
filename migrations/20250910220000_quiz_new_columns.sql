-- +goose Up
ALTER TABLE quiz ADD COLUMN options_json TEXT NOT NULL DEFAULT '{}';

-- +goose Down
CREATE TABLE quiz_tmp (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    q_id INTEGER NOT NULL,
    a_id INTEGER NOT NULL,
    date DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1,
    FOREIGN KEY (q_id) REFERENCES questions(id),
    FOREIGN KEY (a_id) REFERENCES answers(id)
);

INSERT INTO quiz_tmp (id, q_id, a_id, date, is_active)
SELECT id, q_id, a_id, date, is_active FROM quiz;

DROP TABLE quiz;
ALTER TABLE quiz_tmp RENAME TO quiz;
