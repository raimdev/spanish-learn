import { createContext, useContext, useState, useEffect } from 'react';
import { get } from '../api/client';

const VocabularyContext = createContext(null);

const ARTICLES = ['el', 'la', 'los', 'las', 'un', 'una', 'unos', 'unas'];

function buildVocabMap(words) {
  const map = new Map();
  for (const word of words) {
    const spanish = word.spanish.toLowerCase();
    map.set(spanish, word.english);
    // Also index bare word without leading article
    const parts = spanish.split(' ');
    if (parts.length >= 2 && ARTICLES.includes(parts[0])) {
      map.set(parts.slice(1).join(' '), word.english);
    }
  }
  return map;
}

export function VocabularyProvider({ children }) {
  const [vocabMap, setVocabMap] = useState(new Map());

  useEffect(() => {
    get('/api/vocabulary')
      .then((words) => setVocabMap(buildVocabMap(words)))
      .catch(() => {});
  }, []);

  function lookup(word) {
    return vocabMap.get(word.toLowerCase()) || null;
  }

  return (
    <VocabularyContext.Provider value={{ vocabMap, lookup }}>
      {children}
    </VocabularyContext.Provider>
  );
}

export function useVocabulary() {
  return useContext(VocabularyContext);
}
