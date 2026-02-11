import { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { get, post } from '../api/client';

export default function GrammarDetailPage() {
  const { slug } = useParams();
  const [lesson, setLesson] = useState(null);

  useEffect(() => {
    get(`/api/grammar/${slug}`).then(setLesson).catch(console.error);
  }, [slug]);

  async function markComplete() {
    await post(`/api/grammar/${lesson.id}/complete`, {});
    setLesson({ ...lesson, completed: true });
  }

  if (!lesson) return <div className="loading">Loading...</div>;

  return (
    <div className="grammar-detail">
      <Link to="/grammar" className="back-link">Back to Grammar Lessons</Link>

      <div className="lesson-header">
        <h1>{lesson.title}</h1>
        {lesson.completed ? (
          <span className="badge badge-success">Completed</span>
        ) : (
          <button onClick={markComplete} className="btn-primary">Mark as Complete</button>
        )}
      </div>

      <div
        className="lesson-content"
        dangerouslySetInnerHTML={{ __html: lesson.content }}
      />

      {lesson.examples && lesson.examples.length > 0 && (
        <div className="examples-section">
          <h2>Examples</h2>
          <div className="examples-list">
            {lesson.examples.map((ex) => (
              <div key={ex.id} className="example-card">
                <div className="example-spanish">{ex.spanish}</div>
                <div className="example-english">{ex.english}</div>
                {ex.explanation && (
                  <div className="example-explanation">{ex.explanation}</div>
                )}
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
