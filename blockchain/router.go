package blockchain

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "healthcare-blockchain/database"
    "healthcare-blockchain/models"
    "github.com/dgrijalva/jwt-go"
    "healthcare-blockchain/config"
    "net/http"
    "strconv"
    "time"
    "fmt"
)

var blockchain = NewBlockchain()

func GetBlockchain(c *gin.Context) {
    blocks := blockchain.GetAllBlocks()
    c.JSON(http.StatusOK, blocks)
}


func GetBlockchainByPatientID(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(config.AppConfig.JWTSecret), nil
    })
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        return
    }
    patientID, ok := claims["user_id"].(string)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in token"})
        return
    }
    var filteredBlocks []*Block
    for _, block := range blockchain.GetAllBlocks() {
        if block.PatientID == patientID {
            filteredBlocks = append(filteredBlocks, block)
        }
    }
    c.JSON(http.StatusOK, filteredBlocks)
}

func AddBlock(c *gin.Context) {
    var newBlock models.BlockMetadata
    if err := c.ShouldBindJSON(&newBlock); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newBlock.TransactionID = uuid.New().String()
    newBlock.BlockID = uuid.New().String()
    newBlock.Timestamp = time.Now().Unix()
    var count int64
    if err := database.DB.Model(&models.BlockMetadata{}).Count(&count).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate block index"})
        return
    }
    newBlock.BlockIndex = uint(count + 1)
    if err := database.DB.Create(&newBlock).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add block"})
        return
    }
    blockchain.AddBlock(newBlock.Data, "passphrase", strconv.Itoa(int(newBlock.PatientID)), newBlock.TransactionID)
    c.JSON(http.StatusOK, gin.H{"message": "Block added successfully"})
}

func RegisterRoutes(r *gin.Engine) {
    blockchainGroup := r.Group("/blockchain")
    {
        blockchainGroup.GET("/", GetBlockchain)
        blockchainGroup.GET("/patient", GetBlockchainByPatientID)
        blockchainGroup.POST("/add", AddBlock)
    }
}