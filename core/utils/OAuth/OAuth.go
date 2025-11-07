package OAuth

import (
	"RedRock/core"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func InitOAuth() (*OAuth, error) {
	root := OAuth{}
	err := root.initOAuth()
	return &root, err
}

func (root *OAuth) initOAuth() error {
	//var err error
	root.userForm = UserInfo{}
	//ctx := context.Background()
	//root.provider, err = oidc.NewProvider(ctx, core.GlobalConf.OidcProvider)

	//if err != nil {
	//	return err
	//}

	root.oAuth2Config = oauth2.Config{
		ClientID:     core.GlobalConf.ClientID,
		ClientSecret: core.GlobalConf.ClientSecret,
		RedirectURL:  core.GlobalConf.RedirectURL,

		Scopes:   []string{"", "email", "profile"},
		Endpoint: github.Endpoint,
	}

	return nil
}

func (root *OAuth) OAuth2LoginLogic(c *gin.Context) {
	state := uuid.New().String()
	c.SetCookie("state",
		state,
		3600,
		"/",
		core.GlobalConf.Domain,
		false,
		true)
	authURL := root.oAuth2Config.AuthCodeURL(state)
	c.Redirect(http.StatusFound, authURL)
}

func (root *OAuth) CallBackLogic(c *gin.Context) (error, UserInfo) {
	state := c.Query("state")
	stateInCookie, _ := c.Cookie("state")

	if stateInCookie != state {
		//c.JSON(http.StatusBadRequest, gin.H{"Error": "StateIncorrect"})
		return fmt.Errorf("state Incorrect"), root.userForm
	}

	code := c.Query("code")
	token, err := root.oAuth2Config.Exchange(context.Background(), code)

	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"ExchangeTokenError": err.Error()})
		return err, root.userForm
	}

	client := root.oAuth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")

	if err != nil {
		return err, root.userForm
	}

	defer resp.Body.Close()

	userInfo, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, root.userForm
	}

	err = json.Unmarshal(userInfo, &root.userForm)
	if err != nil {
		return err, root.userForm
	}

	return nil, root.userForm
}
