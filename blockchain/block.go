package blockchain

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "io"
    "strconv"
    "time"
)

type Block struct {
    Index         int    `json:"index"`
    Timestamp     string `json:"timestamp"`
    EncryptedData string `json:"encrypted_data"`
    PreviousHash  string `json:"previous_hash"`
    Hash          string `json:"hash"`
    PatientID     string `json:"patient_id"`
}

func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.EncryptedData + b.PreviousHash
    h := sha256.New()
    h.Write([]byte(record))
    return hex.EncodeToString(h.Sum(nil))
}

func encrypt(data, passphrase string) (string, error) {
    block, err := aes.NewCipher([]byte(passphrase))
    if err != nil {
        return "", err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

    return hex.EncodeToString(ciphertext), nil
}

func NewBlock(index int, data, previousHash, passphrase, patientID string) (*Block, error) {
    encryptedData, err := encrypt(data, passphrase)
    if err != nil {
        return nil, err
    }

    block := &Block{
        Index:         index,
        Timestamp:     time.Now().UTC().String(),
        EncryptedData: encryptedData,
        PreviousHash:  previousHash,
        PatientID:     patientID,
    }
    block.Hash = block.CalculateHash()
    return block, nil
}