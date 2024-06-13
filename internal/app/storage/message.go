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
