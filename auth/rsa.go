package auth

import (
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "math/big"
    mrand "math/rand"
)

func GenerateRSAKeys(email string, password string, keySize int) (string, string, error) {
    combined := email + password
    hash := sha256.Sum256([]byte(combined))
    seed := new(big.Int).SetBytes(hash[:])
    reader := mrand.New(mrand.NewSource(seed.Int64()))

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