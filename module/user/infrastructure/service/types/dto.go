package types

type CreateUser struct {
	Email    string
	Password string
	Name     string
}

type CreateUserResult struct {
	SSS2 string
	SSS3 string
}

type UpdateUser struct {
	WalletAddress string
	Name          string
}

type UpdateUserPassword struct {
	WalletAddress   string
	CurrentPassword string
	NewPassword     string
}
