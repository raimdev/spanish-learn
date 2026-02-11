import { useState, useEffect } from 'react';
import { get, post } from '../api/client';

const QUALITY_BUTTONS = [
  { value: 0, label: 'Blackout', color: '#e74c3c' },
  { value: 1, label: 'Wrong', color: '#e67e22' },
  { value: 2, label: 'Hard', color: '#f39c12' },
  { value: 3, label: 'OK', color: '#27ae60' },
  { value: 4, label: 'Good', color: '#2ecc71' },
  { value: 5, label: 'Easy', color: '#00d4aa' },
];

export default function FlashcardsPage() {
  const [decks, setDecks] = useState([]);
  const [selectedDeck, setSelectedDeck] = useState('');
  const [cards, setCards] = useState([]);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [flipped, setFlipped] = useState(false);
  const [sessionActive, setSessionActive] = useState(false);
  const [sessionComplete, setSessionComplete] = useState(false);
  const [reviewed, setReviewed] = useState(0);

  useEffect(() => {
    get('/api/flashcards/decks').then(setDecks).catch(console.error);
  }, []);

  async function startSession() {
    const params = new URLSearchParams({ limit: '20' });
    if (selectedDeck) params.set('deck_id', selectedDeck);
    const data = await get(`/api/flashcards/due?${params}`);
    if (data.length === 0) {
      alert('No cards due for review! Check back later.');
      return;
    }
    setCards(data);
    setCurrentIndex(0);
    setFlipped(false);
    setReviewed(0);
    setSessionActive(true);
    setSessionComplete(false);
  }

  async function rateCard(quality) {
    const card = cards[currentIndex];
    await post(`/api/flashcards/${card.id}/review`, { quality });
    setReviewed((r) => r + 1);

    if (currentIndex + 1 >= cards.length) {
      setSessionComplete(true);
    } else {
      setCurrentIndex((i) => i + 1);
      setFlipped(false);
    }
  }

  if (!sessionActive) {
    return (
      <div className="flashcards">
        <h1>Flashcard Review</h1>
        <p className="subtitle">Spaced repetition for long-term memory</p>

        <div className="deck-list">
          {decks.map((deck) => (
            <div
              key={deck.id}
              className={`deck-card ${selectedDeck === String(deck.id) ? 'selected' : ''}`}
              onClick={() => setSelectedDeck(selectedDeck === String(deck.id) ? '' : String(deck.id))}
            >
              <h3>{deck.name}</h3>
              {deck.description && <p>{deck.description}</p>}
              <span className="deck-count">{deck.card_count} cards</span>
            </div>
          ))}
        </div>

        <button onClick={startSession} className="btn-primary">
          {selectedDeck ? 'Review Selected Deck' : 'Review All Due Cards'}
        </button>
      </div>
    );
  }

  if (sessionComplete) {
    return (
      <div className="flashcards">
        <h1>Session Complete!</h1>
        <div className="session-result">
          <div className="result-score">Reviewed {reviewed} cards</div>
          <button onClick={startSession} className="btn-primary">Start Another Session</button>
          <button onClick={() => setSessionActive(false)} className="btn-secondary">Back to Decks</button>
        </div>
      </div>
    );
  }

  const card = cards[currentIndex];

  return (
    <div className="flashcards">
      <div className="session-progress">
        Card {currentIndex + 1} of {cards.length} | Reviewed: {reviewed}
      </div>

      <div className="flashcard-container" onClick={() => !flipped && setFlipped(true)}>
        <div className={`flashcard ${flipped ? 'flipped' : ''}`}>
          <div className="flashcard-front">
            <div className="flashcard-label">Front</div>
            <div className="flashcard-text">{card.front}</div>
            <div className="flashcard-instruction">Click to reveal answer</div>
          </div>
          <div className="flashcard-back">
            <div className="flashcard-label">Back</div>
            <div className="flashcard-text">{card.back}</div>
          </div>
        </div>
      </div>

      {flipped && (
        <div className="rating-section">
          <p className="rating-prompt">How well did you remember?</p>
          <div className="rating-buttons">
            {QUALITY_BUTTONS.map((btn) => (
              <button
                key={btn.value}
                onClick={() => rateCard(btn.value)}
                className="rating-btn"
                style={{ borderColor: btn.color, color: btn.color }}
              >
                <span className="rating-value">{btn.value}</span>
                <span className="rating-label">{btn.label}</span>
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
