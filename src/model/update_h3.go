package model

import (
	"context"

	"github.com/kajiLabTeam/mr-platform-contents-server/lib"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func RemoveRelationH3CellToContentIdForNeo4j(contentId string) error {
	ctx, driver, err := lib.ConnectNeo4j()
	if err != nil {
		return err
	}
	defer func() { err = lib.HandleClose(ctx, driver, err) }()
	session := driver.NewSession(ctx, neo4jSessionConfig)
	defer session.Close(ctx)

	err = createContentNode(ctx, session, contentId)
	if err != nil {
		return err
	}

	err = removeRelationH3CellToContentId(ctx, session, contentId)
	if err != nil {
		return err
	}

	return nil
}

/*
最終的に作られるクエリ
MATCH (c:Content {content: $contentId})
MATCH (h0:H3_Cell_0 {cell: $cell0})
WITH h0, c
MERGE (h0)-[r:H3CellToContentId]->(c)
DELETE r
WITH h0, c
MATCH (h1:H3_Cell_1 {cell: $cell1}) -[:Child_Cell]->(h0:H3_Cell_0)
WITH h0, h1, c
MERGE (h1)-[r:H3CellToContentId]->(c)
DELETE r
WITH h1, c
MATCH (h2:H3_Cell_2 {cell: $cell2}) -[:Child_Cell]->(h1:H3_Cell_1)
WITH h1, h2, c
MERGE (h2)-[r:H3CellToContentId]->(c)
DELETE r
WITH h2, c
MATCH (h3:H3_Cell_3 {cell: $cell3}) -[:Child_Cell]->(h2:H3_Cell_2)
WITH h2, h3, c
MERGE (h3)-[r:H3CellToContentId]->(c)
DELETE r
RETURN c
*/
func removeRelationH3CellToContentId(ctx context.Context, session neo4j.SessionWithContext, contentId string) error {
	// contentIdと繋がる関連性"H3CellToContentId"を削除
	query := `MATCH (c:Content {content: $contentId})
						MATCH ()-[r:H3CellToContentId]->(c) DELETE r
						RETURN c`

	params := map[string]interface{}{
		"contentId": contentId,
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		txResult, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		_, err = txResult.Collect(ctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}
