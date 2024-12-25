package api_test

import (
    "os"
    "testing"
    "healthcare-blockchain/config"
    "healthcare-blockchain/database"
)

func TestMain(m *testing.M) {
    config.LoadConfig()
    database.Connect()
    database.Migrate(
        // Add models here if needed
    )
    os.Exit(m.Run())
}