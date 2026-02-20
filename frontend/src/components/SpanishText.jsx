import { useVocabulary } from '../context/VocabularyContext';

// Split text into tokens of words and non-word characters (punctuation, spaces)
function tokenize(text) {
  return text.match(/[\p{L}\p{M}]+|[^\p{L}\p{M}]+/gu) || [];
}

function stripPunctuation(word) {
  return word.replace(/^[^\p{L}\p{M}]+|[^\p{L}\p{M}]+$/gu, '');
}

export default function SpanishText({ text }) {
  const vocab = useVocabulary();
  if (!vocab || !text) return text || null;

  const { lookup } = vocab;
  const tokens = tokenize(text);

  return (
    <>
      {tokens.map((token, i) => {
        const cleaned = stripPunctuation(token);
        if (!cleaned) return token;
        const translation = lookup(cleaned);
        if (translation) {
          return (
            <span key={i} className="spanish-word" data-tooltip={translation}>
              {token}
            </span>
          );
        }
        return token;
      })}
    </>
  );
}
