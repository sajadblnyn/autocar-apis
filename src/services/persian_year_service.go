package services

import (
	"context"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/data/models"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {

	return &PersianYearService{base: (NewBaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse](cfg, []preload{}))}
}

func (s *PersianYearService) CreatePersianYear(c context.Context, r *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Create(c, r)
}

func (s *PersianYearService) UpdatePersianYear(c context.Context, id int, r *dto.UpdatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Update(c, id, r)
}

func (s *PersianYearService) DeletePersianYear(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *PersianYearService) GetPersianYear(c context.Context, id int) (*dto.PersianYearResponse, error) {
	return s.base.GetById(c, id)
}

func (s *PersianYearService) GetPersianYearByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PersianYearResponse], error) {
	return s.base.GetByFilter(c, req)
}
