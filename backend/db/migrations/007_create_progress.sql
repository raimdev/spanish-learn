CREATE TABLE IF NOT EXISTS user_flashcard_progress (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id          INTEGER NOT NULL REFERENCES users(id),
    flashcard_id     INTEGER NOT NULL REFERENCES flashcards(id),
    repetitions      INTEGER NOT NULL DEFAULT 0,
    easiness_factor  REAL    NOT NULL DEFAULT 2.5,
    interval_days    INTEGER NOT NULL DEFAULT 0,
    next_review_date TEXT    NOT NULL DEFAULT (datetime('now')),
    last_reviewed_at TEXT,
    UNIQUE(user_id, flashcard_id)
);

CREATE TABLE IF NOT EXISTS review_log (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id          INTEGER NOT NULL REFERENCES users(id),
    flashcard_id     INTEGER NOT NULL REFERENCES flashcards(id),
    quality          INTEGER NOT NULL,
    reviewed_at      TEXT    NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS user_exercise_progress (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id      INTEGER NOT NULL REFERENCES users(id),
    exercise_id  INTEGER NOT NULL REFERENCES exercises(id),
    is_correct   INTEGER NOT NULL,
    user_answer  TEXT,
    completed_at TEXT    NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS user_grammar_progress (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id     INTEGER NOT NULL REFERENCES users(id),
    lesson_id   INTEGER NOT NULL REFERENCES grammar_lessons(id),
    completed   INTEGER NOT NULL DEFAULT 0,
    completed_at TEXT,
    UNIQUE(user_id, lesson_id)
);
