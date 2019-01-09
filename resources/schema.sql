# sql statements
# TODO:
# add Keybase data

# AUTH
CREATE TABLE coindrop_auth (
  ID SERIAL NOT NULL UNIQUE,
  user_id TEXT NOT NULL UNIQUE,
  PRIMARY KEY(ID)
)

# REDDIT
CREATE TABLE coindrop_reddit (
	ID INTEGER NOT NULL UNIQUE REFERENCES coindrop_auth(ID),
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
	ID SERIAL NOT NULL UNIQUE REFERENCES coindrop_auth(ID),
	exchange_account_id INTEGER NOT NULL UNIQUE,
	user_id INTEGER NOT NULL UNIQUE,
	display_name TEXT NOT NULL,
	accounts TEXT ARRAY NOT NULL,
	posted_verification_code TEXT NOT NULL,
	stored_verification_code TEXT NOT NULL,
	is_verified BOOLEAN NOT NULL
)
