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

func (db *Database) GetRoute(ctx context.Context, routeId int) (model.Route, error) {
	i := model.Route{}

	err := db.Bun.NewSelect().Model(&i).
		Where("id = ?", routeId).
		Scan(ctx)
	if err != nil {
		return i, err
	}

	return i, nil
}

func (db *Database) SaveRoute(ctx context.Context, route model.Route) error {
	_, err := db.Bun.NewUpdate().Model(&route).Column("load").WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
