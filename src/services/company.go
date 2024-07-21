package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {

	return &CompanyService{base: (NewBaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse](cfg, []preload{{string: "Country"}}))}
}

func (s *CompanyService) CreateCompany(c context.Context, r *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Create(c, r)
}

func (s *CompanyService) UpdateCompany(c context.Context, id int, r *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *CompanyService) DeleteCompany(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CompanyService) GetCompany(c context.Context, id int) (*dto.CompanyResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CompanyService) GetCompanyByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CompanyResponse], error) {
	return s.base.GetByFilter(c, req)
}
