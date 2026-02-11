CREATE TABLE IF NOT EXISTS vocabulary_words (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    spanish         TEXT    NOT NULL,
    english         TEXT    NOT NULL,
    part_of_speech  TEXT,
    category        TEXT,
    example_sentence TEXT,
    created_at      TEXT    NOT NULL DEFAULT (datetime('now'))
);
