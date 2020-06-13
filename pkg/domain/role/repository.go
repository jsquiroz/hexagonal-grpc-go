package role

import "errors"

var (
	// ErrSQLStatement returned if the query was incorrect
	ErrSQLStatement = errors.New("Could not prepare statement")
	// ErrSQLInsert returned if the row could not insert into db
	ErrSQLInsert = errors.New("Could not insert row")
)

// Repository provides access to role repository
type Repository interface {
	AddRole(Role) error
}
