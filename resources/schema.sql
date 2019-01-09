# sql statements
# TODO:
# add Keybase data

# AUTH
CREATE TABLE coindrop_auth (
  ID SERIAL NOT NULL UNIQUE,
  auth_user_id TEXT NOT NULL UNIQUE,
  PRIMARY KEY(ID)
)

# REDDIT
CREATE TABLE coindrop_reddit (
	ID SERIAL NOT NULL UNIQUE,
	auth_user_id TEXT NOT NULL UNIQUE,
	username TEXT NOT NULL UNIQUE,
	wallet_address TEXT NOT NULL UNIQUE,
	comment_karma INTEGER NOT NULL,
	link_karma INTEGER NOT NULL,
	subreddits TEXT ARRAY NOT NULL,
	trophies TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	stored_verification_code TEXT NOT NULL,
	is_verified BOOLEAN NOT NULL
)

# STACK OVERFLOW
CREATE TABLE coindrop_stackoverflow (
	ID SERIAL NOT NULL UNIQUE,
	auth_user_id TEXT NOT NULL UNIQUE,
	exchange_account_id INTEGER UNIQUE NOT NULL,
	user_id INTEGER UNIQUE NOT NULL,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	stored_verification_code TEXT NOT NULL,
	is_verified BOOLEAN NOT NULL
)

# TASKS
CREATE TABLE coindrop_tasks (
	task_id SERIAL NOT NULL UNIQUE,
	task_name TEXT NOT NULL,
	task_author TEXT NOT NULL UNIQUE,
	task_description TEXT NOT NULL,
	task_token_name TEXT NOT NULL,
	task_token_allocation INTEGER NOT NULL,
	task_badge TEXT NOT NULL UNIQUE
)
