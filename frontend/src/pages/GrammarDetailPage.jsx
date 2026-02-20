import { useState, useEffect, useRef, useCallback } from 'react';
import { useParams, Link } from 'react-router-dom';
import { get, post } from '../api/client';
import { useVocabulary } from '../context/VocabularyContext';
import SpanishText from '../components/SpanishText';

function annotateTextNode(textNode, vocabMap) {
  const text = textNode.textContent;
  const tokens = text.match(/[\p{L}\p{M}]+|[^\p{L}\p{M}]+/gu);
  if (!tokens) return;

  let hasMatch = false;
  for (const token of tokens) {
    const cleaned = token.replace(/^[^\p{L}\p{M}]+|[^\p{L}\p{M}]+$/gu, '');
    if (cleaned && vocabMap.get(cleaned.toLowerCase())) {
      hasMatch = true;
      break;
    }
  }
  if (!hasMatch) return;

  const fragment = document.createDocumentFragment();
  for (const token of tokens) {
    const cleaned = token.replace(/^[^\p{L}\p{M}]+|[^\p{L}\p{M}]+$/gu, '');
    const translation = cleaned ? vocabMap.get(cleaned.toLowerCase()) : null;
    if (translation) {
      const span = document.createElement('span');
      span.className = 'spanish-word';
      span.setAttribute('data-tooltip', translation);
      span.textContent = token;
      fragment.appendChild(span);
    } else {
      fragment.appendChild(document.createTextNode(token));
    }
  }
  textNode.parentNode.replaceChild(fragment, textNode);
}

function annotateContainer(container, vocabMap) {
  const walker = document.createTreeWalker(container, NodeFilter.SHOW_TEXT, null);
  const textNodes = [];
  let node;
  while ((node = walker.nextNode())) {
    textNodes.push(node);
  }
  for (const tn of textNodes) {
    annotateTextNode(tn, vocabMap);
  }
}

export default function GrammarDetailPage() {
  const { slug } = useParams();
  const [lesson, setLesson] = useState(null);
  const contentRef = useRef(null);
  const annotatedRef = useRef(false);
  const vocab = useVocabulary();

  useEffect(() => {
    get(`/api/grammar/${slug}`).then(setLesson).catch(console.error);
    annotatedRef.current = false;
  }, [slug]);

  useEffect(() => {
    if (contentRef.current && vocab?.vocabMap?.size > 0 && lesson && !annotatedRef.current) {
      annotateContainer(contentRef.current, vocab.vocabMap);
      annotatedRef.current = true;
    }
  }, [lesson, vocab?.vocabMap]);

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
        ref={contentRef}
        dangerouslySetInnerHTML={{ __html: lesson.content }}
      />

      {lesson.examples && lesson.examples.length > 0 && (
        <div className="examples-section">
          <h2>Examples</h2>
          <div className="examples-list">
            {lesson.examples.map((ex) => (
              <div key={ex.id} className="example-card">
                <div className="example-spanish">
                  <SpanishText text={ex.spanish} />
                </div>
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
