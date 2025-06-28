package user

type SignUpInput struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken  string  `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	User         UserDTO `json:"user"`
}

type UserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=NewPassword"`
}

type UpdateProfileRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Username  *string `json:"username"`
}
