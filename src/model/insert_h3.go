package model

import (
	"context"
	"fmt"

	"github.com/kajiLabTeam/mr-platform-contents-server/lib"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/uber/h3-go/v4"
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
		return err
	}
	return nil
}

func InsertCellToContentIdRelations(cells []h3.Cell, contentId string) error {
	ctx, driver, err := lib.ConnectNeo4j()
	if err != nil {
		return err
	}
	defer func() { err = lib.HandleClose(ctx, driver, err) }()
	session := driver.NewSession(ctx, neo4jSessionConfig)
	defer session.Close(ctx)

	err = insertH3CellToContentIdRelation(ctx, session, cells, contentId)
	if err != nil {
		return err
	}
	return nil
}

func insertH3CellToContentIdRelation(ctx context.Context, session neo4j.SessionWithContext, cells []h3.Cell, contentId string) error {
	query := `MATCH (c:Content {content: $contentId}) `
	for i := range cells {
		cell := h3.Cell(0)
		cell = cells[i]

		query += `MATCH `
		if i > 0 {
			query += fmt.Sprintf(`(h%d)-[:Child_Cell]->`, cell.Resolution()-1)
		}
		query += fmt.Sprintf(`(h%d:H3_Cell_%d {cell: $cell%d}) `, i, cell.Resolution(), i)

		query += fmt.Sprintf(`MERGE (h%d)-[:H3CellToContentId]->(c) `, i)

		query += fmt.Sprintf(`WITH h%d,c `, i)
	}
	query += `RETURN c`

	params := map[string]interface{}{
		"contentId": contentId,
	}
	for index := range cells {
		params[fmt.Sprintf("cell%d", index)] = cells[index].String()
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
