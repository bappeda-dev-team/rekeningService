package repository

import (
	"context"
	"database/sql"
	"rekeningService/model/domain"
)

type RekeningRepositoryImpl struct {
}

func NewRekeningRepositoryImpl() *RekeningRepositoryImpl {
	return &RekeningRepositoryImpl{}
}

func (repository *RekeningRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, rekening domain.Rekening) (domain.Rekening, error) {
	query := "INSERT INTO tb_rekening (kode_rekening, nama_rekening, tahun) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, rekening.KodeRekening, rekening.NamaRekening, rekening.Tahun)
	if err != nil {
		return domain.Rekening{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Rekening{}, err
	}

	rekening.Id = int(id)

	return rekening, nil
}

func (repository *RekeningRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, rekening domain.Rekening) (domain.Rekening, error) {
	query := "UPDATE tb_rekening SET kode_rekening = ?, nama_rekening = ?, tahun = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, rekening.KodeRekening, rekening.NamaRekening, rekening.Tahun, rekening.Id)
	if err != nil {
		return domain.Rekening{}, err
	}
	return rekening, nil
}

func (repository *RekeningRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM tb_rekening WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *RekeningRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Rekening, error) {
	query := "SELECT id, kode_rekening, nama_rekening, tahun FROM tb_rekening WHERE id = ?"
	row := tx.QueryRowContext(ctx, query, id)
	var rekening domain.Rekening
	err := row.Scan(&rekening.Id, &rekening.KodeRekening, &rekening.NamaRekening, &rekening.Tahun)
	if err != nil {
		return domain.Rekening{}, err
	}
	return rekening, nil
}

func (repository *RekeningRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Rekening, error) {
	query := "SELECT id, kode_rekening, nama_rekening, tahun FROM tb_rekening"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rekenings []domain.Rekening
	for rows.Next() {
		var r domain.Rekening
		err := rows.Scan(&r.Id, &r.KodeRekening, &r.NamaRekening, &r.Tahun)
		if err != nil {
			return nil, err
		}
		rekenings = append(rekenings, r)
	}
	return rekenings, nil
}
