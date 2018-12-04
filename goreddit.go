package main

import (
	"errors"
	"os"

	"github.com/jzelinskie/geddit"
)

var (
	// TwoFASubredditName name of subreddit used for 2FA
	TwoFASubredditName = "testing_QA_adChain"
)

// InitRedditAuth initializes reddit OAuth session
func (a *AuthSessions) InitRedditAuth() (*AuthSessions, error) {
	// initialize OAuth session with credentials as environment variables
	oAuthSession, err := geddit.NewOAuthSession(
		os.Getenv("REDDIT_CLIENT_ID"),
		os.Getenv("REDDIT_CLIENT_SECRET"),
		os.Getenv("REDDIT_USER_AGENT"),
		"http://metax.io",
	)
	if err != nil {
		return a, err
	}

	// create new auth token for confidential clients (personal scripts/apps).
	err = oAuthSession.LoginAuth(os.Getenv("REDDIT_USERNAME"), os.Getenv("REDDIT_PASSWORD"))
	if err != nil {
		return a, err
	}
	// create new unauthenticated session
	session := geddit.NewSession(os.Getenv("REDDIT_USER_AGENT"))

	// assign OAuth and NoAuth sessions to User struct
	a.OAuthSession = oAuthSession
	a.NoAuthSession = session

	return a, nil
}

// GetUserTrophies method to retrieve slice of user trophies
func (a *AuthSessions) GetUserTrophies(redditObj *User) error {
	// get trophies of reddit user
	trophies, err := a.OAuthSession.UserTrophies(redditObj.Info.RedditData.Username)
	if err != nil {
		return err
	}
	// initialize new slice to store trophies
	var trophySlice []string

	// check if the user has no trophies
	if len(trophies) <= 0 {
		// create empty slice to prevent null value
		redditObj.Info.RedditData.Trophies = []string{""}
		return nil
	}

	// iterate over trophies object to add only trophy name to trophySlice
	for _, trophy := range trophies {
		trophySlice = append(trophySlice, trophy.Name)
	}
	// assign trophySlice to User struct
	redditObj.Info.RedditData.Trophies = trophySlice
	return nil
}

// GetRecentPostsFromSubreddit method to watch and pull last 5 posts from subreddit to match 2FA code
func (a *AuthSessions) GetRecentPostsFromSubreddit(redditObj *User) (*User, error) {
	// get 5 newest submissions from the subreddit
	submissions, err := a.OAuthSession.SubredditSubmissions(TwoFASubredditName, "new", geddit.ListingOptions{Count: 5})
	if err != nil {
		return redditObj, err
	}

	// iterate over the submissions
	for _, submission := range submissions {
		// check to ensure both author and 2FA code match
		if submission.Author == redditObj.Info.RedditData.Username && submission.Title == redditObj.Info.TwoFAData.StoredTwoFACode {
			// assign submission title (posted 2FA code) to user struct
			redditObj.Info.TwoFAData.PostedTwoFACode = submission.Title
			if redditObj.Info.TwoFAData.StoredTwoFACode == redditObj.Info.TwoFAData.PostedTwoFACode {
				// flip bool flag once 2FA code validated
				redditObj.Info.TwoFAData.IsValidated = true
				return redditObj, nil
			}
		}
	}
	// if no 2FA match return error message
	err = errors.New("[!] 2FA code not matched")
	return redditObj, err
}

// GetAboutInfo method to retrieve general information about user
func (a *AuthSessions) GetAboutInfo(redditObj *User) error {
	// get about information of reddit user
	redditProfile, err := a.OAuthSession.AboutRedditor(redditObj.Info.RedditData.Username)
	if err != nil {
		return err
	}

	// store select reddit profile info in user struct
	redditObj.Info.RedditData.CommentKarma = redditProfile.CommentKarma
	redditObj.Info.RedditData.LinkKarma = redditProfile.LinkKarma
	redditObj.Info.RedditData.AccountCreatedUTC = redditProfile.Created

	return nil
}

// GetSubmittedInfo method to retrieve slice of user's submitted posts
func (a *AuthSessions) GetSubmittedInfo(redditObj *User) error {
	// get submissions of reddit user
	submissions, err := a.NoAuthSession.RedditorSubmissions(redditObj.Info.RedditData.Username, geddit.ListingOptions{Count: 25})
	if err != nil {
		return err
	}

	// TODO:
	// initialize new map to store subreddits and associated score
	// var subredditMap map[string]int

	// initialize new slice to store subreddit names user has submitted to
	var subredditSlice []string

	// check if the user has not submitted anything
	if len(submissions) <= 0 {
		// create empty slice to prevent null value
		redditObj.Info.RedditData.Subreddits = []string{""}
		return nil
	}

	// iterate over submissions object to add subreddit name to subredditSlice
	for _, submission := range submissions {
		subredditSlice = append(subredditSlice, submission.Subreddit)
	}

	// return a unique slice version of the subredditSlice
	uniqueSubredditSlice := removeDuplicates(subredditSlice)

	// assign uniqueSubredditSlice to user struct
	redditObj.Info.RedditData.Subreddits = uniqueSubredditSlice

	return nil
}
