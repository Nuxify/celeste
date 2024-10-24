package repository

import (
	"celeste/infrastructures/database/mysql/types"
)

// IAMQueryRepository handles the iam query repository logic
type IAMQueryRepository struct {
	types.MySQLDBHandlerInterface
}
