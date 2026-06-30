package category

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m-mahmoud-alsaid/prim-backend/internal/shared/validation"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/api"
)

type CategoryHandler struct {
	cservice *CategoryService
}

func NewHandler(
	s *CategoryService,
) *CategoryHandler {
	return &CategoryHandler{
		cservice: s,
	}
}

type CreateCategoryRequest struct {
	Name     string     `json:"name" binding:"required"`
	ParentID *uuid.UUID `json:"parent_id" binding:"omitempty,uuid"`
}

type CreateCategoryResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	ParentID  *uuid.UUID `json:"parent_id"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.ValidationError(c, err)
		return
	}

	in := CreateCategoryInput{
		Name:     req.Name,
		ParentID: req.ParentID,
	}

	ctx := c.Request.Context()
	category, err := ch.cservice.CreateCategory(
		ctx,
		in,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(
		http.StatusCreated,
		api.DataResponse{
			Data: CreateCategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				Slug:      category.Slug,
				ParentID:  category.ParentID,
				CreatedAt: category.CreatedAt.Format(time.RFC3339),
				UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
			},
		},
	)
}
