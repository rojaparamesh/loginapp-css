package main

import (
    "database/sql"
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
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
    r.LoadHTMLGlob("templates/*") // Load HTML templates

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

    // Check for existing user
    var existingUser string
    err := db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUser)
    if err != nil && err != sql.ErrNoRows {
        log.Printf("Error checking existing user: %v", err)
        c.String(http.StatusInternalServerError, "Internal Server Error")
        return
    }

    if existingUser != "" {
        c.String(http.StatusConflict, "Username already exists")
        return
    }

    // Hash the password
    hashedPassword, err := hashPassword(password)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        c.String(http.StatusInternalServerError, "Error hashing password")
        return
    }

    // Insert new user
    _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
    if err != nil {
        log.Printf("Error inserting user: %v", err)
        c.String(http.StatusInternalServerError, "Error creating user: %v", err)
        return
    }

    c.String(http.StatusOK, "User created successfully!")
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func resetPassword(c *gin.Context) {
    username := c.PostForm("username")
    // Logic to reset the password (send email, etc.)
    c.String(http.StatusOK, "Password reset link sent to %s", username)
}