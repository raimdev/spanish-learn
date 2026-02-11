import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { get } from '../api/client';

export default function GrammarListPage() {
  const [lessons, setLessons] = useState([]);

  useEffect(() => {
    get('/api/grammar').then(setLessons).catch(console.error);
  }, []);

  return (
    <div className="grammar-list">
      <h1>Grammar Lessons</h1>
      <p className="subtitle">Learn the rules and patterns of Spanish grammar</p>

      <div className="lesson-list">
        {lessons.map((lesson) => (
          <Link to={`/grammar/${lesson.slug}`} key={lesson.id} className="lesson-card">
            <div className="lesson-status">
              {lesson.completed ? (
                <span className="badge badge-success">Completed</span>
              ) : (
                <span className="badge badge-pending">Not started</span>
              )}
            </div>
            <h3>{lesson.title}</h3>
            <p>{lesson.summary}</p>
          </Link>
        ))}
      </div>
    </div>
  );
}
