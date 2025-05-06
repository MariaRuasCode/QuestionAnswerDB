-- ===============================
-- 📄 UPDATED SCHEMA WITH ANSWER STATUS
-- ===============================

-- 1. Create ENUM type first
CREATE TYPE answer_status AS ENUM (
    'normal',
    'spoiler',
    'no_comment',
    'hold_onto_that'
);

-- 2. Tag Types Table 
CREATE TABLE tag_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL 
);

-- 3. Tags Table 
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    tag_type_id INTEGER NOT NULL REFERENCES tag_types(id),
    value TEXT NOT NULL,
    UNIQUE(tag_type_id, value)
);

-- 4. Questions Table 
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    author TEXT,
    source TEXT,
    season INT,
    episode INT,
    timestamp TIMESTAMP
);

-- 5. Question-Tags Table 
CREATE TABLE question_tags (
    question_id INTEGER NOT NULL REFERENCES questions(id),
    tag_id INTEGER NOT NULL REFERENCES tags(id),
    PRIMARY KEY (question_id, tag_id)
);

-- 6. Updated Answers Table with Status
CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL REFERENCES questions(id),
    content TEXT NOT NULL,
    status answer_status NOT NULL DEFAULT 'normal',
    timestamp TIMESTAMP,
    is_followup BOOLEAN DEFAULT FALSE,
    "order" INT DEFAULT 1
);

-- 7. Add index for faster status filtering
CREATE INDEX idx_answers_status ON answers(status);