package db

import (
	"context"
	"fmt"

	"github.com/chirag1807/websocket-go/config"
	"github.com/jackc/pgx/v5"
)

func DBConnection() (conn *pgx.Conn, err error) {
	DATABASE_URL := "postgresql://" + config.Config.Database.Username + ":" + config.Config.Database.Password + "@127.0.0.1:" + config.Config.Database.Port + "/" + config.Config.Database.Name + "?sslmode=" + config.Config.Database.SSLMode
	connConfig, err := pgx.ParseConfig(DATABASE_URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	conn, err = pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conn, nil
}
