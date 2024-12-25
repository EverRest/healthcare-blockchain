package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "github.com/google/uuid"
    "healthcare-blockchain/database"
    "healthcare-blockchain/models"
    "strconv"
    "time"
)

func hashData(data string) string {
    h := sha256.New()
    h.Write([]byte(data))
    return hex.EncodeToString(h.Sum(nil))
}

func AddBlockMetadata(index int, patientID, transactionID, data string) error {
    patientIDUint, err := strconv.ParseUint(patientID, 10, 32)
    if err != nil {
        return err
    }
    hashedData := hashData(data)
    metadata := models.BlockMetadata{
        BlockID:       uuid.New().String(),
        BlockIndex:    uint(index),
        PatientID:     uint(patientIDUint),
        TransactionID: transactionID,
        Timestamp:     time.Now().Unix(),
        Data:          hashedData,
    }
    return database.DB.Create(&metadata).Error
}

func GetMetadataByPatientID(patientID string) ([]models.BlockMetadata, error) {
    patientIDUint, err := strconv.ParseUint(patientID, 10, 32)
    if err != nil {
        return nil, err
    }
    var metadata []models.BlockMetadata
    result := database.DB.Where("patient_id = ?", uint(patientIDUint)).Find(&metadata)
    return metadata, result.Error
}