package brand

import (
	"context"
	"errors"
	"net/http"

	"github.com/m-mahmoud-alsaid/prim-backend/internal/model"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/api/security"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/database"
)

type BrandService struct {
	brepo     *BrandRepository
	qexecuter database.Runner
}

func NewService(
	r database.Runner,
	br *BrandRepository,
) *BrandService {
	return &BrandService{
		qexecuter: r,
		brepo:     br,
	}
}

type CreateBrandInput struct {
	Name      string
	LogoURL   string
	LogoLabel string
}

func (bs *BrandService) CreateBrand(
	ctx context.Context,
	in CreateBrandInput,
) (*model.Brand, error) {
	brand := model.NewBrand(
		in.Name,
		in.LogoURL,
		in.LogoLabel,
	)
	err := bs.qexecuter.WithDB(
		ctx,
		func(db database.QueryExecutor) error {
			return bs.brepo.Create(
				ctx,
				db,
				brand,
			)
		},
	)
	if err != nil {
		mappedError := database.MapError(err)
		switch {
		case errors.Is(
			mappedError,
			database.ErrConflict,
		):
			return nil, security.NewSecureError(
				http.StatusConflict,
				security.CodeConflict,
				"resourse alread existed",
				err,
			)
		default:
			return nil, security.NewSecureError(
				http.StatusInternalServerError,
				security.CodeInternal,
				"internal server error",
				err,
			)
		}
	}

	return brand, nil
}
