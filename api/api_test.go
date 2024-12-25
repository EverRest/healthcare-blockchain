package api_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "healthcare-blockchain/api"
)

func TestHealthCheck(t *testing.T) {
    router := api.SetupRouter()

    req, _ := http.NewRequest("GET", "/healthcheck", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }
}