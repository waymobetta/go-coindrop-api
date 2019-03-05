package reddit

import (
	"errors"
	"os"

	"github.com/jzelinskie/geddit"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/types"
)

var (
	// VerificationSubredditName name of subreddit used for verification
	VerificationSubredditName = "testing_QA_adChain"
	// ErrVerificationNotMatch .
	ErrVerificationNotMatch = errors.New("Verification code does not match")
)

// NewRedditAuth initializes reddit OAuth session
func NewRedditAuth() (*AuthSessions, error) {
	// initialize OAuth session with credentials as environment variables
	oAuthSession, err := geddit.NewOAuthSession(
		os.Getenv("REDDIT_CLIENT_ID"),
		os.Getenv("REDDIT_CLIENT_SECRET"),
		os.Getenv("REDDIT_USER_AGENT"),
		"http://metax.io",
	)
	if err != nil {
		return nil, err
	}

	// create new auth token for confidential clients (personal scripts/apps).
	err = oAuthSession.LoginAuth(os.Getenv("REDDIT_USERNAME"), os.Getenv("REDDIT_PASSWORD"))
	if err != nil {
		return nil, err
	}
	// create new unauthenticated session
	session := geddit.NewSession(os.Getenv("REDDIT_USER_AGENT"))

	// assign OAuth and NoAuth sessions to User struct
	return &AuthSessions{
		OAuthSession:  oAuthSession,
		NoAuthSession: session,
	}, nil
}

// GetRedditUserTrophies method to retrieve slice of user trophies
func (a *AuthSessions) GetRedditUserTrophies(user *types.User) error {
	// get trophies of reddit user
	trophies, err := a.OAuthSession.UserTrophies(user.Social.Reddit.Username)
	if err != nil {
		return err
	}
	// initialize new slice to store trophies
	var trophySlice []string

	// iterate over trophies object to add only trophy name to trophySlice
	for _, trophy := range trophies {
		trophySlice = append(trophySlice, trophy.Name)
	}

	if len(trophySlice) == 0 {
		trophySlice = []string{""}
	}

	// assign trophySlice to User struct
	user.Social.Reddit.Trophies = trophySlice

	return nil
}

// GetRecentPostsFromSubreddit method to watch and pull last 5 posts from subreddit to match verification code
func (a *AuthSessions) GetRecentPostsFromSubreddit(user *types.User) error {
	// get 5 newest submissions from the subreddit
	submissions, err := a.OAuthSession.SubredditSubmissions(VerificationSubredditName, "new", geddit.ListingOptions{Count: 1})
	if err != nil {
		return err
	}

	// iterate over the submissions
	for _, submission := range submissions {
		// check to ensure both author and verification code match
		if submission.Author == user.Social.Reddit.Username &&
			submission.Title == user.Social.Reddit.Verification.ConfirmedVerificationCode {
			// assign submission title (posted verification code) to user struct
			user.Social.Reddit.Verification.PostedVerificationCode = submission.Title
			if user.Social.Reddit.Verification.ConfirmedVerificationCode == user.Social.Reddit.Verification.PostedVerificationCode {
				// flip bool flag once verification code validated
				user.Social.Reddit.Verification.Verified = true
				return nil
			}
		}
	}
	// if no verification match return error message
	log.Errorln("[reddit] Verification code not matched")
	return ErrVerificationNotMatch
}

// GetAboutInfo method to retrieve general information about user
func (a *AuthSessions) GetAboutInfo(user *types.User) error {
	// get about information of reddit user
	redditProfile, err := a.OAuthSession.AboutRedditor(user.Social.Reddit.Username)
	if err != nil {
		return err
	}

	// store select reddit profile info in user struct
	user.Social.Reddit.CommentKarma = redditProfile.CommentKarma
	user.Social.Reddit.LinkKarma = redditProfile.LinkKarma

	return nil
}

// GetSubmittedInfo method to retrieve slice of user's submitted posts
func (a *AuthSessions) GetSubmittedInfo(user *types.User) error {
	// get submissions of reddit user
	submissions, err := a.NoAuthSession.RedditorSubmissions(user.Social.Reddit.Username, geddit.ListingOptions{Count: 25})
	if err != nil {
		return err
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
	user.Social.Reddit.Subreddits = uniqueSubredditSlice

	return nil
}

// GetOverview method to retrieve overview of user account
// TODO:
// func (u *User) GetOverview() *User {
// 	overviewURL := fmt.Sprintf("https://www.reddit.com/user/%s/overview.json", u.Reddit.Username)
// 	return u
// }
