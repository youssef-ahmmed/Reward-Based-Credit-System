package app

import (
	"Start/internal/shared/database"
	"github.com/gin-gonic/gin"
)

func RegisterModules(r *gin.Engine) {
	apiGroup := r.Group("/api")

	db := database.GetDB()

	RegisterAdminModule(apiGroup, db)
	RegisterAuthModule(apiGroup, db)
	RegisterUserModule(apiGroup, db)
	RegisterCategoryModule(apiGroup, db)
	RegisterCreditPackageModule(apiGroup, db)
	RegisterProductModule(apiGroup, db)
	RegisterPurchaseModule(apiGroup, db)
	RegisterRedemptionModule(apiGroup, db)
	RegisterWalletModule(apiGroup, db)
	RegisterAIModule(apiGroup, db)
}
