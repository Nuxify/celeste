package types

type CreateUser struct {
	WalletAddress string
	Email         string
	Password      string
	SSS1          string
	Name          string
	SSS3          string
}

type DeactivateUser struct {
	WalletAddress string
	Email         string
	Password      string
	SSS1          string
	Name          string
}

type UpdateUser struct {
	WalletAddress string
	Name          string
}

type UpdateUserPassword struct {
	WalletAddress string
	Password      string
}
