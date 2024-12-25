package blockchain

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/hex"
    "encoding/pem"
    "os"
)

func EncryptData(publicKeyPath, data string) (string, error) {
    pubKeyFile, err := os.ReadFile(publicKeyPath)
    if err != nil {
        return "", err
    }
    block, _ := pem.Decode(pubKeyFile)
    publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return "", err
    }
    encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), []byte(data))
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(encrypted), nil
}

func DecryptData(privateKeyPath, encryptedData string) (string, error) {
    privKeyFile, err := os.ReadFile(privateKeyPath)
    if err != nil {
        return "", err
    }
    block, _ := pem.Decode(privKeyFile)
    privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return "", err
    }
    data, err := hex.DecodeString(encryptedData)
    if err != nil {
        return "", err
    }
    decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
    if err != nil {
        return "", err
    }
    return string(decrypted), nil
}