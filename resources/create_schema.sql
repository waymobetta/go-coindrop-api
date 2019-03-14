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
	cognito_auth_user_id TEXT NOT NULL UNIQUE
);

-- WALLETS
CREATE TABLE coindrop_wallets (
    id uuid DEFAULT uuid_generate_v4() UNIQUE,
    address text,
    user_id uuid REFERENCES coindrop_auth(id),
    type text
);

-- REDDIT
CREATE TABLE IF NOT EXISTS coindrop_reddit (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth(id),
	username TEXT NOT NULL,
	comment_karma INTEGER NOT NULL,
	link_karma INTEGER NOT NULL,
	subreddits TEXT ARRAY NOT NULL,
	trophies TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code TEXT DEFAULT gen_verif_code() UNIQUE,
	verified BOOLEAN NOT NULL
);

-- STACK OVERFLOW
CREATE TABLE IF NOT EXISTS coindrop_stackoverflow (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth(id),
	exchange_account_id INTEGER NOT NULL,
	stack_user_id INTEGER NOT NULL,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code TEXT DEFAULT gen_verif_code() UNIQUE,
	verified BOOLEAN NOT NULL
);

-- BADGE
CREATE TABLE IF NOT EXISTS coindrop_badges (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL,
	description TEXT
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

-- QUIZZES
CREATE TABLE IF NOT EXISTS coindrop_quizzes (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	title TEXT NOT NULL,
	typeform_form_url TEXT NOT NULL,
	typeform_form_id TEXT NOT NULL
);

-- QUIZ RESULTS
CREATE TABLE IF NOT EXISTS coindrop_quiz_results (
	quiz_id uuid REFERENCES coindrop_quizzes(id),
	typeform_form_id TEXT,
	user_id uuid REFERENCES coindrop_auth(id),
	questions_correct INTEGER NOT NULL,
	questions_incorrect INTEGER NOT NULL,
	quiz_taken BOOLEAN
);

-- TASKS SPECIFIC TO USER
CREATE TABLE IF NOT EXISTS coindrop_user_tasks (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth(id),
	task_id uuid REFERENCES coindrop_tasks(id),
	completed BOOLEAN DEFAULT false
);

CREATE UNIQUE INDEX coindrop_wallets_user_id_type_uniq_idx ON coindrop_wallets(user_id uuid_ops,type text_ops);

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

CREATE TABLE "public"."coindrop_profiles" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "user_id" uuid UNIQUE,
    "name" text,
    "username" text,
    "image_url" text,
    PRIMARY KEY ("id")
);
