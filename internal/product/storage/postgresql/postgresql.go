package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"shop-product-service/internal/product/modal"
	storage2 "shop-product-service/internal/product/storage"
)

type storage struct {
	logger *slog.Logger
	db     *pgxpool.Pool
}

func NewStorage(client *pgxpool.Pool, logger *slog.Logger) storage2.Storage {

	return &storage{
		logger: logger,
		db:     client,
	}
}

func (s storage) FindOne(ctx context.Context, id string) (modal.Product, error) {
	q := `SELECT id, name, price, count, images from product."product" WHERE id = ($1)`

	var p modal.Product

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	if err := s.db.QueryRow(ctx, q, id).Scan(); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Sprintf("SQL Query Error: %s, Code: %s", pgErr.Message, pgErr.Code)
			s.logger.Error(newErr)
		}
		return modal.Product{}, err
	}
	return p, nil

}

func (s storage) FindByCategory(ctx context.Context, categoryId int64) ([]modal.Product, error) {
	q := `SELECT id, name, price, count, images,description,categoryId from product."product" WHERE categoryId = ($1)`

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	products := make([]modal.Product, 0)

	rows, err := s.db.Query(ctx, q, categoryId)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Sprintf("SQL Query Error: %s, Code: %s", pgErr.Message, pgErr.Code)
			s.logger.Error(newErr)
		}
		return nil, err
	}

	for rows.Next() {
		var p modal.Product

		err := rows.Scan(&p.Uuid, &p.Name, &p.Price, &p.Count, &p.Image, &p.Description, &p.CategoryId)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s storage) Create(ctx context.Context, product modal.Product) (string, error) {
	q := `INSERT INTO product."product" (id, name, price, count, images, description, categoryId)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	if err := s.db.QueryRow(ctx, q, product.Uuid, product.Name, product.Count, product.Image, product.Description, product.CategoryId).Scan(); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			if pgErr.Code == "23505" {
				return "", errors.New("user already exists")
			}
			newErr := fmt.Sprintf("SQL Query Error: %s, Code: %s", pgErr.Message, pgErr.Code)
			s.logger.Error(newErr)
		}
		return "", err
	}

	return product.Uuid, nil
}

func (s storage) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s storage) Update(ctx context.Context, product modal.Product) error {
	//TODO implement me
	panic("implement me")
}
