package neo4j_subscription_storage

import (
	"context"
    "log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/neo4j"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	testPassword := "letmein!"

	neo4jContainer, err := neo4j.RunContainer(ctx,
		testcontainers.WithImage("docker.io/neo4j:4.4"),
		neo4j.WithAdminPassword(testPassword),
		neo4j.WithNeo4jSetting("dbms.tx_log.rotation.size", "42M"),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := neo4jContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
}
