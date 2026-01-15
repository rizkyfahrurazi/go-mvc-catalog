package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("super-secret-key"))



func GetSession(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "online-shop-session")
}
