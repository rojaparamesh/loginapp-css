package main

import (
    "html/template"
    "net/http"
)

var users = make(map[string]string)

func main() {
    http.HandleFunc("/", loginPage)
    http.HandleFunc("/signup", signupPage)
    http.HandleFunc("/forgot-password", forgotPasswordPage)
    http.ListenAndServe(":8080", nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")
        if _, ok := users[username]; ok && users[username] == password {
            http.Redirect(w, r, "/success", http.StatusSeeOther)
            return
        }
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
    renderTemplate(w, "login.html")
}

func signupPage(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")
        users[username] = password
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    renderTemplate(w, "signup.html")
}

func forgotPasswordPage(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "forgot_password.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, _ := template.ParseFiles("templates/" + tmpl)
    t.Execute(w, nil)
}