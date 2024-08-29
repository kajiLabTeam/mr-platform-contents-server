package model

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/lib"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InsertContentForNeo4j(contentId string) error {
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
	return nil
}

func createContentNode(ctx context.Context, session neo4j.SessionWithContext, contentId string) error {
	query := `MERGE (c:Content {content: $contentId})
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
		return errors.New("failed to execute query")
	}
	return nil
}

func InsertCellToContentRelations(cellAndReses []common.CellAndRes, contentId string) error {
	ctx, driver, err := lib.ConnectNeo4j()
	if err != nil {
		return err
	}
	defer func() { err = lib.HandleClose(ctx, driver, err) }()
	session := driver.NewSession(ctx, neo4jSessionConfig)
	defer session.Close(ctx)

	err = insertH3CellToContentRelation(ctx, session, cellAndReses, contentId)
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
MERGE (h0)-[:CellHas]->(c)
WITH h0, c
MATCH (h1:H3_Cell_1 {cell: $cell1}) -[:Child_Cell]->(h0:H3_Cell_0)
WITH h0, h1, c
MERGE (h1)-[:CellHas]->(c)
WITH h1, c
MATCH (h2:H3_Cell_2 {cell: $cell2}) -[:Child_Cell]->(h1:H3_Cell_1)
WITH h1, h2, c
MERGE (h2)-[:CellHas]->(c)
WITH h2, c
MATCH (h3:H3_Cell_3 {cell: $cell3}) -[:Child_Cell]->(h2:H3_Cell_2)
WITH h2, h3, c
MERGE (h3)-[:CellHas]->(c)
RETURN c
*/
func insertH3CellToContentRelation(ctx context.Context, session neo4j.SessionWithContext, cellAndReses []common.CellAndRes, contentId string) error {
	query := `MATCH (c:Content {content: $contentId}) `
	for i := 0; i < len(cellAndReses); i++ {
		query += fmt.Sprintf(`MATCH (h%d:H3_Cell_%d {cell: $cell%d})`, i, cellAndReses[i].Res, i)

		if i > 0 {
			query += fmt.Sprintf(`-[:Child_Cell]->(h%d:H3_Cell_%d) WITH h%d, h%d, c `, i-1, cellAndReses[i-1].Res, i-1, i)
		} else {
			query += `WITH h0, c `
		}

		query += fmt.Sprintf(`MERGE (h%d)-[:CellHas]->(c) `, i)
		if i < len(cellAndReses)-1 {
			query += fmt.Sprintf(`WITH h%d, c `, i)
		}
	}
	query += `RETURN c`

	params := map[string]interface{}{
		"contentId": contentId,
	}
	for i := 0; i < len(cellAndReses); i++ {
		params["cell"+strconv.Itoa(i)] = cellAndReses[i].Cell
	}

	fmt.Println(query)
	fmt.Println(params)

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
