package brand

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m-mahmoud-alsaid/prim-backend/internal/shared/validation"
	"github.com/m-mahmoud-alsaid/prim-backend/pkg/api"
)

type BrandHandler struct {
	bservice *BrandService
}

func NewHandler(
	sb *BrandService,
) *BrandHandler {
	return &BrandHandler{
		bservice: sb,
	}
}

type CreateBrandRequest struct {
	Name      string `json:"name" binding:"required" example:"apple"`
	LogoURL   string `json:"logo_url" binding:"required" example:"https://pictures.com/apple.png"`
	LogoLable string `json:"logo_label" binding:"required" example:"apple logo"`
}

type BrandResponse struct {
	ID        uuid.UUID `json:"id" example:"358b2e03-0b3f-40a4-8163-ebed0cb252ee"`
	Name      string    `json:"name" example:"nvidia"`
	LogoURL   string    `json:"logo_url" example:"https://pictures.com/nvidia.png"`
	LogoLabel string    `json:"logo_label" example:"nvidia logo"`
	CreatedAt string    `json:"created_at" example:"2026-07-01T05:04:38Z"`
	UpdatedAt string    `json:"udpated_at" example:"2026-07-01T05:04:38Z"`
}

// CreateBrand godoc
// @Summary create a new products brand
// @Description create a new products brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param brand body CreateBrandRequest true "brand data"
// @Failure 404 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Success 201 {object} api.DataResponse{data=BrandResponse}
// @Router /brands [post]
func (bh *BrandHandler) CreateBrand(c *gin.Context) {
	var req CreateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.ValidationError(c, err)
		return
	}

	in := CreateBrandInput{
		Name:      req.Name,
		LogoURL:   req.LogoURL,
		LogoLabel: req.LogoLable,
	}

	ctx := c.Request.Context()
	brand, err := bh.bservice.CreateBrand(
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
			Data: BrandResponse{
				ID:        brand.ID,
				Name:      brand.Name,
				LogoURL:   brand.LogoURL,
				LogoLabel: brand.LogoLabel,
				CreatedAt: brand.CreatedAt.Format(time.RFC3339),
				UpdatedAt: brand.UpdatedAt.Format(time.RFC3339),
			},
		},
	)

}
