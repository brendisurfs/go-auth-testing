// GO AUTH PRACTICE
// in the theme  of today, fuck it we tryin authorization in go.

package main

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/spotify"
)

func main() {
	fmt.Println("we out here lmao")

	authApp()
}

// authorize app through spotify
func authApp() {

	clientID := Auth().CLIENT_ID
	clientSecret := Auth().CLIENT_SECRET

	goth.UseProviders(
		spotify.New(clientID, clientSecret, "callbackurlidk"),
	)

	gothic.SetState = func(r *http.Request) string {
		return r.URL.Query().Get("state")
	}
}
