package storage

import (
	"context"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
)

func (db *Database) SaveRequest(ctx context.Context, req model.Request) error {
	_, err := db.Bun.NewInsert().Model(&req).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ListRequest(ctx context.Context) ([]model.Request, int, error) {
	i := make([]model.Request, 0)

	co, err := db.Bun.NewSelect().Model(&i).Relation("Person").Relation("Org").
		ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return i, co, nil
}
