import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { get } from '../api/client';
import { useAuth } from '../context/AuthContext';

export default function DashboardPage() {
  const { user } = useAuth();
  const [summary, setSummary] = useState(null);

  useEffect(() => {
    get('/api/progress/summary').then(setSummary).catch(console.error);
  }, []);

  if (!summary) return <div className="loading">Loading...</div>;

  return (
    <div className="dashboard">
      <h1>Welcome, {user?.username}!</h1>
      <p className="subtitle">Your Spanish learning dashboard</p>

      <div className="stats-grid">
        <Link to="/grammar" className="stat-card">
          <div className="stat-icon">📖</div>
          <div className="stat-info">
            <div className="stat-value">{summary.grammar.completed} / {summary.grammar.total}</div>
            <div className="stat-label">Grammar Lessons</div>
          </div>
        </Link>

        <Link to="/exercises" className="stat-card">
          <div className="stat-icon">✏️</div>
          <div className="stat-info">
            <div className="stat-value">{summary.exercises.completed} / {summary.exercises.total}</div>
            <div className="stat-label">Exercises Done</div>
          </div>
        </Link>

        <Link to="/flashcards" className="stat-card highlight">
          <div className="stat-icon">🃏</div>
          <div className="stat-info">
            <div className="stat-value">{summary.flashcards.due}</div>
            <div className="stat-label">Cards Due for Review</div>
          </div>
        </Link>
      </div>

      <div className="quick-actions">
        <h2>Quick Start</h2>
        <div className="action-grid">
          <Link to="/grammar" className="action-card">
            <h3>Study Grammar</h3>
            <p>Learn Spanish grammar rules and patterns</p>
          </Link>
          <Link to="/vocabulary" className="action-card">
            <h3>Practice Vocabulary</h3>
            <p>Test your knowledge of Spanish words</p>
          </Link>
          <Link to="/exercises" className="action-card">
            <h3>Do Exercises</h3>
            <p>Fill-in-the-blank, translation, conjugation</p>
          </Link>
          <Link to="/flashcards" className="action-card">
            <h3>Review Flashcards</h3>
            <p>Spaced repetition for long-term memory</p>
          </Link>
        </div>
      </div>
    </div>
  );
}
