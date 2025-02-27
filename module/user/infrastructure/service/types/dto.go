package types

type CreateUser struct {
	Email    string
	Password string
	Name     string
}

type CreateUserResult struct {
	WalletAddress string
	SSS2          string
	SSS3          string
}

type UpdateUser struct {
	WalletAddress string
	Name          string
}

type UpdateUserPassword struct {
	WalletAddress string
	Password      string
}
