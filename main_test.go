package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestSignup(t *testing.T) {
    req, _ := http.NewRequest("POST", "/signup", strings.NewReader("username=testuser&password=testpassword"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    w := httptest.NewRecorder()
    signupPage(w, req)

    if w.Code != http.StatusSeeOther {
        t.Errorf("Expected redirect after signup, got %v", w.Code)
    }

    // Verify that the user was added
    if password, ok := users["testuser"]; !ok || password != "testpassword" {
        t.Errorf("Expected user 'testuser' to exist with password 'testpassword'")
    }
}

func TestLoginSuccessful(t *testing.T) {
    // Simulate a signup
    users["testuser"] = "testpassword"

    req, _ := http.NewRequest("POST", "/", strings.NewReader("username=testuser&password=testpassword"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    w := httptest.NewRecorder()
    loginPage(w, req)

    if w.Code != http.StatusSeeOther {
        t.Errorf("Expected redirect on successful login, got %v", w.Code)
    }
}

func TestLoginFailure(t *testing.T) {
    // Simulate a signup
    users["testuser"] = "testpassword"

    req, _ := http.NewRequest("POST", "/", strings.NewReader("username=testuser&password=wrongpassword"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    w := httptest.NewRecorder()
    loginPage(w, req)

    if w.Code != http.StatusUnauthorized {
        t.Errorf("Expected unauthorized on failed login, got %v", w.Code)
    }
}

func TestForgotPasswordPage(t *testing.T) {
    req, _ := http.NewRequest("GET", "/forgot-password", nil)
    w := httptest.NewRecorder()
    forgotPasswordPage(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status OK for forgot password page, got %v", w.Code)
    }
}