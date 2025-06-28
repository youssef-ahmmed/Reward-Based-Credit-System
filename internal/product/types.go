package product

type CategorySimple struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductFilters struct {
	CategoryID string
	IsOffer    *bool
	MinPoints  int
	MaxPoints  int
}

type PaginationMeta struct {
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
	TotalItems   int `json:"totalItems"`
	ItemsPerPage int `json:"itemsPerPage"`
}

type CreateCategoryRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	ParentCategoryID string `json:"parentCategoryId"`
}

type UpdateCategoryRequest struct {
	Name             *string `json:"name"`
	Description      *string `json:"description"`
	ParentCategoryID *string `json:"parentCategoryId"`
}

type CategoryResponse struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	ParentCategoryID *string `json:"parentCategoryId,omitempty"`
}

type CreateProductRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	CategoryID       string   `json:"categoryId"`
	RedemptionPoints int      `json:"redemptionPoints"`
	StockQuantity    int      `json:"stockQuantity"`
	IsOffer          bool     `json:"isOffer"`
	ImageURL         *string  `json:"imageUrl,omitempty"`
	Tags             []string `json:"tags"`
}

type UpdateProductRequest struct {
	Name             *string  `json:"name"`
	Description      *string  `json:"description"`
	CategoryID       *string  `json:"categoryId"`
	RedemptionPoints *int     `json:"redemptionPoints"`
	StockQuantity    *int     `json:"stockQuantity"`
	IsOffer          *bool    `json:"isOffer"`
	ImageURL         *string  `json:"imageUrl,omitempty"`
	Tags             []string `json:"tags"`
}

type ProductResponse struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Category         *CategorySummary `json:"category,omitempty"`
	RedemptionPoints int              `json:"redemptionPoints"`
	StockQuantity    int              `json:"stockQuantity"`
	IsOffer          bool             `json:"isOffer"`
	ImageURL         *string          `json:"imageUrl,omitempty"`
	Tags             []string         `json:"tags"`
	CreatedAt        string           `json:"createdAt"`
}

type CategorySummary struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryDetailsResponse struct {
	Category   *Category          `json:"category"`
	Products   []*ProductResponse `json:"products"`
	Pagination *PaginationMeta    `json:"pagination"`
}
