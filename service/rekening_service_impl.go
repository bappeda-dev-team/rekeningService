package service

import (
	"context"
	"database/sql"
	"errors"
	"rekeningService/helper"
	"rekeningService/model/domain"
	"rekeningService/model/web"
	"rekeningService/repository"

	"github.com/go-playground/validator/v10"
)

type RekeningServiceImpl struct {
	RekeningRepository repository.RekeningRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewRekeningServiceImpl(rekeningRepository repository.RekeningRepository, db *sql.DB, validator *validator.Validate) *RekeningServiceImpl {
	return &RekeningServiceImpl{
		RekeningRepository: rekeningRepository,
		DB:                 db,
		Validator:          validator,
	}
}

func (service *RekeningServiceImpl) Create(ctx context.Context, request web.RekeningCreateRequest) (web.RekeningResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return web.RekeningResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RekeningResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rekeningDomain := domain.Rekening{
		KodeRekening: request.KodeRekening,
		NamaRekening: &request.NamaRekening,
		Tahun:        request.Tahun,
	}

	rekeningDomain, err = service.RekeningRepository.Create(ctx, tx, rekeningDomain)
	if err != nil {
		return web.RekeningResponse{}, err
	}

	return web.RekeningResponse{
		Id:           rekeningDomain.Id,
		KodeRekening: rekeningDomain.KodeRekening,
		NamaRekening: rekeningDomain.NamaRekening,
		Tahun:        rekeningDomain.Tahun,
	}, nil
}

func (service *RekeningServiceImpl) Update(ctx context.Context, request web.RekeningUpdateRequest) (web.RekeningResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return web.RekeningResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RekeningResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Cek apakah rekening ada
	_, err = service.RekeningRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return web.RekeningResponse{}, errors.New("rekening tidak ditemukan")
		}
		return web.RekeningResponse{}, err
	}

	rekeningDomain := domain.Rekening{
		Id:           request.Id,
		KodeRekening: request.KodeRekening,
		NamaRekening: &request.NamaRekening,
		Tahun:        request.Tahun,
	}

	rekeningDomain, err = service.RekeningRepository.Update(ctx, tx, rekeningDomain)
	if err != nil {
		return web.RekeningResponse{}, err
	}

	return web.RekeningResponse{
		Id:           rekeningDomain.Id,
		KodeRekening: rekeningDomain.KodeRekening,
		NamaRekening: rekeningDomain.NamaRekening,
		Tahun:        rekeningDomain.Tahun,
	}, nil
}

func (service *RekeningServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	// Cek apakah rekening ada
	_, err = service.RekeningRepository.FindById(ctx, tx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("rekening tidak ditemukan")
		}
		return err
	}

	err = service.RekeningRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *RekeningServiceImpl) FindById(ctx context.Context, id int) (web.RekeningResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RekeningResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rekeningDomain, err := service.RekeningRepository.FindById(ctx, tx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return web.RekeningResponse{}, errors.New("rekening tidak ditemukan")
		}
		return web.RekeningResponse{}, err
	}

	return web.RekeningResponse{
		Id:           rekeningDomain.Id,
		KodeRekening: rekeningDomain.KodeRekening,
		NamaRekening: rekeningDomain.NamaRekening,
		Tahun:        rekeningDomain.Tahun,
	}, nil
}

func (service *RekeningServiceImpl) FindAll(ctx context.Context) ([]web.RekeningResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.RekeningResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rekeningDomains, err := service.RekeningRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.RekeningResponse{}, err
	}

	var rekeningResponses []web.RekeningResponse
	for _, rekening := range rekeningDomains {
		rekeningResponses = append(rekeningResponses, web.RekeningResponse{
			Id:           rekening.Id,
			KodeRekening: rekening.KodeRekening,
			NamaRekening: rekening.NamaRekening,
			Tahun:        rekening.Tahun,
		})
	}

	return rekeningResponses, nil
}
