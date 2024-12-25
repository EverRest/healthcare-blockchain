package blockchain_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "healthcare-blockchain/blockchain"
    "healthcare-blockchain/database"
)

func TestGetBlockchain(t *testing.T) {
    database.Connect()
    database.DB.Exec("DELETE FROM block_metadata") // Clean up the table before testing

    router := gin.Default()
    blockchain.RegisterRoutes(router)

    req, _ := http.NewRequest("GET", "/blockchain/", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }
}

func TestAddBlock(t *testing.T) {
    database.Connect()
    database.DB.Exec("DELETE FROM block_metadata") // Clean up the table before testing

    router := gin.Default()
    blockchain.RegisterRoutes(router)

    blockJSON := `{"block_id": 1, "block_index": 1, "patient_id": 1, "transaction_id": "tx123", "timestamp": 1234567890, "data": "test data"}`
    req, _ := http.NewRequest("POST", "/blockchain/add", bytes.NewBufferString(blockJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }
}