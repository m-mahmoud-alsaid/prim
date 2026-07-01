package brand

import (
	"context"
	"fmt"

	"github.com/m-mahmoud-alsaid/prim-backend/internal/model"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/database"
)

type BrandRepository struct {
}

func NewRepository() *BrandRepository {
	return &BrandRepository{}
}

func (br *BrandRepository) Create(
	ctx context.Context,
	qe database.QueryExecutor,
	brand *model.Brand,
) error {
	_, err := qe.Exec(
		ctx,
		`
		INSERT INTO brands(
			id,
			name,
			logo_url,
			logo_label,
			created_at,
			updated_at
		) 
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
		`,
		brand.ID,
		brand.Name,
		brand.LogoURL,
		brand.LogoLabel,
		brand.CreatedAt,
		brand.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create brand :%w", err)
	}
	return nil
}
