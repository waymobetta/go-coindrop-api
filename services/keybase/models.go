package keybase

// KeybaseData profile info
type Keybase struct {
	Bio                string `json:"bio"`
	Location           string `json:"location"`
	FullName           string `json:"full_name"`
	GithubUsername     string `json:"github_username"`
	TwitterUsername    string `json:"twitter_username"`
	HackerNewsUsername string `json:"hackernews_username"`
}

/*
// curl commands

/// TODO
// POST: GET KEYBASE INFO

# update user + Keybase info
curl -H "Content-type: application/json" -d '{"info": {"reddit_data": {"reddit_username": "qa_adchain_registry"}}}' 'localhost:5000/api/v1/updatekeybaseinfo'
*/
