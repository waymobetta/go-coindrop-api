# sql statements
# TODO:
# add Keybase data

# profile info
CREATE TABLE coindropdb (
	ID INTEGER NOT NULL PRIMARY KEY UNIQUE,
	reddit_username TEXT NOT NULL UNIQUE,
	wallet_address TEXT NOT NULL UNIQUE,
	comment_karma INTEGER NOT NULL,
	link_karma INTEGER NOT NULL,
	subreddits TEXT ARRAY NOT NULL,
	trophies TEXT ARRAY NOT NULL,
	posted_twofa_code TEXT NOT NULL,
	stored_twofa_code TEXT NOT NULL,
	is_validated BOOLEAN NOT NULL,
	FOREIGN KEY(ID) REFERENCES coindropdbusers(id)
)

# auth info
create table coindropdbusers (
  id SERIAL PRIMARY KEY UNIQUE
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
)
