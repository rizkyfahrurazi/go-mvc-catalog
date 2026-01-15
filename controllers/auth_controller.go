package controllers

import (
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Login",
	}

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			data["Error"] = "Email dan password wajib diisi"
		} else if email != "admin@mail.com" || password != "123456" {
			data["Error"] = "Email atau password salah"
		} else {
			// simulasi login sukses
			http.SetCookie(w, &http.Cookie{
				Name:  "is_logged_in",
				Value: "true",
				Path:  "/",
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles(
		"views/layouts/auth.html",
		"views/auth/login.html",
	))
	tmpl.ExecuteTemplate(w, "auth", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// fake register success → langsung login
		http.SetCookie(w, &http.Cookie{
			Name:  "is_logged_in",
			Value: "true",
			Path:  "/",
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"views/layouts/auth.html",
		"views/auth/register.html",
	))

	data := map[string]interface{}{
		"Title": "Register",
	}

	tmpl.ExecuteTemplate(w, "auth", data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "is_logged_in",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
