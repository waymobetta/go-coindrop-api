-- psql [db name] < resources/create_schema.sql

-- AUTH
CREATE TABLE IF NOT EXISTS coindrop_auth (
  ID SERIAL NOT NULL UNIQUE,
  auth_user_id TEXT NOT NULL UNIQUE,
  wallet_address TEXT UNIQUE,
  PRIMARY KEY(ID)
);

-- REDDIT
CREATE TABLE IF NOT EXISTS coindrop_reddit (
	ID SERIAL NOT NULL UNIQUE,
	auth_user_id TEXT NOT NULL UNIQUE,
	username TEXT NOT NULL,
	comment_karma INTEGER NOT NULL,
	link_karma INTEGER NOT NULL,
	subreddits TEXT ARRAY NOT NULL,
	trophies TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	stored_verification_code TEXT NOT NULL,
	is_verified BOOLEAN NOT NULL,
	PRIMARY KEY(ID)
);

-- STACK OVERFLOW
CREATE TABLE IF NOT EXISTS IF NOT EXISTS coindrop_stackoverflow (
	ID SERIAL NOT NULL UNIQUE,
	auth_user_id TEXT NOT NULL UNIQUE,
	exchange_account_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	stored_verification_code TEXT NOT NULL,
	is_verified BOOLEAN NOT NULL,
	PRIMARY KEY(ID)
);

-- TASKS
CREATE TABLE IF NOT EXISTS coindrop_tasks (
	id SERIAL NOT NULL UNIQUE,
	title TEXT NOT NULL,
	type TEXT NOT NULL,
	author TEXT NOT NULL,
	description TEXT NOT NULL,
	token_name TEXT,
	token_allocation INTEGER,
	badge TEXT,
	PRIMARY KEY(ID)
);

-- QUIZ RESULTS
CREATE TABLE IF NOT EXISTS coindrop_quiz_results (
	id SERIAL NOT NULL UNIQUE,
	title TEXT NOT NULL,
	auth_user_id TEXT NOT NULL,
	questions_correct INT NOT NULL,
	questions_incorrect INT NOT NULL,
	has_taken_quiz BOOLEAN
);

-- TASKS SPECIFIC TO USER
CREATE TABLE IF NOT EXISTS coindrop_user_tasks (
	id SERIAL NOT NULL UNIQUE,
	auth_user_id TEXT NOT NULL UNIQUE,
	assigned TEXT ARRAY NOT NULL,
	completed TEXT ARRAY NOT NULL
);

-- QUIZZES
CREATE TABLE IF NOT EXISTS coindrop_quizzes (
	id SERIAL NOT NULL UNIQUE,
	title TEXT NOT NULL,
	quiz_data TEXT NOT NULL
);
