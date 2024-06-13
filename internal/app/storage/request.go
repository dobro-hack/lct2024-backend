package storage

import (
	"context"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
)

func (db *Database) SaveRequest(ctx context.Context, req model.SaveRequest) error {
	_, err := db.Bun.NewInsert().Model(&req).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
