package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLoginPage(t *testing.T) {
    req, _ := http.NewRequest("GET", "/login", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status OK, got %v", w.Code)
    }
}