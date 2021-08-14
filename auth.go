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

	// I. get client ID + client Secret from env.go
	clientID := Auth().CLIENT_ID
	clientSecret := Auth().CLIENT_SECRET

	// II. create provider for spotify
	goth.UseProviders(
		spotify.New(clientID, clientSecret, "callbackurlidk"),
	)

	// III. ???
	gothic.SetState = func(r *http.Request) string {
		return r.URL.Query().Get("state")
	}

	// IV. Profit

}
