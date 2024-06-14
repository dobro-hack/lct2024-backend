package storage

import (
	"context"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
)

func (db *Database) SaveDevice(ctx context.Context, req model.Device) error {
	_, err := db.Bun.NewInsert().Model(&req).Ignore().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
