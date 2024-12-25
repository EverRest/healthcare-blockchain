package models

import (
    "time"
    "gorm.io/gorm"
)

type BlockMetadata struct {
    BlockID       string         `json:"block_id" gorm:"type:text"`
    BlockIndex    uint           `json:"block_index"`
    PatientID     uint           `json:"patient_id"`
    TransactionID string         `json:"transaction_id"`
    Timestamp     int64          `json:"timestamp"`
    Data          string         `json:"data"`
    CreatedAt     time.Time
    UpdatedAt     time.Time
    DeletedAt     gorm.DeletedAt `gorm:"index"`
}