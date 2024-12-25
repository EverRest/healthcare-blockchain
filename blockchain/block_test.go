package blockchain_test

import (
	"healthcare-blockchain/blockchain"
	"testing"
	"math/rand"
	"fmt"
)

func generateFakeData() string {
	receipts := []string{"Receipt A", "Receipt B", "Receipt C"}
	doctorVisits := []string{"Visit A", "Visit B", "Visit C"}
	insurances := []string{"Insurance A", "Insurance B", "Insurance C"}

	receipt := receipts[rand.Intn(len(receipts))]
	doctorVisit := doctorVisits[rand.Intn(len(doctorVisits))]
	insurance := insurances[rand.Intn(len(insurances))]

	return fmt.Sprintf("Receipt: %s, Doctor Visit: %s, Insurance: %s", receipt, doctorVisit, insurance)
}

func TestGenerateFakeBlocks(t *testing.T) {
	bc := blockchain.NewBlockchain()

	for i := 0; i < 20; i++ {
		fakeData := generateFakeData()
		bc.AddBlock(fakeData)
	}

	if len(bc.Blocks) != 21 { // Including the genesis block
		t.Fatalf("Expected 21 blocks, but got %d", len(bc.Blocks))
	}

	for i, block := range bc.Blocks {
		t.Logf("Block %d: %+v\n", i, block)
	}
}