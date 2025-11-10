package service

import (
	"context"
	"rekeningService/model/web"
)

type RekeningService interface {
	Create(ctx context.Context, request web.RekeningCreateRequest) (web.RekeningResponse, error)
	Update(ctx context.Context, request web.RekeningUpdateRequest) (web.RekeningResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.RekeningResponse, error)
	FindAll(ctx context.Context) ([]web.RekeningResponse, error)
}
