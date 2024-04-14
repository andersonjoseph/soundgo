package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/andersonjoseph/soundgo/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

type config struct {
	port               int
	dbConnString       string
	sessionTokenSecret []byte
}

func getConfig() config {
	dbConnString := flag.String("db", "", "database connection string")
	sessionTokenSecret := flag.String("sessionTokenSecret", "secret", "the secret to sign JWT session token, should at least be 32 characters long, but the longer, the better")
	port := flag.Int("port", -1, "http port")

	flag.Parse()

	return config{
		port:               *port,
		dbConnString:       *dbConnString,
		sessionTokenSecret: []byte(*sessionTokenSecret),
	}
}

func run() error {
	config := getConfig()

	conn, err := pgxpool.New(context.Background(), config.dbConnString)

	if err != nil {
		return err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return err
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	srv := server.NewServer(
		conn,
		logger,
		config.sessionTokenSecret,
	)

	logger.Info("server listening", "port", config.port)

	addr := fmt.Sprintf(":%d", config.port)

	return srv.Listen(addr)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
