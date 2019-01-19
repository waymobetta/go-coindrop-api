# sql statements
# TODO:
# add Keybase data

# AUTH
CREATE TABLE coindrop_auth (
  ID SERIAL NOT NULL UNIQUE,
  auth_user_id TEXT NOT NULL UNIQUE,
  wallet_address TEXT UNIQUE,
  PRIMARY KEY(ID)
)

# REDDIT
CREATE TABLE coindrop_reddit (
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
)

# STACK OVERFLOW
CREATE TABLE coindrop_stackoverflow (
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
)

# TASKS
CREATE TABLE coindrop_tasks (
	id SERIAL NOT NULL UNIQUE,
	title TEXT NOT NULL,
	type TEXT NOT NULL,
	author TEXT NOT NULL,
	description TEXT NOT NULL,
	token_name TEXT,
	token_allocation INTEGER,
	badge TEXT,
	PRIMARY KEY(ID)
)

# TASK RESULTS
CREATE TABLE coindrop_tasks_results (
	id SERIAL NOT NULL UNIQUE,	
	users_enlisted TEXT ARRAY NOT NULL,
	users_passed TEXT ARRAY NOT NULL,
	users_failed TEXT ARRAY NOT NULL
)
