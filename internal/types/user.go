package types

type UserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type UpdateProfileRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Username  *string `json:"username"`
}

type DashboardStatsResponse struct {
	TotalUsers    int    `json:"totalUsers"`
	TotalOrders   int    `json:"totalOrders"`
	CreditsIssued string `json:"creditsIssued"`
	PointsEarned  string `json:"pointsEarned"`
}

type UpdateRedemptionStatusRequest struct {
	Status string `json:"status" binding:"required"`
	Notes  string `json:"notes"`
}

type ManageCreditsRequest struct {
	Action string `json:"action" binding:"required"` // "add" or "subtract"
	Amount int    `json:"amount" binding:"required,gt=0"`
}

type ModerateUserRequest struct {
	Status string `json:"status" binding:"required"` // active, suspended, banned
	Reason string `json:"reason"`
}
