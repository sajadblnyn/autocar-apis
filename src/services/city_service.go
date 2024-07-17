package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.GetCityResponse]
}

func NewCityService(cfg *config.Config) *CityService {

	return &CityService{base: (NewBaseService[models.City, dto.CreateUpdateCityRequest, dto.CreateUpdateCityRequest, dto.GetCityResponse](cfg, []preload{{string: "Country"}}))}
}

func (s *CityService) CreateCity(c context.Context, r *dto.CreateUpdateCityRequest) (*dto.GetCityResponse, error) {
	return s.base.Create(c, r)
}

func (s *CityService) UpdateCity(c context.Context, id int, r *dto.CreateUpdateCityRequest) (*dto.GetCityResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CityService) DeleteCity(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CityService) GetCity(c context.Context, id int) (*dto.GetCityResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CityService) GetCityByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GetCityResponse], error) {
	return s.base.GetByFilter(c, req)
}
