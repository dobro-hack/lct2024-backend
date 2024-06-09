package storage

import (
	"context"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
)

func (db *Database) ListRoute(ctx context.Context) ([]model.Route, int, error) {
	i := make([]model.Route, 0)

	co, err := db.Bun.NewSelect().Model(&i).
		Relation("Park").
		Relation("Places").
		ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return i, co, nil
}
