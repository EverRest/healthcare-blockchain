package auth

import (
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "errors"
    "golang.org/x/crypto/bcrypt"
    "healthcare-blockchain/database"
    "healthcare-blockchain/models"
    "math/rand"
)

func CreateUser(username string, password string, email string, role string, keySize int) error {
    if role != models.RoleAdmin && role != models.RoleDoctor && role != models.RoleUser {
        return errors.New("invalid role")
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    publicKey, privateKey, err := GenerateRSAKeys(email, password, keySize)
    if err != nil {
        return err
    }
    user := models.User{
        Username:     username,
        Password:     password,
        PasswordHash: string(hashedPassword),
        Role:         role,
        Email:        email,
        PublicKey:    publicKey,
        PrivateKey:   privateKey,
    }
    return database.DB.Create(&user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    result := database.DB.Where("username = ?", username).First(&user)
    return &user, result.Error
}

func GenerateRSAKeysFromPassword(password string, keySize int) (string, string, error) {
    hash := sha256.Sum256([]byte(password))
    reader := rand.New(rand.NewSource(int64(hash[0])))
    privateKey, err := rsa.GenerateKey(reader, keySize)
    if err != nil {
        return "", "", err
    }
    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
    privateKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: privateKeyBytes,
    })
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
    if err != nil {
        return "", "", err
    }
    publicKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    })
    return string(publicKeyPEM), string(privateKeyPEM), nil
}