package lib

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var Neo4jConfiguration common.Neo4jConfiguration

func ConnectNeo4j() (context.Context, neo4j.DriverWithContext, error) {
	ctx := context.Background()
	dbUri := os.Getenv("NEO4J_URL_API")
	dbUser := os.Getenv("NEO4J_USER")
	dbPassword := os.Getenv("NEO4J_PASSWORD")

	Neo4jConfiguration = common.Neo4jConfiguration{
		URL:      dbUri,
		Username: dbUser,
		Password: dbPassword,
	}

	driver, err := neo4j.NewDriverWithContext(Neo4jConfiguration.URL, neo4j.BasicAuth(Neo4jConfiguration.Username, Neo4jConfiguration.Password, ""))
	if err != nil {
		return ctx, nil, err
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return ctx, nil, err
	}

	return ctx, driver, nil
}

func HandleClose(ctx context.Context, closer interface{ Close(context.Context) error }, previousError error) error {
	err := closer.Close(ctx)
	if err == nil {
		return previousError
	}
	if previousError == nil {
		return err
	}
	return fmt.Errorf("%v closure error occurred:\n%s\ninitial error was:\n%w", reflect.TypeOf(closer), err.Error(), previousError)
}
