package service

import (
	"context"
	"errors"
	"learning_tools/ddd-auth2-example/domain/entity"
	"learning_tools/ddd-auth2-example/domain/repo"
	"learning_tools/ddd-auth2-example/domain/repo/specification"
	"learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
)

type Merchant struct {
	merchantRepo repo.MerchantRepo
}

func (m *Merchant) CreateMerChant(ctx context.Context, data *entity.Merchant) error {
	if err := data.CheckBase(); err != nil {
		return err
	}
	_, err := m.merchantRepo.QueryMerChant(ctx, specification.NewMerchantSpecificationByAPPID(data.APPID))
	if errors.Is(err, hcode.ResourcesNotFindErr) {
		return m.merchantRepo.CreateMerChant(ctx, data)
	} else if err != nil {
		return err
	} else {
		return hcode.ResourcesAlreadyExistsErr
	}
}

func (m *Merchant) QueryMerChant(ctx context.Context, repo repo.MerChantSpecificationRepo) (data *entity.Merchant, err error) {
	return m.merchantRepo.QueryMerChant(ctx, repo)
}
