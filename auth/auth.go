// GO AUTH PRACTICE
// in the theme  of today, fuck it we tryin authorization in go.

package Auth

import (
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/spotify"
)

// authorize app through spotify
func AuthApp() {

	// I. get client ID + client Secret from env.go
	clientID := Env().CLIENT_ID
	clientSecret := Env().CLIENT_SECRET

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
