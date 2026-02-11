import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { get } from '../api/client';

const TYPE_LABELS = {
  fill_blank: 'Fill in the Blank',
  translation: 'Translation',
  conjugation: 'Conjugation',
};

const DIFFICULTY_LABELS = {
  beginner: 'Beginner',
  intermediate: 'Intermediate',
  advanced: 'Advanced',
};

export default function ExercisesListPage() {
  const [exercises, setExercises] = useState([]);
  const [filterType, setFilterType] = useState('');
  const [filterDifficulty, setFilterDifficulty] = useState('');

  useEffect(() => {
    const params = new URLSearchParams();
    if (filterType) params.set('type', filterType);
    if (filterDifficulty) params.set('difficulty', filterDifficulty);
    const query = params.toString();
    get(`/api/exercises${query ? `?${query}` : ''}`).then(setExercises).catch(console.error);
  }, [filterType, filterDifficulty]);

  return (
    <div className="exercises-list">
      <h1>Exercises</h1>
      <p className="subtitle">Practice your Spanish with various exercise types</p>

      <div className="filters">
        <div className="form-group">
          <label>Type</label>
          <select value={filterType} onChange={(e) => setFilterType(e.target.value)}>
            <option value="">All types</option>
            <option value="fill_blank">Fill in the Blank</option>
            <option value="translation">Translation</option>
            <option value="conjugation">Conjugation</option>
          </select>
        </div>
        <div className="form-group">
          <label>Difficulty</label>
          <select value={filterDifficulty} onChange={(e) => setFilterDifficulty(e.target.value)}>
            <option value="">All levels</option>
            <option value="beginner">Beginner</option>
            <option value="intermediate">Intermediate</option>
            <option value="advanced">Advanced</option>
          </select>
        </div>
      </div>

      <div className="exercise-grid">
        {exercises.map((ex) => (
          <Link to={`/exercises/${ex.id}`} key={ex.id} className="exercise-card">
            <div className="exercise-badges">
              <span className={`badge badge-type-${ex.type}`}>
                {TYPE_LABELS[ex.type] || ex.type}
              </span>
              <span className={`badge badge-diff-${ex.difficulty}`}>
                {DIFFICULTY_LABELS[ex.difficulty] || ex.difficulty}
              </span>
            </div>
            <p className="exercise-prompt">{ex.prompt}</p>
          </Link>
        ))}
        {exercises.length === 0 && <p>No exercises found.</p>}
      </div>
    </div>
  );
}
