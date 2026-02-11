CREATE TABLE IF NOT EXISTS grammar_lessons (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT    NOT NULL,
    slug        TEXT    NOT NULL UNIQUE,
    summary     TEXT    NOT NULL,
    content     TEXT    NOT NULL,
    sort_order  INTEGER NOT NULL DEFAULT 0,
    created_at  TEXT    NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS grammar_examples (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    lesson_id   INTEGER NOT NULL REFERENCES grammar_lessons(id),
    spanish     TEXT    NOT NULL,
    english     TEXT    NOT NULL,
    explanation TEXT,
    sort_order  INTEGER NOT NULL DEFAULT 0
);
