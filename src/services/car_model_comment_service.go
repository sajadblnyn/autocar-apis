package services

import (
	"context"
	"errors"

	"github.com/sajadblnyn/autocar-apis/apis/dto"
	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/constants"
	"github.com/sajadblnyn/autocar-apis/data/models"
	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {

	return &CarModelCommentService{base: (NewBaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse](cfg, []preload{
		{string: "User"}}))}
}

func (s *CarModelCommentService) CreateCarModelComment(c context.Context, r *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	r.UserId = int(c.Value(constants.UserIdKey).(float64))
	return s.base.Create(c, r)
}

func (s *CarModelCommentService) UpdateCarModelComment(c context.Context, id int, r *dto.UpdateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	comment, err := s.base.GetById(c, id)
	if err != nil {
		return nil, err
	}
	userId := int(c.Value(constants.UserIdKey).(float64))
	if comment.UserId != userId {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied, TechnicalMessage: "user do not have permission for edit this comment", Err: errors.New(service_errors.PermissionDenied)}
	}
	return s.base.Update(c, id, r)
}

func (s *CarModelCommentService) DeleteCarModelComment(c context.Context, id int) error {

	return s.base.Delete(c, id)

}

func (s *CarModelCommentService) GetCarModelComment(c context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return s.base.GetById(c, id)
}

func (s *CarModelCommentService) GetCarModelCommentByFilter(c context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelCommentResponse], error) {
	return s.base.GetByFilter(c, req)
}
