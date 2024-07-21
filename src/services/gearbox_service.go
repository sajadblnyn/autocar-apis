package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {

	return &GearboxService{base: (NewBaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse](cfg, []preload{}))}
}

func (s *GearboxService) CreateGearbox(c context.Context, r *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Create(c, r)
}

func (s *GearboxService) UpdateGearbox(c context.Context, id int, r *dto.UpdateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *GearboxService) DeleteGearbox(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *GearboxService) GetGearbox(c context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(c, id)
}

func (s *GearboxService) GetGearboxByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(c, req)
}
