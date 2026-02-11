CREATE TABLE IF NOT EXISTS exercises (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    type            TEXT    NOT NULL,
    difficulty      TEXT    NOT NULL DEFAULT 'beginner',
    prompt          TEXT    NOT NULL,
    correct_answer  TEXT    NOT NULL,
    options         TEXT,
    hint            TEXT,
    lesson_id       INTEGER REFERENCES grammar_lessons(id),
    created_at      TEXT    NOT NULL DEFAULT (datetime('now'))
);
