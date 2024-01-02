package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/floriwan/srcm/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/steam"
)

func Registration(c *gin.Context) {
	goth.UseProviders(
		steam.New(os.Getenv("STEAM_KEY"), "/steam/callback"),
	)

	values := c.Request.URL.Query()
	values.Add("provider", "steam")
	c.Request.URL.RawQuery = values.Encode()

	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		// TODO add user to database
		log.Printf("user: %+v\n", gothUser)
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(c.Writer, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title":    config.GlobalConfig.HomepageName,
		"subtitle": "login",
	})
}

func SteamCallback(c *gin.Context) {
	fmt.Printf("SteamCallback\n")
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("steam user: %+v\n", user)

}

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
