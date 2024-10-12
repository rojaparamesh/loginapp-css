package main

import (
    "database/sql"
    "net/http"
    "log"
    "html/template"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/userdb")
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    r := gin.Default()
    r.Static("/static", "./static")
    r.LoadHTMLGlob("templates/*")

    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    r.GET("/signup", func(c *gin.Context) {
        c.HTML(http.StatusOK, "signup.html", nil)
    })

    r.POST("/signin", signIn)
    r.POST("/signup", signUp)
    r.GET("/forgot", func(c *gin.Context) {
        c.HTML(http.StatusOK, "forgot_password.html", nil)
    })
    r.POST("/reset_password", resetPassword)

    r.Run(":8080")
}

func signIn(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    var dbPassword string
    err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPassword)
    if err != nil {
        c.String(http.StatusUnauthorized, "Invalid username or password")
        return
    }

    // Here you should validate the password (use bcrypt for hashing)
    if password != dbPassword {
        c.String(http.StatusUnauthorized, "Invalid username or password")
        return
    }

    c.String(http.StatusOK, "Welcome, %s!", username)
}

func signUp(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    // Here you should hash the password (use bcrypt for security)
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error creating user")
        return
    }

    c.String(http.StatusOK, "User created successfully!")
}

func resetPassword(c *gin.Context) {
    username := c.PostForm("username")
    // Logic to reset the password (send email, etc.)
    c.String(http.StatusOK, "Password reset link sent to %s", username)
}