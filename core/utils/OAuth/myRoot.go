package OAuth

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OAuth struct {
	provider     *oidc.Provider
	oAuth2Config oauth2.Config
	userForm     UserInfo
}

type UserInfo struct {
	Login             string      `json:"login"`
	ID                int64       `json:"id"`
	NodeId            string      `json:"node_id"`
	AvatarUrl         string      `json:"avatar_url"`
	GravatarId        string      `json:"gravatar_id"`
	Url               string      `json:"url"`
	HtmlUrl           string      `json:"html_url"`
	FollowersUrl      string      `json:"followers_url"`
	FollowingUrl      string      `json:"following_url"`
	GistsUrl          string      `json:"gists_url"`
	StarredUrl        string      `json:"starred_url"`
	SubscriptionsUrl  string      `json:"subscriptions_url"`
	OrganizationsUrl  string      `json:"organizations_url"`
	ReposUrl          string      `json:"repos_url"`
	EventsUrl         string      `json:"events_url"`
	ReceivedEventsUrl string      `json:"received_events_url"`
	Type              string      `json:"type"`
	UserViewType      string      `json:"user_view_type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               interface{} `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	NotificationEmail interface{} `json:"notification_email"`
	PublicRepos       int64       `json:"public_repos"`
	PublicGists       int64       `json:"public_gists"`
	Followers         int64       `json:"followers"`
	Following         int64       `json:"following"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
}
