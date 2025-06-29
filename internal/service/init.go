package service

import (
	"Start/internal/store"
	"Start/internal/types"
	"gorm.io/gorm"
)

type CreditPackageService interface {
	GetAllCreditPackages(page, limit int, activeFilter *bool) ([]types.CreditCreditPackageResponse, types.PaginationMeta, error)
	GetCreditPackageByID(id string) (*types.CreditCreditPackageResponse, error)
	CreateCreditPackage(input types.CreateCreditPackageRequest) (*types.CreditCreditPackageResponse, error)
	UpdateCreditPackages(id string, input types.UpdateCreditPackageRequest) (*types.CreditCreditPackageResponse, error)
	DeleteCreditPackage(id string) error
}

type PurchaseService interface {
	CreatePurchase(userID string, input types.CreatePurchaseRequest) (*types.PurchaseResponse, error)
	GetUserPurchases(userID string, status string, page int, limit int) ([]types.PurchaseResponse, types.PaginationMeta, error)
	GetPurchaseByID(userID string, purchaseID string) (*types.PurchaseResponse, error)
	CountTotalPurchases() (int, error)
	SumCreditsIssued() (int, error)
}

type ProductService interface {
	GetAllProducts(filters types.ProductFilters, page, limit int, sortBy, sortOrder string) ([]store.Product, types.PaginationMeta, error)
	SearchProducts(query string, filters types.ProductFilters, page, limit int) ([]store.Product, types.PaginationMeta, error)
	CreateProduct(input *types.CreateProductRequest) (*types.ProductResponse, error)
	UpdateProduct(id string, input *types.UpdateProductRequest) (*types.ProductResponse, error)
	DeleteProduct(id string) error
}

type CategoryService interface {
	CreateCategory(c *types.CreateCategoryRequest) (*types.CategoryResponse, error)
	GetAllCategories(parentID *string) ([]store.Category, error)
	GetCategoryDetails(categoryID string, page, limit int) (*types.CategoryDetailsResponse, error)
	UpdateCategory(id string, input *types.UpdateCategoryRequest) (*types.CategoryResponse, error)
	DeleteCategory(id string) error
}

type RedemptionService interface {
	CreateRedemption(userID string, input types.CreateRedemptionRequest) (*types.RedemptionResponse, error)
	GetRedemptionByID(userID, id string) (*types.RedemptionResponse, error)
	GetUserRedemptions(userID string, page, limit int) ([]*types.RedemptionResponse, int64, error)
}

type AuthService interface {
	SignUp(input types.SignUpInput) (*store.User, error)
	Login(input types.LoginInput) (*types.LoginResponse, error)
	RefreshToken(refreshToken string) (*types.TokenPair, error)
	ChangePassword(userID, currentPassword, newPassword string) error
}

type UserService interface {
	GetProfile(userID string) (*types.UserDTO, error)
	UpdateProfile(userID string, input types.UpdateProfileRequest) error
}

type WalletService interface {
	GetWallet(userID string) (*store.Wallet, error)
	DeductPointsTx(tx *gorm.DB, userID string, points int) error
}

type AdminService interface {
	GetAdminDashboardStats() (*types.DashboardStatsResponse, error)
	GetAllUsers(page, limit int, search, sortBy, sortOrder string) ([]*types.UserDTO, int, error)
	GetAllPurchases(page, limit int, status, dateFrom, dateTo string) ([]*types.PurchaseResponse, int, error)
	GetAllRedemptions(page, limit int, status, dateFrom, dateTo string) ([]*types.RedemptionResponse, int, error)
	UpdateRedemptionStatus(id, status, notes string) error
	ManageUserCredits(userID, action string, amount int) error
	UpdateUserStatus(userID, status string) error
}
