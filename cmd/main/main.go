package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"os"
	"shop-product-service/internal/config"
	"shop-product-service/internal/product/modal"
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

	productService.AddProduct(ctx, modal.CreateProductDTO{
		Name:        "Honor 9",
		Price:       20333,
		Count:       10,
		Image:       nil,
		Description: "Notebook",
		CategoryId:  1,
	})

}
