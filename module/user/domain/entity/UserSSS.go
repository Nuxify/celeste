package entity

import "time"

// UserSSS holds the user sss entity fields
type UserSSS struct {
	UserWalletAddress string    `db:"user_wallet_address"`
	SSS3              string    `db:"sss_3"`
	CreatedAt         time.Time `db:"created_at"`
}

// GetModelName returns the model name of user sss entity that can be used for naming schemas
func (u *UserSSS) GetModelName() string {
	return "user_sss"
}
