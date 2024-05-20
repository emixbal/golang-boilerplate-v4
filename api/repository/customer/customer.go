package customer

import (
	"context"
	"database/sql"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"log"
	"strconv"
)

type mysqlCustomerRepository struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) repository.CustomerRepository {
	return &mysqlCustomerRepository{Conn}
}

// List implements repository.CustomerRepository.
func (db *mysqlCustomerRepository) List(ctx context.Context) (customers []*models.Customer, err error) {

	query := `SELECT id, name, phone, created_at FROM customers`

	rows, err := db.Conn.QueryContext(ctx, query)
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
func (db *mysqlCustomerRepository) Delete(ctx context.Context, id int) (err error) {
	ids := strconv.Itoa(id)
	query := `DELETE FROM customers WHERE id = ?`

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, query, ids)
	if err != nil {
		return err
	}
	return
}

// Detail implements repository.CustomerRepository.
func (db *mysqlCustomerRepository) Detail(ctx context.Context, id int) (customer *models.Customer, err error) {
	customer = &models.Customer{}

	query := `SELECT id, name, phone, created_at FROM customers WHERE id = ?`
	err = db.Conn.QueryRowContext(ctx, query, id).Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.CreatedAt)

	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return customer, nil
}

// Insert implements repository.CustomerRepository.
func (*mysqlCustomerRepository) Insert(ctx context.Context, Customer models.Customer) (*models.Customer, error) {
	panic("unimplemented")
}

// Update implements repository.CustomerRepository.
func (*mysqlCustomerRepository) Update(ctx context.Context, datas map[string]interface{}, id int) error {
	panic("unimplemented")
}
