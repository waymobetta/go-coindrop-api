package gostackoverflow

// AboutProfileResponse struct contins profile info
type AboutProfileResponse struct {
	Items []struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		ViewCount               int    `json:"view_count"`
		DownVoteCount           int    `json:"down_vote_count"`
		UpVoteCount             int    `json:"up_vote_count"`
		AnswerCount             int    `json:"answer_count"`
		QuestionCount           int    `json:"question_count"`
		AccountID               int    `json:"account_id"`
		IsEmployee              bool   `json:"is_employee"`
		LastModifiedDate        int    `json:"last_modified_date"`
		LastAccessDate          int    `json:"last_access_date"`
		ReputationChangeYear    int    `json:"reputation_change_year"`
		ReputationChangeQuarter int    `json:"reputation_change_quarter"`
		ReputationChangeMonth   int    `json:"reputation_change_month"`
		ReputationChangeWeek    int    `json:"reputation_change_week"`
		ReputationChangeDay     int    `json:"reputation_change_day"`
		Reputation              int    `json:"reputation"`
		CreationDate            int    `json:"creation_date"`
		UserType                string `json:"user_type"`
		UserID                  int    `json:"user_id"`
		AboutMe                 string `json:"about_me"`
		Location                string `json:"location"`
		WebsiteURL              string `json:"website_url"`
		Link                    string `json:"link"`
		ProfileImage            string `json:"profile_image"`
		DisplayName             string `json:"display_name"`
	} `json:"items"`
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}

// AssociatedCommunitiesResponse struct contains info from
type AssociatedCommunitiesResponse struct {
	Items []struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		QuestionCount  int    `json:"question_count"`
		AnswerCount    int    `json:"answer_count"`
		LastAccessDate int    `json:"last_access_date"`
		CreationDate   int    `json:"creation_date"`
		AccountID      int    `json:"account_id"`
		Reputation     int    `json:"reputation"`
		UserID         int    `json:"user_id"`
		SiteURL        string `json:"site_url"`
		SiteName       string `json:"site_name"`
	} `json:"items"`
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}
