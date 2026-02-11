import { useState, useEffect } from 'react';
import { get } from '../api/client';

export default function ProgressPage() {
  const [summary, setSummary] = useState(null);
  const [exerciseHistory, setExerciseHistory] = useState([]);
  const [reviewHistory, setReviewHistory] = useState([]);

  useEffect(() => {
    get('/api/progress/summary').then(setSummary).catch(console.error);
    get('/api/progress/exercises').then(setExerciseHistory).catch(console.error);
    get('/api/progress/reviews').then(setReviewHistory).catch(console.error);
  }, []);

  if (!summary) return <div className="loading">Loading...</div>;

  const exerciseAccuracy = exerciseHistory.length > 0
    ? Math.round((exerciseHistory.filter(e => e.is_correct).length / exerciseHistory.length) * 100)
    : 0;

  return (
    <div className="progress-page">
      <h1>Your Progress</h1>
      <p className="subtitle">Track your Spanish learning journey</p>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-value">{summary.grammar.completed} / {summary.grammar.total}</div>
          <div className="stat-label">Grammar Lessons Completed</div>
          <div className="progress-bar">
            <div
              className="progress-fill"
              style={{ width: `${summary.grammar.total > 0 ? (summary.grammar.completed / summary.grammar.total) * 100 : 0}%` }}
            />
          </div>
        </div>

        <div className="stat-card">
          <div className="stat-value">{summary.exercises.completed} / {summary.exercises.total}</div>
          <div className="stat-label">Exercises Attempted</div>
          <div className="progress-bar">
            <div
              className="progress-fill"
              style={{ width: `${summary.exercises.total > 0 ? (summary.exercises.completed / summary.exercises.total) * 100 : 0}%` }}
            />
          </div>
        </div>

        <div className="stat-card">
          <div className="stat-value">{exerciseAccuracy}%</div>
          <div className="stat-label">Exercise Accuracy</div>
        </div>

        <div className="stat-card">
          <div className="stat-value">{summary.flashcards.due}</div>
          <div className="stat-label">Flashcards Due</div>
        </div>
      </div>

      <div className="history-section">
        <h2>Recent Exercise Attempts</h2>
        {exerciseHistory.length === 0 ? (
          <p className="empty-state">No exercises attempted yet.</p>
        ) : (
          <table className="history-table">
            <thead>
              <tr>
                <th>Exercise</th>
                <th>Type</th>
                <th>Result</th>
                <th>Your Answer</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {exerciseHistory.slice(0, 20).map((item, i) => (
                <tr key={i}>
                  <td>{item.prompt}</td>
                  <td>{item.type}</td>
                  <td>
                    <span className={`badge ${item.is_correct ? 'badge-success' : 'badge-error'}`}>
                      {item.is_correct ? 'Correct' : 'Incorrect'}
                    </span>
                  </td>
                  <td>{item.user_answer || '-'}</td>
                  <td>{new Date(item.completed_at).toLocaleDateString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>

      <div className="history-section">
        <h2>Recent Flashcard Reviews</h2>
        {reviewHistory.length === 0 ? (
          <p className="empty-state">No flashcards reviewed yet.</p>
        ) : (
          <table className="history-table">
            <thead>
              <tr>
                <th>Card</th>
                <th>Answer</th>
                <th>Rating</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {reviewHistory.slice(0, 20).map((item, i) => (
                <tr key={i}>
                  <td>{item.front}</td>
                  <td>{item.back}</td>
                  <td>{item.quality}/5</td>
                  <td>{new Date(item.reviewed_at).toLocaleDateString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
}
