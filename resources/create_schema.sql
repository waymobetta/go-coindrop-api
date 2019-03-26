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
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

-- WALLETS
CREATE TABLE coindrop_wallets (
  id uuid DEFAULT uuid_generate_v4() UNIQUE,
  address text,
  user_id uuid REFERENCES coindrop_auth(id),
  type text,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone
);

-- REDDIT
CREATE TABLE coindrop_reddit (
  id uuid DEFAULT uuid_generate_v4() UNIQUE,
  user_id uuid REFERENCES coindrop_auth(id) UNIQUE,
  username text NOT NULL UNIQUE,
  comment_karma integer NOT NULL,
  link_karma integer NOT NULL,
  subreddits text[] NOT NULL,
  trophies text[] NOT NULL,
  posted_verification_code text NOT NULL,
  confirmed_verification_code text DEFAULT gen_verif_code() UNIQUE,
  verified boolean NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone
);

-- STACK OVERFLOW
CREATE TABLE IF NOT EXISTS coindrop_stackoverflow (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth(id) UNIQUE,
	exchange_account_id INTEGER NOT NULL,
	stack_user_id INTEGER NOT NULL UNIQUE,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	confirmed_verification_code TEXT DEFAULT gen_verif_code() UNIQUE,
	verified BOOLEAN NOT NULL,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

-- BADGE
CREATE TABLE IF NOT EXISTS coindrop_badges (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL UNIQUE,
	description TEXT,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone,
	logo_url TEXT
);

-- TASKS
CREATE TABLE IF NOT EXISTS coindrop_tasks (
  id uuid DEFAULT uuid_generate_v4() UNIQUE,
  title text NOT NULL,
  type text NOT NULL,
  author text NOT NULL,
  description text NOT NULL,
  token_name text,
  token_allocation integer,
  badge_id uuid REFERENCES coindrop_badges(id),
  logo_url text,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone
);

-- QUIZZES
CREATE TABLE IF NOT EXISTS coindrop_quizzes (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	title TEXT NOT NULL,
	typeform_form_url TEXT NOT NULL,
	typeform_form_id TEXT NOT NULL,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

-- QUIZ RESULTS
CREATE TABLE IF NOT EXISTS coindrop_quiz_results (
	quiz_id uuid REFERENCES coindrop_quizzes(id),
	typeform_form_id TEXT,
	user_id uuid REFERENCES coindrop_auth(id),
	questions_correct INTEGER NOT NULL,
	questions_incorrect INTEGER NOT NULL,
	quiz_taken BOOLEAN,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

-- TASKS SPECIFIC TO USER
CREATE TABLE IF NOT EXISTS coindrop_user_tasks (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	user_id uuid REFERENCES coindrop_auth(id),
	task_id uuid REFERENCES coindrop_tasks(id),
	completed BOOLEAN DEFAULT false,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

CREATE UNIQUE INDEX coindrop_wallets_user_id_type_uniq_idx ON coindrop_wallets(user_id uuid_ops,type text_ops);

-- NOT USED YET --

-- REDDIT COMMUNITIES
CREATE TABLE IF NOT EXISTS coindrop_reddit_communities (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

-- STACK OVERFLOW COMMUNITIES
CREATE TABLE IF NOT EXISTS coindrop_stackoverflow_communities (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	name TEXT NOT NULL,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);

CREATE TABLE "public"."coindrop_profiles" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "user_id" uuid UNIQUE,
    "name" text,
    "username" text,
    "image_url" text,
    PRIMARY KEY ("id"),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone
);
