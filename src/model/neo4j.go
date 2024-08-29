package model

import (
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var neo4jSessionConfig neo4j.SessionConfig = neo4j.SessionConfig{
	DatabaseName: os.Getenv("NEO4J_DATABASE"),
}
