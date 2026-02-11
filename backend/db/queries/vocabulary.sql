-- name: ListVocabulary :many
SELECT * FROM vocabulary_words ORDER BY category, spanish;

-- name: ListVocabularyByCategory :many
SELECT * FROM vocabulary_words WHERE category = ? ORDER BY spanish;

-- name: ListVocabularyCategories :many
SELECT DISTINCT category FROM vocabulary_words WHERE category IS NOT NULL ORDER BY category;

-- name: GetRandomVocabulary :many
SELECT * FROM vocabulary_words ORDER BY RANDOM() LIMIT ?;

-- name: GetRandomVocabularyByCategory :many
SELECT * FROM vocabulary_words WHERE category = ? ORDER BY RANDOM() LIMIT ?;
