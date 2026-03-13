package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	var config = Config {
		addr: ":8080",
		db: DbConfig {},
	}

	// Database connection
	var ctx = context.Background()
	var dbStr = "postgres://postgres:password@localhost:5110/postgres"
	var conn, errConn = pgx.Connect(ctx, dbStr)
	if errConn != nil {
		slog.Error("Postgres connection error.", "error", errConn)
		panic(errConn)
	}
	slog.Info("Postgres connected.")
	defer conn.Close(ctx)

	var app = Application {
		config: config,
		db: conn,
	}

	var handler = app.mount()
	// if err := app.run(handler); err != nil {
	// 	log.Printf("Error to run the server: %s", err)
	// }
	var err = app.run(handler)
	if err != nil {
		// log.Printf("Error to run the server: %s", err)
		slog.Error("Server failed to start. Error: ", "error", err)
		os.Exit(1)
	}
}
