import { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { get, post } from '../api/client';

const TYPE_LABELS = {
  fill_blank: 'Fill in the Blank',
  translation: 'Translation',
  conjugation: 'Conjugation',
};

export default function ExercisePage() {
  const { id } = useParams();
  const [exercise, setExercise] = useState(null);
  const [answer, setAnswer] = useState('');
  const [result, setResult] = useState(null);

  useEffect(() => {
    get(`/api/exercises/${id}`).then(setExercise).catch(console.error);
    setAnswer('');
    setResult(null);
  }, [id]);

  async function handleSubmit(e) {
    e.preventDefault();
    const data = await post(`/api/exercises/${id}/submit`, { answer });
    setResult(data);
  }

  function handleKeyDown(e) {
    if (e.key === 'Enter' && !result) {
      handleSubmit(e);
    }
  }

  if (!exercise) return <div className="loading">Loading...</div>;

  return (
    <div className="exercise-page">
      <Link to="/exercises" className="back-link">Back to Exercises</Link>

      <div className="exercise-detail">
        <div className="exercise-badges">
          <span className={`badge badge-type-${exercise.type}`}>
            {TYPE_LABELS[exercise.type] || exercise.type}
          </span>
          <span className={`badge badge-diff-${exercise.difficulty}`}>
            {exercise.difficulty}
          </span>
        </div>

        <div className="exercise-prompt-large">{exercise.prompt}</div>

        {exercise.hint && !result && (
          <div className="exercise-hint">Hint: {exercise.hint}</div>
        )}

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Your Answer</label>
            <input
              type="text"
              value={answer}
              onChange={(e) => setAnswer(e.target.value)}
              onKeyDown={handleKeyDown}
              disabled={!!result}
              autoFocus
              placeholder="Type your answer..."
            />
          </div>

          {!result && (
            <button type="submit" className="btn-primary" disabled={!answer.trim()}>
              Submit Answer
            </button>
          )}
        </form>

        {result && (
          <div className={`exercise-result ${result.correct ? 'result-correct' : 'result-incorrect'}`}>
            <div className="result-icon">{result.correct ? 'Correct!' : 'Incorrect'}</div>
            {!result.correct && (
              <div className="result-answer">
                Correct answer: <strong>{result.correct_answer}</strong>
              </div>
            )}
            <div className="result-actions">
              <Link to="/exercises" className="btn-secondary">Back to Exercises</Link>
              <button
                onClick={() => { setAnswer(''); setResult(null); }}
                className="btn-primary"
              >
                Try Again
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
