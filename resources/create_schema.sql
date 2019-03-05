-- psql [db name] < resources/create_schema.sql

-- create uuid_generate_v4 function
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- generate verification code function
CREATE FUNCTION gen_verif_code()
RETURNS TEXT AS $$
BEGIN
	RETURN CONCAT('[coindrop.io]-', uuid_generate_v4());
END; $$
LANGUAGE plpgsql;

-- AUTH
CREATE TABLE IF NOT EXISTS coindrop_auth (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	cognito_auth_user_id TEXT NOT NULL UNIQUE,
);

-- WALLETS
CREATE TABLE IF NOT EXISTS coindrop_wallets (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	address TEXT,
	user_id uuid REFERENCES coindrop_auth2(id)
);

CREATE UNIQUE INDEX "coindrop_wallets_address_user_id_uniq_idx" ON "public"."coindrop_wallets"("address","user_id");

CREATE UNIQUE INDEX "coindrop_wallets_user_id_uniq_idx" ON "public"."coindrop_wallets"("user_id");

-- REDDIT
CREATE TABLE IF NOT EXISTS coindrop_reddit (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth2(id),
	username TEXT NOT NULL,
	comment_karma INTEGER NOT NULL,
	link_karma INTEGER NOT NULL,
	subreddits TEXT ARRAY NOT NULL,
	trophies TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code text DEFAULT gen_verif_code() UNIQUE,
	verified BOOLEAN NOT NULL
);

-- STACK OVERFLOW
CREATE TABLE IF NOT EXISTS coindrop_stackoverflow (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth2(id),
	exchange_account_id INTEGER NOT NULL,
	stack_user_id INTEGER NOT NULL,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code text DEFAULT gen_verif_code() UNIQUE,
	verified BOOLEAN NOT NULL
);

-- TASKS
CREATE TABLE IF NOT EXISTS coindrop_tasks (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	title TEXT NOT NULL,
	type TEXT NOT NULL,
	author TEXT NOT NULL,
	description TEXT NOT NULL,
	token_name TEXT,
	token_allocation INTEGER,
	badge_id uuid REFERENCES coindrop_badges(id)
);

-- QUIZ RESULTS
CREATE TABLE IF NOT EXISTS coindrop_quiz_results (
	quiz_id uuid REFERENCES coindrop_quizzes2(id),
	user_id uuid REFERENCES coindrop_auth2(id),
	questions_correct INTEGER NOT NULL,
	questions_incorrect INTEGER NOT NULL,
	quiz_taken BOOLEAN
);

-- TASKS SPECIFIC TO USER
CREATE TABLE IF NOT EXISTS coindrop_user_tasks (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth2(id),
	task_id uuid REFERENCES coindrop_tasks2(id),
	completed BOOLEAN DEFAULT false;
);

-- QUIZZES
CREATE TABLE IF NOT EXISTS coindrop_quizzes (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	title TEXT NOT NULL,
	quiz_url TEXT NOT NULL,
	quiz_id TEXT NOT NULL
);

-- BADGE
CREATE TABLE IF NOT EXISTS coindrop_badges (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL,
	description TEXT
);

-- NOT USED YET --

-- REDDIT COMMUNITIES
CREATE TABLE IF NOT EXISTS coindrop_reddit_communities (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL
);

-- STACK OVERFLOW COMMUNITIES
CREATE TABLE IF NOT EXISTS coindrop_stackoverflow_communities (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL
);
