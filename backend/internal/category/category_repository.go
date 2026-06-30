package category

import (
	"context"

	"github.com/m-mahmoud-alsaid/prim-backend/internal/model"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/database"
)

type CategoryRepository struct {
}

func NewRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (cr *CategoryRepository) Create(
	ctx context.Context,
	qe database.QueryExecutor,
	category *model.Category,
) error {
	_, err := qe.Exec(
		ctx,
		`
		INSERT INTO
		categories (
			id,
			name,
			slug,
			parent_id,
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
		category.ID,
		category.Name,
		category.Slug,
		category.ParentID,
		category.CreatedAt,
		category.UpdatedAt,
	)
	return err
}
