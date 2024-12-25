package blockchain

import (
    "errors"
)

type Blockchain struct {
    Blocks []*Block
}

func NewBlockchain() *Blockchain {
    genesisBlock, _ := NewBlock(0, "Genesis Block", "0", "passphrase", "0")
    return &Blockchain{Blocks: []*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data, passphrase, patientID, transactionID string) error {
    if data == "" {
        return errors.New("data cannot be empty")
    }
    previousBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock, err := NewBlock(len(bc.Blocks), data, previousBlock.Hash, passphrase, patientID)
    if err != nil {
        return err
    }
    bc.Blocks = append(bc.Blocks, newBlock)
    err = AddBlockMetadata(newBlock.Index, patientID, transactionID, data)
    if err != nil {
        return err
    }
    return nil
}

func (bc *Blockchain) GetAllBlocks() []*Block {
    return bc.Blocks
}