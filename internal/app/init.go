package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterModules(r *gin.Engine, db *gorm.DB) {
	apiGroup := r.Group("/api")

	RegisterAdminModule(apiGroup, db)
	RegisterAuthModule(apiGroup, db)
	RegisterUserModule(apiGroup, db)
	RegisterCategoryModule(apiGroup, db)
	RegisterCreditPackageModule(apiGroup, db)
	RegisterProductModule(apiGroup, db)
	RegisterPurchaseModule(apiGroup, db)
	RegisterRedemptionModule(apiGroup, db)
	RegisterWalletModule(apiGroup, db)
}
