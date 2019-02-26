-- psql [db name] < resources/create_schema2.sql

-- AUTH 2
CREATE TABLE IF NOT EXISTS coindrop_auth2 (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	cognito_auth_user_id TEXT NOT NULL UNIQUE,
	wallet_id uuid REFERENCES coindrop_wallets(id)
);

-- WALLETS 2
CREATE TABLE IF NOT EXISTS coindrop_wallets (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	address TEXT
);

-- REDDIT 2
CREATE TABLE IF NOT EXISTS coindrop_reddit2 (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth2(id),
	username TEXT NOT NULL,
	comment_karma INTEGER NOT NULL,
	link_karma INTEGER NOT NULL,
	subreddits TEXT ARRAY NOT NULL,
	trophies TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code TEXT NOT NULL,
	verified BOOLEAN NOT NULL
);

-- STACK OVERFLOW 2
CREATE TABLE IF NOT EXISTS coindrop_stackoverflow2 (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth2(id),
	exchange_account_id INTEGER NOT NULL,
	stack_user_id INTEGER NOT NULL,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code TEXT NOT NULL,
	verified BOOLEAN NOT NULL
);

-- TASKS 2
CREATE TABLE IF NOT EXISTS coindrop_tasks2 (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	title TEXT NOT NULL,
	type TEXT NOT NULL,
	author TEXT NOT NULL,
	description TEXT NOT NULL,
	token_name TEXT,
	token_allocation INTEGER,
	badge TEXT
);

-- QUIZ RESULTS 2
CREATE TABLE IF NOT EXISTS coindrop_quiz_results2 (
	quiz_id uuid REFERENCES coindroo_quizzes2(id),
	user_id uuid REFERENCES coindrop_auth2(id),
	questions_correct INT NOT NULL,
	questions_incorrect INT NOT NULL,
	quiz_taken BOOLEAN
);

-- TASKS SPECIFIC TO USER 2
CREATE TABLE IF NOT EXISTS coindrop_user_tasks2 (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth2(id),
	task_id uuid REFERENCES coindrop_tasks2(id),
	completed BOOLEAN
);

-- QUIZZES 2
CREATE TABLE IF NOT EXISTS coindrop_quizzes2 (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	title TEXT NOT NULL,
	quiz_url TEXT NOT NULL
);
