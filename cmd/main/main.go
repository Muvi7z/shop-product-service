package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"os"
	"shop-product-service/internal/config"
	"shop-product-service/internal/product/service"
	"shop-product-service/internal/product/storage/postgresql"
)

func main() {
	//router := gin.Default()

	cfg := config.GetConfig()

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Storage.Username, cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Database)

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	productStorage := postgresql.NewStorage(db, logger)

	productService := service.New(logger, productStorage)

	productService.GetProductByCategory(ctx, 1)

}
