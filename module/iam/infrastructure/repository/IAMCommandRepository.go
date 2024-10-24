package repository

import (
	"celeste/infrastructures/database/mysql/types"
)

// IAMCommandRepository handles the iam command repository logic
type IAMCommandRepository struct {
	types.MySQLDBHandlerInterface
}
