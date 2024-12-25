package api

import (
    "net/http"
    "healthcare-blockchain/blockchain"
    "github.com/gin-gonic/gin"
)

var blockchainInstance = blockchain.NewBlockchain()

func AddBlock(c *gin.Context) {
    var request struct {
        Data          string `json:"data"`
        Passphrase    string `json:"passphrase"`
        PatientID     string `json:"patient_id"`
        TransactionID string `json:"transaction_id"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := blockchainInstance.AddBlock(request.Data, request.Passphrase, request.PatientID, request.TransactionID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Block added successfully"})
}