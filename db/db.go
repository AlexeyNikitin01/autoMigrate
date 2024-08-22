package db

import (
	"context"
	"fmt"

	"task2/config"

	"github.com/jackc/pgx/v5"
)

func Connect(cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.TODO(),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			cfg.Postgres.User,
			cfg.Postgres.Pass,
			cfg.Postgres.Port,
			cfg.Postgres.Host,
			cfg.Postgres.Dbname,
		))
	if err != nil {
		fmt.Println("no connect db")
		return nil, err
	}
	return conn, nil
}
