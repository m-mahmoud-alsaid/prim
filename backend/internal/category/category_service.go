package category

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/m-mahmoud-alsaid/prim-backend/internal/model"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/api/security"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/database"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/utils"
)

type CategoryService struct {
	crepository *CategoryRepository
	qexecuter   database.Runner
}

func NewService(
	qexecuter database.Runner,
	r *CategoryRepository,
) *CategoryService {
	return &CategoryService{
		crepository: r,
		qexecuter:   qexecuter,
	}
}

type CreateCategoryInput struct {
	Name     string
	ParentID *uuid.UUID
}

func (cs *CategoryService) CreateCategory(
	ctx context.Context,
	in CreateCategoryInput,
) (*model.Category, error) {
	slug := utils.Slugify(in.Name)

	category := model.NewCategory(
		in.Name,
		slug,
		in.ParentID,
	)

	err := cs.qexecuter.WithDB(
		ctx,
		func(db database.QueryExecutor) error {
			return cs.crepository.Create(
				ctx,
				db,
				category,
			)
		},
	)

	if err != nil {
		mappedErr := database.MapError(err)
		switch {
		case errors.Is(
			mappedErr,
			database.ErrConflict,
		):
			return nil, security.NewSecureError(
				http.StatusConflict,
				security.CodeConflict,
				"resource already exists",
				err,
			)
		default:
			return nil, security.NewSecureError(
				http.StatusInternalServerError,
				security.CodeInternal,
				"failed to create a new resource",
				err,
			)
		}
	}
	return category, err
}
