package repository_accrual

const (
	queryExistOrderID = `
		SELECT 1 FROM orders WHERE order_number = $1 LIMIT 1;
	`
	queryCreateOrder = `
		INSERT INTO orders (order_number, description, price)
		VALUES ($1, $2, $3)
	`
)