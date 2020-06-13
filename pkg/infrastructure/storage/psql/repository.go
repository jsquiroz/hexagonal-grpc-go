package psql

import (
	"github.com/jsquiroz/hexagonal-grpc-go/pkg/config"
	"github.com/jsquiroz/hexagonal-grpc-go/pkg/domain/role"
	log "github.com/sirupsen/logrus"
)

const (
	psqlInsert = "INSERT INTO role (name, id_company) VALUES ( $1,  $2 ) RETURNING id"
)

// Storage keeps data in db
type Storage struct{}

// AddRole saves the given Role to the repository
func (m *Storage) AddRole(a role.Role) error {

	db := config.GetConnection()

	stmt, err := db.Prepare(psqlInsert)
	if err != nil {
		log.WithFields(log.Fields{
			"Name":       a.Name,
			"ID Company": a.IDCompany,
		}).Errorf("Error Prepare Statement Insert: %v", err)
		return role.ErrSQLStatement
	}

	defer stmt.Close()
	err = stmt.QueryRow(
		a.Name,
		a.IDCompany,
	).Scan(&a.ID)
	if err != nil {
		log.WithFields(log.Fields{
			"Name":       a.Name,
			"ID Company": a.IDCompany,
		}).Errorf("Error Insert %v:", err)
		return role.ErrSQLInsert
	}

	return nil
}
