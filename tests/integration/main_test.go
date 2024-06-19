package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/tests/testhelpers"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
)

type AppTestSuite struct {
	suite.Suite
	db          *bun.DB
	pgContainer *testhelpers.PostgresContainer
}

// SetupSuite runs once before the suite starts and sets up the test environment.
func (suite *AppTestSuite) SetupSuite() {
	os.Setenv("ENV", "test")
	ctx := context.Background()

	var err error
	suite.pgContainer, err = testhelpers.CreatePostgresContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL container: %s", err)
	}

	suite.db, err = utils.SetupDBConnection(suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
}

// TearDownSuite runs once after all the tests in the suite have finished.
func (suite *AppTestSuite) TearDownSuite() {
	ctx := context.Background()
	if err := suite.pgContainer.Terminate(ctx); err != nil {
		log.Fatalf("Failed to terminate PostgreSQL container: %s", err)
	}

	suite.db.Close()
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
