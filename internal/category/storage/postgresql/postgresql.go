package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"shop-product-service/internal/category/modal"
	storage3 "shop-product-service/internal/category/storage"
)

type storage struct {
	logger *slog.Logger
	db     *pgxpool.Pool
}

func NewStorage(client *pgxpool.Pool, logger *slog.Logger) storage3.Storage {
	return &storage{
		logger: logger,
		db:     client,
	}
}

func (s storage) FindOne(ctx context.Context, id int64) (modal.Category, error) {
	q := `SELECT id, name, parent_category from categories WHERE id = ($1)`

	var c modal.Category

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	if err := s.db.QueryRow(ctx, q, id).Scan(&c.Id, &c.Name, &c.ParentCategory); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Sprintf("SQL Query Error: %s, Code: %s", pgErr.Message, pgErr.Code)
			s.logger.Error(newErr)
		}
		return modal.Category{}, err
	}
	return c, nil

}

func (s storage) FindByCategory(ctx context.Context, categoryId int64) ([]modal.Category, error) {
	q := `SELECT id, name, parent_category from categories WHERE parent_category = ($1)`

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	categories := make([]modal.Category, 0)

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
		var c modal.Category

		err := rows.Scan(&c.Id, &c.Name, &c.ParentCategory)
		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (s storage) FindRootCategories(ctx context.Context) ([]modal.Category, error) {
	q := `SELECT id, name, parent_category from categories WHERE parent_category IS NULL`

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	categories := make([]modal.Category, 0)

	rows, err := s.db.Query(ctx, q)
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
		var c modal.Category

		err := rows.Scan(&c.Id, &c.Name, &c.ParentCategory)
		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (s storage) Create(ctx context.Context, c modal.Category) (int64, error) {
	q := `INSERT INTO public."categories" ( name, parent_category)
			VALUES ($1, $2) RETURNING id`

	s.logger.Info(fmt.Sprintf("SQL Query: %s", q))

	if err := s.db.QueryRow(ctx, q, c.Name, c.ParentCategory).Scan(&c.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			//pgErr = err.(*pgconn.PgError)
			if pgErr.Code == "23505" {
				return 0, errors.New("user already exists")
			}
			newErr := fmt.Sprintf("SQL Query Error: %s, Code: %s", pgErr.Message, pgErr.Code)
			s.logger.Error(newErr)
		}
		return 0, err
	}

	return c.Id, nil
}

func (s storage) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s storage) Update(ctx context.Context, category modal.Category) error {
	//TODO implement me
	panic("implement me")
}
