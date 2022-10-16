package service

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/fir1/rest-api/internal/todo/model"
	"github.com/fir1/rest-api/pkg/erru"
	"time"
)

type CreateParams struct {
	Name        string       `valid:"required"`
	Description string       `valid:"required"`
	Status      model.Status `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (int, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, erru.ErrArgument{Wrapped: err}
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.ToDo{
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
		CreatedOn:   time.Now().UTC(),
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return entity.ID, err
}
