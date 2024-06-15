package storage

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dobro-hack/lct2024-backend/internal/app/config"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
)

// Database represents database instance
type Database struct {
	Bun *bun.DB
	DB  *sql.DB
}

func createTLSConf() tls.Config {

	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile("root.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}

	return tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true, // needed for self signed certs
	}
}

// New returns new Database with specified dsn, handler and options.
func Open(ctx context.Context, cfg config.Config) (*Database, error) {
	tlsConf := createTLSConf()
	err := mysql.RegisterTLSConfig("custom", &tlsConf)
	if err != nil {
		return nil, err
	}

	dbConn, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	dbConn.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	dbConn.SetConnMaxLifetime(cfg.DB.ConnMaxLifetime)

	bundb := bun.NewDB(dbConn, mysqldialect.New(), bun.WithDiscardUnknownColumns())
	// To enable database debug uncomment next line
	bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	bundb.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName("lct2024")))

	db := &Database{bundb, dbConn}
	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil

}

// Close closes an active database.
func (db *Database) Close() error {
	if err := db.Bun.Close(); err != nil {
		return fmt.Errorf("closing database: %w", err)
	}

	return nil
}

// Ping pings the database. Used for healthchecks.
func (db *Database) Ping(ctx context.Context) error {
	if err := db.Bun.PingContext(ctx); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	return nil
}

// WithTx runs the specified function in a new transaction.
func (db *Database) WithTx(ctx context.Context, readOnly bool, f func(ctx context.Context, tx bun.Tx) error) error {
	return db.Bun.RunInTx(ctx, &sql.TxOptions{ReadOnly: readOnly}, f)
}
