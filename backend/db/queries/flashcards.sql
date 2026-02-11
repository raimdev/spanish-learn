-- name: ListFlashcardDecks :many
SELECT * FROM flashcard_decks ORDER BY id;

-- name: GetFlashcardDeck :one
SELECT * FROM flashcard_decks WHERE id = ? LIMIT 1;

-- name: CountCardsInDeck :one
SELECT COUNT(*) FROM flashcards WHERE deck_id = ?;

-- name: GetDueFlashcards :many
SELECT f.id, f.deck_id, f.front, f.back, f.created_at
FROM flashcards f
LEFT JOIN user_flashcard_progress ufp ON ufp.flashcard_id = f.id AND ufp.user_id = ?1
WHERE (ufp.next_review_date IS NULL OR ufp.next_review_date <= datetime('now'))
ORDER BY ufp.next_review_date ASC
LIMIT ?2;

-- name: GetDueFlashcardsByDeck :many
SELECT f.id, f.deck_id, f.front, f.back, f.created_at
FROM flashcards f
LEFT JOIN user_flashcard_progress ufp ON ufp.flashcard_id = f.id AND ufp.user_id = ?1
WHERE f.deck_id = ?2 AND (ufp.next_review_date IS NULL OR ufp.next_review_date <= datetime('now'))
ORDER BY ufp.next_review_date ASC
LIMIT ?3;

-- name: GetFlashcardProgress :one
SELECT * FROM user_flashcard_progress WHERE user_id = ? AND flashcard_id = ? LIMIT 1;

-- name: UpsertFlashcardProgress :exec
INSERT INTO user_flashcard_progress (user_id, flashcard_id, repetitions, easiness_factor, interval_days, next_review_date, last_reviewed_at)
VALUES (?1, ?2, ?3, ?4, ?5, ?6, datetime('now'))
ON CONFLICT(user_id, flashcard_id) DO UPDATE SET
    repetitions = ?3,
    easiness_factor = ?4,
    interval_days = ?5,
    next_review_date = ?6,
    last_reviewed_at = datetime('now');

-- name: CreateReviewLog :exec
INSERT INTO review_log (user_id, flashcard_id, quality) VALUES (?, ?, ?);

-- name: CountDueFlashcards :one
SELECT COUNT(*)
FROM flashcards f
LEFT JOIN user_flashcard_progress ufp ON ufp.flashcard_id = f.id AND ufp.user_id = ?
WHERE ufp.next_review_date IS NULL OR ufp.next_review_date <= datetime('now');

-- name: GetReviewHistory :many
SELECT rl.*, f.front, f.back FROM review_log rl
JOIN flashcards f ON f.id = rl.flashcard_id
WHERE rl.user_id = ?
ORDER BY rl.reviewed_at DESC
LIMIT ?;

-- name: GetFlashcard :one
SELECT * FROM flashcards WHERE id = ? LIMIT 1;
