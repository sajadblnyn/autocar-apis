package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.GetCountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {

	return &CountryService{base: (NewBaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.GetCountryResponse](cfg))}
}

func (s *CountryService) CreateCountry(c context.Context, r *dto.CreateUpdateCountryRequest) (*dto.GetCountryResponse, error) {
	return s.base.Create(c, r)
}

func (s *CountryService) UpdateCountry(c context.Context, id int, r *dto.CreateUpdateCountryRequest) (*dto.GetCountryResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CountryService) DeleteCountry(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CountryService) GetCountry(c context.Context, id int) (*dto.GetCountryResponse, error) {
	return s.base.GetById(c, id)
}
