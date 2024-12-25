// auth/auth_test.go
package auth_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "healthcare-blockchain/auth"
    "healthcare-blockchain/database"
)

func TestRegisterUser(t *testing.T) {
    database.Connect()
    database.DB.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT, password TEXT)")
    database.DB.Exec("DELETE FROM users")

    router := gin.Default()
    auth.RegisterRoutes(router)

    userJSON := `{"username": "testuser", "password": "testpassword"}`
    req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBufferString(userJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }
}