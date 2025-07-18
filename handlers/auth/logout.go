package auth

import (
	"forum-app/app"
	"forum-app/middleware"
	"forum-app/models"
	"net/http"
)

// Logout handles user logout by clearing the session and redirecting to the login page.
func Logout(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middleware.UserKey).(*models.Users)

		if !ok || user == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		session, _ := app.DB.GetSession("userId", user.ID)
		if session != nil {
			app.DB.DeleteSession(session.ID)
		}

		expireCookie := http.Cookie{
			Name:   "auth-token",
			Value:  "",
			MaxAge: -1,
		}

		http.SetCookie(w, &expireCookie)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
