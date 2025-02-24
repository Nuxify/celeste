package types

type CreateUser struct {
	Email string
	Name  string
}

type UpdateUser struct {
	WalletAddress string
	Name          string
}

type UpdateUserPassword struct {
	WalletAddress string
	Password      string
}
