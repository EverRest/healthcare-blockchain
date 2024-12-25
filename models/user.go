package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID           uint           `json:"id" gorm:"primaryKey"`
    Username     string         `json:"username"`
    Password     string         `json:"password"`
    PasswordHash string         `json:"-"`
    Role         string         `json:"role"`
    Email        string         `json:"email"`
    PublicKey    string         `json:"public_key"`
    PrivateKey   string         `json:"private_key"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`
}

const (
    RoleAdmin  = "admin"
    RoleDoctor = "doctor"
    RoleUser   = "user"
)