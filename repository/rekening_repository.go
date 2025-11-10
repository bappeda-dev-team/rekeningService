package repository

import (
	"context"
	"database/sql"
	"rekeningService/model/domain"
)

type RekeningRepository interface {
	Create(ctx context.Context, tx *sql.Tx, rekening domain.Rekening) (domain.Rekening, error)
	Update(ctx context.Context, tx *sql.Tx, rekening domain.Rekening) (domain.Rekening, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Rekening, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Rekening, error)
}
