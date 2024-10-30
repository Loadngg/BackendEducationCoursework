CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    login    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE lectures
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE,
    text  TEXT         NOT NULL
);

CREATE TABLE lecture_materials
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(255) NOT NULL UNIQUE,
    file       BYTEA        NOT NULL,
    lecture_id INTEGER REFERENCES lectures (id)
);

CREATE TABLE quiz
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE quiz_questions
(
    id                SERIAL PRIMARY KEY,
    question          TEXT NOT NULL,
    correct_answer_id INTEGER,
    quiz_id           INTEGER REFERENCES quiz (id)
);

CREATE TABLE answers
(
    id          SERIAL PRIMARY KEY,
    text        TEXT NOT NULL,
    question_id INTEGER REFERENCES quiz_questions (id)
);

CREATE TABLE user_answers
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER REFERENCES users (id),
    question_id INTEGER REFERENCES quiz_questions (id),
    answer_id   INTEGER REFERENCES answers (id),
    is_correct  BOOLEAN NOT NULL
);

CREATE TABLE user_quiz
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),
    quiz_id INTEGER REFERENCES quiz (id),
    score   INTEGER NOT NULL
);
