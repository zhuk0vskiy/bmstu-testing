//go:build integration_test

package integration_tests

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testDbInstance *pgxpool.Pool

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase()
	defer testDB.TearDown()
	testDbInstance = testDB.DbInstance
	err := SeedTestData(testDbInstance)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(m.Run())
}
