import { useState, useEffect } from 'react';
import { get } from '../api/client';

export default function VocabularyPage() {
  const [categories, setCategories] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState('');
  const [words, setWords] = useState([]);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [userAnswer, setUserAnswer] = useState('');
  const [feedback, setFeedback] = useState(null);
  const [score, setScore] = useState({ correct: 0, total: 0 });
  const [quizActive, setQuizActive] = useState(false);
  const [quizComplete, setQuizComplete] = useState(false);
  const [direction, setDirection] = useState('es-en');

  useEffect(() => {
    get('/api/vocabulary/categories').then(setCategories).catch(console.error);
  }, []);

  async function startQuiz() {
    const params = new URLSearchParams({ count: '10' });
    if (selectedCategory) params.set('category', selectedCategory);
    const data = await get(`/api/vocabulary/quiz?${params}`);
    setWords(data);
    setCurrentIndex(0);
    setScore({ correct: 0, total: 0 });
    setFeedback(null);
    setUserAnswer('');
    setQuizActive(true);
    setQuizComplete(false);
  }

  function checkAnswer() {
    const word = words[currentIndex];
    const expected = direction === 'es-en' ? word.english : word.spanish;
    const isCorrect = userAnswer.trim().toLowerCase() === expected.toLowerCase();

    setFeedback({
      correct: isCorrect,
      expected,
    });
    setScore((s) => ({
      correct: s.correct + (isCorrect ? 1 : 0),
      total: s.total + 1,
    }));
  }

  function nextWord() {
    if (currentIndex + 1 >= words.length) {
      setQuizComplete(true);
    } else {
      setCurrentIndex((i) => i + 1);
      setUserAnswer('');
      setFeedback(null);
    }
  }

  function handleKeyDown(e) {
    if (e.key === 'Enter') {
      if (feedback) {
        nextWord();
      } else {
        checkAnswer();
      }
    }
  }

  if (!quizActive) {
    return (
      <div className="vocabulary">
        <h1>Vocabulary Training</h1>
        <p className="subtitle">Test your knowledge of Spanish words</p>

        <div className="quiz-setup">
          <div className="form-group">
            <label>Category</label>
            <select value={selectedCategory} onChange={(e) => setSelectedCategory(e.target.value)}>
              <option value="">All categories</option>
              {categories.map((c) => (
                <option key={c} value={c}>{c}</option>
              ))}
            </select>
          </div>

          <div className="form-group">
            <label>Direction</label>
            <select value={direction} onChange={(e) => setDirection(e.target.value)}>
              <option value="es-en">Spanish to English</option>
              <option value="en-es">English to Spanish</option>
            </select>
          </div>

          <button onClick={startQuiz} className="btn-primary">Start Quiz</button>
        </div>
      </div>
    );
  }

  if (quizComplete) {
    return (
      <div className="vocabulary">
        <h1>Quiz Complete!</h1>
        <div className="quiz-result">
          <div className="result-score">
            {score.correct} / {score.total}
          </div>
          <div className="result-percentage">
            {Math.round((score.correct / score.total) * 100)}% correct
          </div>
          <button onClick={startQuiz} className="btn-primary">Try Again</button>
          <button onClick={() => setQuizActive(false)} className="btn-secondary">Change Settings</button>
        </div>
      </div>
    );
  }

  const word = words[currentIndex];
  const prompt = direction === 'es-en' ? word.spanish : word.english;
  const promptLabel = direction === 'es-en' ? 'Spanish' : 'English';
  const answerLabel = direction === 'es-en' ? 'English' : 'Spanish';

  return (
    <div className="vocabulary">
      <div className="quiz-progress">
        Word {currentIndex + 1} of {words.length} | Score: {score.correct}/{score.total}
      </div>

      <div className="quiz-card">
        <div className="quiz-prompt-label">{promptLabel}</div>
        <div className="quiz-prompt">{prompt}</div>

        {word.part_of_speech && (
          <div className="quiz-hint">({word.part_of_speech})</div>
        )}

        <div className="form-group">
          <label>Your answer ({answerLabel})</label>
          <input
            type="text"
            value={userAnswer}
            onChange={(e) => setUserAnswer(e.target.value)}
            onKeyDown={handleKeyDown}
            disabled={!!feedback}
            autoFocus
            placeholder={`Type the ${answerLabel} translation...`}
          />
        </div>

        {!feedback ? (
          <button onClick={checkAnswer} className="btn-primary" disabled={!userAnswer.trim()}>
            Check Answer
          </button>
        ) : (
          <div>
            <div className={`feedback ${feedback.correct ? 'feedback-correct' : 'feedback-incorrect'}`}>
              {feedback.correct ? 'Correct!' : `Incorrect. Answer: ${feedback.expected}`}
            </div>
            <button onClick={nextWord} className="btn-primary" autoFocus>
              {currentIndex + 1 >= words.length ? 'See Results' : 'Next Word'}
            </button>
          </div>
        )}
      </div>
    </div>
  );
}
