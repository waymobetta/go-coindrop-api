package reddit

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
	"github.com/waymobetta/go-coindrop-api/db"
)

var (
	// VerificationSubredditName name of subreddit used for 2FA
	VerificationSubredditName = "testing_QA_adChain"
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
func (a *AuthSessions) GetUserTrophies(user *db.User) error {
	// get trophies of reddit user
	trophies, err := a.OAuthSession.UserTrophies(user.Info.RedditData.Username)
	if err != nil {
		return err
	}
	// initialize new slice to store trophies
	var trophySlice []string

	// iterate over trophies object to add only trophy name to trophySlice
	for _, trophy := range trophies {
		trophySlice = append(trophySlice, trophy.Name)
	}
	// assign trophySlice to User struct
	user.Info.RedditData.Trophies = trophySlice
	return nil
}

// GetRecentPostsFromSubreddit method to watch and pull last 5 posts from subreddit to match 2FA code
func (a *AuthSessions) GetRecentPostsFromSubreddit(user *db.User) (*db.User, error) {
	// get 5 newest submissions from the subreddit
	submissions, err := a.OAuthSession.SubredditSubmissions(VerificationSubredditName, "new", geddit.ListingOptions{Count: 1})
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	// iterate over the submissions
	for _, submission := range submissions {
		// check to ensure both author and 2FA code match
		if submission.Author == user.Info.RedditData.Username && submission.Title == user.Info.RedditData.VerificationData.StoredVerificationCode {
			// assign submission title (posted 2FA code) to user struct
			user.Info.RedditData.VerificationData.PostedVerificationCode = submission.Title
			if user.Info.RedditData.VerificationData.StoredVerificationCode == user.Info.RedditData.VerificationData.PostedVerificationCode {
				// flip bool flag once 2FA code validated
				user.Info.RedditData.VerificationData.IsVerified = true
				return user, nil
			}
		}
	}
	// if no 2FA match return error message
	err = fmt.Errorf("[!] 2FA code not matched")
	return user, err
}

// GetAboutInfo method to retrieve general information about user
func (a *AuthSessions) GetAboutInfo(user *db.User) (*db.User, error) {
	// get about information of reddit user
	redditProfile, err := a.OAuthSession.AboutRedditor(user.Info.RedditData.Username)
	if err != nil {
		return user, err
	}

	// store select reddit profile info in user struct
	user.Info.RedditData.CommentKarma = redditProfile.CommentKarma
	user.Info.RedditData.LinkKarma = redditProfile.LinkKarma
	user.Info.RedditData.AccountCreatedUTC = redditProfile.Created

	return user, nil
}

// GetSubmittedInfo method to retrieve slice of user's submitted posts
func (a *AuthSessions) GetSubmittedInfo(user *db.User) (*db.User, error) {
	// get submissions of reddit user
	submissions, err := a.NoAuthSession.RedditorSubmissions(user.Info.RedditData.Username, geddit.ListingOptions{Count: 25})
	if err != nil {
		return user, err
	}

	// TODO:
	// initialize new map to store subreddits and associated score
	// var subredditMap map[string]int

	// initialize new slice to store subreddit names user has submitted to
	var subredditSlice []string

	// iterate over submissions object to add subreddit name to subredditSlice
	for _, submission := range submissions {
		subredditSlice = append(subredditSlice, submission.Subreddit)
	}

	// return a unique slice version of the subredditSlice
	uniqueSubredditSlice := removeDuplicates(subredditSlice)

	// assign uniqueSubredditSlice to user struct
	user.Info.RedditData.Subreddits = uniqueSubredditSlice

	return user, nil
}

// GetOverview method to retrieve overview of user account
// TODO:
// func (u *User) GetOverview() *User {
// 	overviewURL := fmt.Sprintf("https://www.reddit.com/user/%s/overview.json", u.RedditData.Username)
// 	return u
// }
