package model

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID uuid.UUID

	Name    string
	LogoURL string
	LogoAlt string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Tag struct {
	ID uuid.UUID

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ProductMedia struct {
	ID        uuid.UUID
	VariantID uuid.UUID

	URL      string
	Alt      string
	MimeType string

	// the first media always the main
	SortOrder int

	Width    int
	Height   int
	FileSize int64

	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ProductCategory struct {
	ID uuid.UUID

	Name string
	Slug string

	ParentID *uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductVariantStatus string

const (
	ProductVariantActive       ProductVariantStatus = "active"
	ProductVariantInactive     ProductVariantStatus = "inactive"
	ProductVariantDiscontinued ProductVariantStatus = "discontinued"
)

type ProductVariant struct {
	ID        uuid.UUID
	ProductID uuid.UUID

	SKU      string
	Slug     string
	Status   ProductVariantStatus
	Price    int64
	Currency string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ProductStatus string

const (
	ProductStatusActive   string = "active"
	ProductStatusDraft    string = "draft"
	ProductStatusArchived string = "archived"
)

type Product struct {
	ID      uuid.UUID
	BrandID uuid.UUID

	Title            string
	ShortDescription string
	Description      string
	Status           ProductStatus

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
