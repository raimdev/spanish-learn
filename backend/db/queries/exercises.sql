-- name: ListExercises :many
SELECT id, type, difficulty, prompt, hint, lesson_id, created_at FROM exercises ORDER BY id;

-- name: ListExercisesByType :many
SELECT id, type, difficulty, prompt, hint, lesson_id, created_at FROM exercises WHERE type = ? ORDER BY id;

-- name: ListExercisesByDifficulty :many
SELECT id, type, difficulty, prompt, hint, lesson_id, created_at FROM exercises WHERE difficulty = ? ORDER BY id;

-- name: ListExercisesByTypeAndDifficulty :many
SELECT id, type, difficulty, prompt, hint, lesson_id, created_at FROM exercises WHERE type = ? AND difficulty = ? ORDER BY id;

-- name: GetExercise :one
SELECT * FROM exercises WHERE id = ? LIMIT 1;

-- name: RecordExerciseAttempt :exec
INSERT INTO user_exercise_progress (user_id, exercise_id, is_correct, user_answer) VALUES (?, ?, ?, ?);

-- name: CountExercisesCompleted :one
SELECT COUNT(DISTINCT exercise_id) FROM user_exercise_progress WHERE user_id = ?;

-- name: CountTotalExercises :one
SELECT COUNT(*) FROM exercises;

-- name: GetExerciseHistory :many
SELECT uep.*, e.type, e.prompt FROM user_exercise_progress uep
JOIN exercises e ON e.id = uep.exercise_id
WHERE uep.user_id = ?
ORDER BY uep.completed_at DESC
LIMIT ?;
