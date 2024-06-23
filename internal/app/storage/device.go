package storage

import (
	"context"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
)

func (db *Database) ListDevice(ctx context.Context) ([]model.Device, error) {
	i := make([]model.Device, 0)

	err := db.Bun.NewSelect().Model(&i).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}
