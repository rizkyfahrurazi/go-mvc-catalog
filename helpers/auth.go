package helpers

import "net/http"

func IsLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("is_logged_in")
	if err != nil {
		return false
	}

	return cookie.Value == "true"
}
func GetUserID(r *http.Request) (int, bool) {
	session, _ := GetSession(r)

	userID, ok := session.Values["user_id"].(int)
	return userID, ok
}

