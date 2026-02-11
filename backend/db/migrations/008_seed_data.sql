-- Test user (password: "test")
INSERT OR IGNORE INTO users (username, password) VALUES
    ('test', '$2a$10$0BnA9Dza7ZpWkW/I/UnsuOt0NAEQEUSXFnC.vi2RmckoNh882vU4e');

-- Grammar Lessons
INSERT OR IGNORE INTO grammar_lessons (id, title, slug, summary, content, sort_order) VALUES
(1, 'Present Tense - Regular Verbs', 'present-tense-regular',
 'Learn how to conjugate regular -ar, -er, and -ir verbs in the present tense.',
 '<h2>Present Tense Regular Verbs</h2>
<p>Spanish regular verbs fall into three groups based on their infinitive endings: <strong>-ar</strong>, <strong>-er</strong>, and <strong>-ir</strong>.</p>

<h3>-AR Verbs (e.g., hablar - to speak)</h3>
<table>
<tr><th>Subject</th><th>Ending</th><th>Example</th></tr>
<tr><td>yo</td><td>-o</td><td>hablo</td></tr>
<tr><td>tu</td><td>-as</td><td>hablas</td></tr>
<tr><td>el/ella/usted</td><td>-a</td><td>habla</td></tr>
<tr><td>nosotros</td><td>-amos</td><td>hablamos</td></tr>
<tr><td>vosotros</td><td>-ais</td><td>hablais</td></tr>
<tr><td>ellos/ustedes</td><td>-an</td><td>hablan</td></tr>
</table>

<h3>-ER Verbs (e.g., comer - to eat)</h3>
<table>
<tr><th>Subject</th><th>Ending</th><th>Example</th></tr>
<tr><td>yo</td><td>-o</td><td>como</td></tr>
<tr><td>tu</td><td>-es</td><td>comes</td></tr>
<tr><td>el/ella/usted</td><td>-e</td><td>come</td></tr>
<tr><td>nosotros</td><td>-emos</td><td>comemos</td></tr>
<tr><td>vosotros</td><td>-eis</td><td>comeis</td></tr>
<tr><td>ellos/ustedes</td><td>-en</td><td>comen</td></tr>
</table>

<h3>-IR Verbs (e.g., vivir - to live)</h3>
<table>
<tr><th>Subject</th><th>Ending</th><th>Example</th></tr>
<tr><td>yo</td><td>-o</td><td>vivo</td></tr>
<tr><td>tu</td><td>-es</td><td>vives</td></tr>
<tr><td>el/ella/usted</td><td>-e</td><td>vive</td></tr>
<tr><td>nosotros</td><td>-imos</td><td>vivimos</td></tr>
<tr><td>vosotros</td><td>-is</td><td>vivis</td></tr>
<tr><td>ellos/ustedes</td><td>-en</td><td>viven</td></tr>
</table>

<p><strong>Tip:</strong> Notice that -er and -ir verbs share the same endings except for nosotros and vosotros forms.</p>', 1),

(2, 'Ser vs Estar', 'ser-vs-estar',
 'Understanding the two Spanish verbs for "to be" and when to use each one.',
 '<h2>Ser vs Estar</h2>
<p>Both <strong>ser</strong> and <strong>estar</strong> mean "to be," but they are used in different contexts.</p>

<h3>Use SER for:</h3>
<ul>
<li><strong>Identity:</strong> Yo soy Maria. (I am Maria.)</li>
<li><strong>Origin:</strong> Ella es de Mexico. (She is from Mexico.)</li>
<li><strong>Occupation:</strong> El es profesor. (He is a teacher.)</li>
<li><strong>Characteristics:</strong> La casa es grande. (The house is big.)</li>
<li><strong>Time/Date:</strong> Son las tres. (It is three o''clock.)</li>
<li><strong>Material:</strong> La mesa es de madera. (The table is made of wood.)</li>
</ul>

<h3>Use ESTAR for:</h3>
<ul>
<li><strong>Location:</strong> El libro esta en la mesa. (The book is on the table.)</li>
<li><strong>Emotions:</strong> Estoy contento. (I am happy.)</li>
<li><strong>Temporary states:</strong> La sopa esta caliente. (The soup is hot.)</li>
<li><strong>Progressive tenses:</strong> Estoy comiendo. (I am eating.)</li>
<li><strong>Results of actions:</strong> La puerta esta abierta. (The door is open.)</li>
</ul>

<h3>Conjugation of SER (Present Tense)</h3>
<table>
<tr><td>yo</td><td>soy</td></tr>
<tr><td>tu</td><td>eres</td></tr>
<tr><td>el/ella</td><td>es</td></tr>
<tr><td>nosotros</td><td>somos</td></tr>
<tr><td>vosotros</td><td>sois</td></tr>
<tr><td>ellos</td><td>son</td></tr>
</table>

<h3>Conjugation of ESTAR (Present Tense)</h3>
<table>
<tr><td>yo</td><td>estoy</td></tr>
<tr><td>tu</td><td>estas</td></tr>
<tr><td>el/ella</td><td>esta</td></tr>
<tr><td>nosotros</td><td>estamos</td></tr>
<tr><td>vosotros</td><td>estais</td></tr>
<tr><td>ellos</td><td>estan</td></tr>
</table>

<p><strong>Memory trick (DOCTOR vs PLACE):</strong></p>
<p>SER = <strong>D</strong>escription, <strong>O</strong>ccupation, <strong>C</strong>haracteristic, <strong>T</strong>ime, <strong>O</strong>rigin, <strong>R</strong>elationship</p>
<p>ESTAR = <strong>P</strong>osition, <strong>L</strong>ocation, <strong>A</strong>ction, <strong>C</strong>ondition, <strong>E</strong>motion</p>', 2),

(3, 'Articles and Gender', 'articles-and-gender',
 'Learn about definite and indefinite articles and noun gender in Spanish.',
 '<h2>Articles and Gender</h2>
<p>In Spanish, all nouns have a gender: <strong>masculine</strong> or <strong>feminine</strong>. Articles must agree with the noun in gender and number.</p>

<h3>Definite Articles (the)</h3>
<table>
<tr><th></th><th>Singular</th><th>Plural</th></tr>
<tr><td>Masculine</td><td>el</td><td>los</td></tr>
<tr><td>Feminine</td><td>la</td><td>las</td></tr>
</table>

<h3>Indefinite Articles (a/an, some)</h3>
<table>
<tr><th></th><th>Singular</th><th>Plural</th></tr>
<tr><td>Masculine</td><td>un</td><td>unos</td></tr>
<tr><td>Feminine</td><td>una</td><td>unas</td></tr>
</table>

<h3>Gender Rules</h3>
<p><strong>Generally masculine:</strong> nouns ending in -o, -or, -aje</p>
<ul><li>el libro (the book), el doctor (the doctor), el viaje (the trip)</li></ul>

<p><strong>Generally feminine:</strong> nouns ending in -a, -cion, -sion, -dad, -tad</p>
<ul><li>la casa (the house), la nacion (the nation), la ciudad (the city)</li></ul>

<p><strong>Common exceptions:</strong></p>
<ul>
<li>el dia (the day) - masculine despite ending in -a</li>
<li>la mano (the hand) - feminine despite ending in -o</li>
<li>el problema (the problem) - masculine (Greek origin)</li>
<li>el mapa (the map) - masculine (Greek origin)</li>
</ul>', 3),

(4, 'Common Irregular Verbs', 'common-irregular-verbs',
 'Master the most frequently used irregular verbs in Spanish.',
 '<h2>Common Irregular Verbs</h2>
<p>Some of the most important Spanish verbs are irregular. Here are the essential ones to memorize.</p>

<h3>Tener (to have)</h3>
<table>
<tr><td>yo tengo</td><td>nosotros tenemos</td></tr>
<tr><td>tu tienes</td><td>vosotros teneis</td></tr>
<tr><td>el/ella tiene</td><td>ellos tienen</td></tr>
</table>
<p><strong>Special uses:</strong> tener hambre (to be hungry), tener sed (to be thirsty), tener frio (to be cold), tener calor (to be hot), tener sueno (to be sleepy), tener anos (to be ... years old)</p>

<h3>Ir (to go)</h3>
<table>
<tr><td>yo voy</td><td>nosotros vamos</td></tr>
<tr><td>tu vas</td><td>vosotros vais</td></tr>
<tr><td>el/ella va</td><td>ellos van</td></tr>
</table>
<p><strong>Ir + a + infinitive</strong> expresses future: Voy a comer. (I am going to eat.)</p>

<h3>Hacer (to do/make)</h3>
<table>
<tr><td>yo hago</td><td>nosotros hacemos</td></tr>
<tr><td>tu haces</td><td>vosotros haceis</td></tr>
<tr><td>el/ella hace</td><td>ellos hacen</td></tr>
</table>

<h3>Poder (to be able to / can)</h3>
<table>
<tr><td>yo puedo</td><td>nosotros podemos</td></tr>
<tr><td>tu puedes</td><td>vosotros podeis</td></tr>
<tr><td>el/ella puede</td><td>ellos pueden</td></tr>
</table>

<h3>Querer (to want)</h3>
<table>
<tr><td>yo quiero</td><td>nosotros queremos</td></tr>
<tr><td>tu quieres</td><td>vosotros quereis</td></tr>
<tr><td>el/ella quiere</td><td>ellos quieren</td></tr>
</table>', 4),

(5, 'Basic Sentence Structure', 'basic-sentence-structure',
 'Learn the fundamental word order and sentence patterns in Spanish.',
 '<h2>Basic Sentence Structure</h2>
<p>Spanish follows a <strong>Subject-Verb-Object (SVO)</strong> order, similar to English, but with more flexibility.</p>

<h3>Basic Pattern: Subject + Verb + Object</h3>
<ul>
<li>Maria come una manzana. (Maria eats an apple.)</li>
<li>Yo leo un libro. (I read a book.)</li>
</ul>

<h3>Subject Pronouns Are Often Omitted</h3>
<p>Because verb conjugations indicate the subject, pronouns are frequently dropped:</p>
<ul>
<li>Como una manzana. (I eat an apple.) - "Yo" is implied</li>
<li>Hablas espanol. (You speak Spanish.) - "Tu" is implied</li>
</ul>

<h3>Adjectives Usually Follow Nouns</h3>
<ul>
<li>un coche rojo (a red car) - NOT "un rojo coche"</li>
<li>una casa grande (a big house)</li>
<li>el libro interesante (the interesting book)</li>
</ul>

<h3>Negation</h3>
<p>Place <strong>no</strong> before the verb:</p>
<ul>
<li>No hablo frances. (I don''t speak French.)</li>
<li>No tengo dinero. (I don''t have money.)</li>
</ul>

<h3>Questions</h3>
<p>Questions can be formed by:</p>
<ul>
<li><strong>Inversion:</strong> Hablas espanol? becomes Hablas tu espanol?</li>
<li><strong>Intonation:</strong> Simply raise your voice at the end: Hablas espanol?</li>
<li><strong>Question words:</strong> Que (what), Donde (where), Cuando (when), Como (how), Por que (why), Quien (who), Cuanto (how much)</li>
</ul>', 5);

-- Grammar Examples
INSERT OR IGNORE INTO grammar_examples (id, lesson_id, spanish, english, explanation, sort_order) VALUES
(1, 1, 'Yo hablo espanol.', 'I speak Spanish.', '-ar verb: hablar -> hablo', 1),
(2, 1, 'Tu comes mucho.', 'You eat a lot.', '-er verb: comer -> comes', 2),
(3, 1, 'Ella vive en Madrid.', 'She lives in Madrid.', '-ir verb: vivir -> vive', 3),
(4, 1, 'Nosotros estudiamos juntos.', 'We study together.', '-ar verb: estudiar -> estudiamos', 4),
(5, 2, 'Yo soy estudiante.', 'I am a student.', 'Ser for identity/occupation', 1),
(6, 2, 'Ella esta cansada.', 'She is tired.', 'Estar for temporary condition', 2),
(7, 2, 'La fiesta es en mi casa.', 'The party is at my house.', 'Ser for events', 3),
(8, 2, 'El cafe esta caliente.', 'The coffee is hot.', 'Estar for temporary state', 4),
(9, 3, 'El gato es negro.', 'The cat is black.', 'Masculine singular: el', 1),
(10, 3, 'La mesa es grande.', 'The table is big.', 'Feminine singular: la', 2),
(11, 3, 'Los libros son interesantes.', 'The books are interesting.', 'Masculine plural: los', 3),
(12, 4, 'Tengo veinte anos.', 'I am twenty years old.', 'Tener for age', 1),
(13, 4, 'Voy al supermercado.', 'I go to the supermarket.', 'Ir: voy (a + el = al)', 2),
(14, 5, 'El nino come una manzana roja.', 'The boy eats a red apple.', 'SVO with adjective after noun', 1),
(15, 5, 'No quiero cafe.', 'I don''t want coffee.', 'Negation with no before verb', 2);

-- Vocabulary Words
INSERT OR IGNORE INTO vocabulary_words (id, spanish, english, part_of_speech, category, example_sentence) VALUES
-- Greetings
(1, 'hola', 'hello', 'interjection', 'greetings', 'Hola, como estas?'),
(2, 'buenos dias', 'good morning', 'phrase', 'greetings', 'Buenos dias, senor.'),
(3, 'buenas tardes', 'good afternoon', 'phrase', 'greetings', 'Buenas tardes a todos.'),
(4, 'buenas noches', 'good night', 'phrase', 'greetings', 'Buenas noches, hasta manana.'),
(5, 'adios', 'goodbye', 'interjection', 'greetings', 'Adios, nos vemos.'),
(6, 'hasta luego', 'see you later', 'phrase', 'greetings', 'Hasta luego, amigo.'),
(7, 'por favor', 'please', 'phrase', 'greetings', 'Un cafe, por favor.'),
(8, 'gracias', 'thank you', 'interjection', 'greetings', 'Muchas gracias por tu ayuda.'),
(9, 'de nada', 'you''re welcome', 'phrase', 'greetings', 'De nada, fue un placer.'),
(10, 'lo siento', 'I''m sorry', 'phrase', 'greetings', 'Lo siento, no entiendo.'),
-- Food
(11, 'la manzana', 'the apple', 'noun', 'food', 'La manzana es roja.'),
(12, 'el pan', 'the bread', 'noun', 'food', 'Quiero pan con mantequilla.'),
(13, 'la leche', 'the milk', 'noun', 'food', 'Bebo leche todos los dias.'),
(14, 'el agua', 'the water', 'noun', 'food', 'Necesito un vaso de agua.'),
(15, 'la carne', 'the meat', 'noun', 'food', 'No como carne.'),
(16, 'el pescado', 'the fish', 'noun', 'food', 'El pescado esta fresco.'),
(17, 'el arroz', 'the rice', 'noun', 'food', 'El arroz con pollo es delicioso.'),
(18, 'la fruta', 'the fruit', 'noun', 'food', 'Me gusta la fruta tropical.'),
(19, 'la verdura', 'the vegetable', 'noun', 'food', 'Las verduras son saludables.'),
(20, 'el cafe', 'the coffee', 'noun', 'food', 'Quiero un cafe con leche.'),
-- Family
(21, 'la madre', 'the mother', 'noun', 'family', 'Mi madre cocina muy bien.'),
(22, 'el padre', 'the father', 'noun', 'family', 'Mi padre trabaja mucho.'),
(23, 'el hermano', 'the brother', 'noun', 'family', 'Tengo un hermano mayor.'),
(24, 'la hermana', 'the sister', 'noun', 'family', 'Mi hermana vive en Barcelona.'),
(25, 'el hijo', 'the son', 'noun', 'family', 'Su hijo tiene cinco anos.'),
(26, 'la hija', 'the daughter', 'noun', 'family', 'La hija estudia medicina.'),
(27, 'el abuelo', 'the grandfather', 'noun', 'family', 'Mi abuelo cuenta historias.'),
(28, 'la abuela', 'the grandmother', 'noun', 'family', 'La abuela hace galletas.'),
(29, 'el tio', 'the uncle', 'noun', 'family', 'Mi tio vive en Mexico.'),
(30, 'la tia', 'the aunt', 'noun', 'family', 'Mi tia es profesora.'),
-- Common verbs
(31, 'hablar', 'to speak', 'verb', 'verbs', 'Yo hablo espanol.'),
(32, 'comer', 'to eat', 'verb', 'verbs', 'Nosotros comemos juntos.'),
(33, 'vivir', 'to live', 'verb', 'verbs', 'Ella vive en Madrid.'),
(34, 'trabajar', 'to work', 'verb', 'verbs', 'Trabajo en una oficina.'),
(35, 'estudiar', 'to study', 'verb', 'verbs', 'Estudio espanol todos los dias.'),
(36, 'leer', 'to read', 'verb', 'verbs', 'Me gusta leer libros.'),
(37, 'escribir', 'to write', 'verb', 'verbs', 'Escribo una carta.'),
(38, 'dormir', 'to sleep', 'verb', 'verbs', 'Duermo ocho horas.'),
(39, 'comprar', 'to buy', 'verb', 'verbs', 'Voy a comprar comida.'),
(40, 'caminar', 'to walk', 'verb', 'verbs', 'Camino al parque.'),
-- Colors
(41, 'rojo', 'red', 'adjective', 'colors', 'El coche es rojo.'),
(42, 'azul', 'blue', 'adjective', 'colors', 'El cielo es azul.'),
(43, 'verde', 'green', 'adjective', 'colors', 'La hierba es verde.'),
(44, 'amarillo', 'yellow', 'adjective', 'colors', 'El sol es amarillo.'),
(45, 'blanco', 'white', 'adjective', 'colors', 'La nieve es blanca.'),
(46, 'negro', 'black', 'adjective', 'colors', 'El gato es negro.'),
-- Numbers
(47, 'uno', 'one', 'numeral', 'numbers', 'Tengo un hermano.'),
(48, 'dos', 'two', 'numeral', 'numbers', 'Dos mas dos son cuatro.'),
(49, 'tres', 'three', 'numeral', 'numbers', 'Hay tres libros.'),
(50, 'diez', 'ten', 'numeral', 'numbers', 'Tengo diez dedos.');

-- Exercises
INSERT OR IGNORE INTO exercises (id, type, difficulty, prompt, correct_answer, options, hint, lesson_id) VALUES
-- Fill in the blank
(1, 'fill_blank', 'beginner', 'Yo ___ (hablar) espanol.', 'hablo', NULL, 'First person singular of -ar verb', 1),
(2, 'fill_blank', 'beginner', 'Tu ___ (comer) pizza.', 'comes', NULL, 'Second person singular of -er verb', 1),
(3, 'fill_blank', 'beginner', 'Nosotros ___ (vivir) en Madrid.', 'vivimos', NULL, 'First person plural of -ir verb', 1),
(4, 'fill_blank', 'beginner', 'Ella ___ (ser) doctora.', 'es', NULL, 'Third person singular of ser', 2),
(5, 'fill_blank', 'beginner', 'Yo ___ (estar) cansado.', 'estoy', NULL, 'First person singular of estar', 2),
(6, 'fill_blank', 'intermediate', 'Ellos ___ (tener) tres gatos.', 'tienen', NULL, 'Third person plural of tener', 4),
(7, 'fill_blank', 'intermediate', 'Nosotros ___ (ir) al cine.', 'vamos', NULL, 'First person plural of ir', 4),
(8, 'fill_blank', 'intermediate', 'Yo ___ (hacer) la tarea.', 'hago', NULL, 'First person singular of hacer', 4),
-- Translation
(9, 'translation', 'beginner', 'Translate: "The cat eats fish."', 'El gato come pescado.', NULL, 'Use present tense of comer', NULL),
(10, 'translation', 'beginner', 'Translate: "I speak Spanish."', 'Yo hablo espanol.', NULL, 'Use hablar', NULL),
(11, 'translation', 'beginner', 'Translate: "The house is big."', 'La casa es grande.', NULL, 'Use ser for characteristics', 2),
(12, 'translation', 'intermediate', 'Translate: "We are going to the park."', 'Vamos al parque.', NULL, 'Use ir, a + el = al', 4),
(13, 'translation', 'intermediate', 'Translate: "She is tired."', 'Ella esta cansada.', NULL, 'Use estar for temporary states', 2),
(14, 'translation', 'intermediate', 'Translate: "I have two brothers."', 'Tengo dos hermanos.', NULL, 'Use tener', 4),
-- Conjugation
(15, 'conjugation', 'beginner', 'Conjugate "hablar" for "yo" in present tense.', 'hablo', NULL, 'Remove -ar, add -o', 1),
(16, 'conjugation', 'beginner', 'Conjugate "comer" for "nosotros" in present tense.', 'comemos', NULL, 'Remove -er, add -emos', 1),
(17, 'conjugation', 'beginner', 'Conjugate "vivir" for "ellos" in present tense.', 'viven', NULL, 'Remove -ir, add -en', 1),
(18, 'conjugation', 'intermediate', 'Conjugate "tener" for "yo" in present tense.', 'tengo', NULL, 'Irregular: stem changes to teng-', 4),
(19, 'conjugation', 'intermediate', 'Conjugate "ir" for "nosotros" in present tense.', 'vamos', NULL, 'Highly irregular verb', 4),
(20, 'conjugation', 'intermediate', 'Conjugate "poder" for "tu" in present tense.', 'puedes', NULL, 'Stem change: o -> ue', 4);

-- Flashcard Decks
INSERT OR IGNORE INTO flashcard_decks (id, name, description) VALUES
(1, 'Basic Greetings', 'Common Spanish greetings and polite expressions'),
(2, 'Food and Drink', 'Vocabulary related to food, drinks, and dining'),
(3, 'Essential Verbs', 'The most important Spanish verbs to know');

-- Flashcards - Basic Greetings
INSERT OR IGNORE INTO flashcards (id, deck_id, front, back) VALUES
(1, 1, 'Hola', 'Hello'),
(2, 1, 'Buenos dias', 'Good morning'),
(3, 1, 'Buenas tardes', 'Good afternoon'),
(4, 1, 'Buenas noches', 'Good night'),
(5, 1, 'Adios', 'Goodbye'),
(6, 1, 'Hasta luego', 'See you later'),
(7, 1, 'Por favor', 'Please'),
(8, 1, 'Gracias', 'Thank you'),
(9, 1, 'De nada', 'You''re welcome'),
(10, 1, 'Lo siento', 'I''m sorry'),
(11, 1, 'Como estas?', 'How are you?'),
(12, 1, 'Bien, gracias', 'Fine, thank you');

-- Flashcards - Food and Drink
INSERT OR IGNORE INTO flashcards (id, deck_id, front, back) VALUES
(13, 2, 'La manzana', 'The apple'),
(14, 2, 'El pan', 'The bread'),
(15, 2, 'La leche', 'The milk'),
(16, 2, 'El agua', 'The water'),
(17, 2, 'La carne', 'The meat'),
(18, 2, 'El pescado', 'The fish'),
(19, 2, 'El arroz', 'The rice'),
(20, 2, 'La fruta', 'The fruit'),
(21, 2, 'El cafe', 'The coffee'),
(22, 2, 'La cerveza', 'The beer'),
(23, 2, 'El pollo', 'The chicken'),
(24, 2, 'La ensalada', 'The salad');

-- Flashcards - Essential Verbs
INSERT OR IGNORE INTO flashcards (id, deck_id, front, back) VALUES
(25, 3, 'Hablar', 'To speak'),
(26, 3, 'Comer', 'To eat'),
(27, 3, 'Vivir', 'To live'),
(28, 3, 'Ser', 'To be (permanent)'),
(29, 3, 'Estar', 'To be (temporary)'),
(30, 3, 'Tener', 'To have'),
(31, 3, 'Ir', 'To go'),
(32, 3, 'Hacer', 'To do / to make'),
(33, 3, 'Poder', 'To be able to / can'),
(34, 3, 'Querer', 'To want'),
(35, 3, 'Saber', 'To know (facts)'),
(36, 3, 'Conocer', 'To know (people/places)');
