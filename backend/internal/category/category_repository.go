package category

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/m-mahmoud-alsaid/prim-backend/internal/model"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/database"
)

type Filter struct {
	ID   *uuid.UUID
	Slug *string
}

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
	if err != nil {
		return fmt.Errorf("create category: %w", err)
	}
	return err
}

func (cr *CategoryRepository) Get(
	ctx context.Context,
	qe database.QueryExecutor,
	filter Filter,
) (*model.Category, error) {
	var category model.Category

	query :=
		`
			SELECT
				id,
				name,
				slug,
				parent_id,
				created_at,
				updated_at
			FROM
				categories
			WHERE
			deleted_at IS NULL
		`

	args := []any{}
	argID := 1
	if filter.ID != nil {
		query += fmt.Sprintf(" AND id = $%d", argID)
		args = append(args, *filter.ID)
		argID++
	}

	if filter.Slug != nil {
		query += fmt.Sprintf(" AND slug = $%d", argID)
		args = append(args, *filter.Slug)
		argID++
	}

	err := qe.QueryRow(
		ctx,
		query,
		args...,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Slug,
		&category.ParentID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("get category: %w", err)
	}

	return &category, nil
}
