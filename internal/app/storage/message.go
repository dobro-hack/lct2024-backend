package storage

import (
	"context"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
)

func (db *Database) SaveMessage(ctx context.Context, req model.Message) error {
	_, err := db.Bun.NewInsert().Model(&req).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ListMessage(ctx context.Context) ([]model.Message, int, error) {
	i := make([]model.Message, 0)

	co, err := db.Bun.NewSelect().Model(&i).
		ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return i, co, nil
}
