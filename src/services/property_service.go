package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.GetPropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {

	return &PropertyService{base: (NewBaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.GetPropertyResponse](cfg, []preload{{string: "Category"}}))}
}

func (s *PropertyService) CreateProperty(c context.Context, r *dto.CreatePropertyRequest) (*dto.GetPropertyResponse, error) {
	return s.base.Create(c, r)
}

func (s *PropertyService) UpdateProperty(c context.Context, id int, r *dto.UpdatePropertyRequest) (*dto.GetPropertyResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *PropertyService) DeleteProperty(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *PropertyService) GetProperty(c context.Context, id int) (*dto.GetPropertyResponse, error) {
	return s.base.GetById(c, id)
}

func (s *PropertyService) GetPropertyByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GetPropertyResponse], error) {
	return s.base.GetByFilter(c, req)
}
