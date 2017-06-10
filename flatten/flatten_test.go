package flatten

import (
	"testing"
)

type TestUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func TestGetPayload(t *testing.T) {
	actual, err := getPayload("https://api.github.com/users/jbelmont/repos")
	testUser := TestUser{
		Login:             "jbelmont",
		ID:                2974744,
		AvatarURL:         "https://avatars0.githubusercontent.com/u/2974744?v=3",
		GravatarID:        "",
		URL:               "https://api.github.com/users/jbelmont",
		HTMLURL:           "https://github.com/jbelmont",
		FollowersURL:      "https://api.github.com/users/jbelmont/followers",
		FollowingURL:      "https://api.github.com/users/jbelmont/following{/other_user}",
		GistsURL:          "https://api.github.com/users/jbelmont/gists{/gist_id}",
		StarredURL:        "https://api.github.com/users/jbelmont/starred{/owner}{/repo}",
		SubscriptionsURL:  "https://api.github.com/users/jbelmont/subscriptions",
		OrganizationsURL:  "https://api.github.com/users/jbelmont/orgs",
		ReposURL:          "https://api.github.com/users/jbelmont/repos",
		EventsURL:         "https://api.github.com/users/jbelmont/events{/privacy}",
		ReceivedEventsURL: "https://api.github.com/users/jbelmont/received_events",
		Type:              "User",
		SiteAdmin:         false,
	}
	if err != nil {
		t.Error(err)
	}

	firstObj := actual[0]
	if firstObj.Owner.ID != testUser.ID {
		t.Errorf("%d should equal %d", firstObj.ID, testUser.ID)
	}
}
