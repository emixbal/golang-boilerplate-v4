package customer

import (
	"context"
	"database/sql"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
)

type mysqlCustomerRepository struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) repository.CustomerRepository {
	return &mysqlCustomerRepository{Conn}
}

// List implements repository.CustomerRepository.
func (m *mysqlCustomerRepository) List(ctx context.Context) (customers []*models.Customer, err error) {

	query := `SELECT id, name, phone, created_at FROM customers`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := new(models.Customer)
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Phone,
			&row.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, row)
	}

	return
}

// Delete implements repository.CustomerRepository.
func (db *mysqlCustomerRepository) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// Detail implements repository.CustomerRepository.
func (*mysqlCustomerRepository) Detail(ctx context.Context, id int) (*models.Customer, error) {
	panic("unimplemented")
}

// Insert implements repository.CustomerRepository.
func (*mysqlCustomerRepository) Insert(ctx context.Context, Customer models.Customer) (*models.Customer, error) {
	panic("unimplemented")
}

// Update implements repository.CustomerRepository.
func (*mysqlCustomerRepository) Update(ctx context.Context, datas map[string]interface{}, id int) error {
	panic("unimplemented")
}
