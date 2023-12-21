package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func TestMain(m *testing.M) {
    ctx := context.Background()
    os.Setenv("JWT_ISSUER", "test_issuer")
    os.Setenv("JWT_SECRET", "test_secret")
    os.Setenv("JWT_VALIDITY_IN_HOURS", "244")
    req := testcontainers.ContainerRequest{
        Image:        "postgres:latest",
        ExposedPorts: []string{"5432/tcp"},
        Env:          map[string]string{"POSTGRES_PASSWORD": "password"},
        WaitingFor:   wait.ForLog("database system is ready to accept connections"),
    }

    postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        log.Fatalf("Failed to start container: %s", err)
    }
    defer postgresContainer.Terminate(ctx)

    // Retrieve the container's mapped port
    port, _ := postgresContainer.MappedPort(ctx, "5432")
    testDBURL := fmt.Sprintf("postgres://postgres:password@localhost:%s/postgres?sslmode=disable", port.Port())
    log.Println("Test DB URL:", testDBURL)
    log.Println("JWT_ISSUER:", os.Getenv("JWT_ISSUER"))
    log.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
    log.Println("JWT_VALIDITY_IN_HOURS:", os.Getenv("JWT_VALIDITY_IN_HOURS"))
    // Set up the database connection
    db = utils.SetupDBConnection(testDBURL)

    code := m.Run()

    db.Close()
    os.Exit(code)
}

func generateValidAuthToken() string {
	user := &models.User{
		UserName: "testuser",
		Email:    "test@example.com",
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		// Print the error and return an empty string or handle it as appropriate
		fmt.Println("Error generating token:", err)
		return ""
	}

	return token
}
