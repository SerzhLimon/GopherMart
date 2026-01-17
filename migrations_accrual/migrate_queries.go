package migrations_accrual

const (
	createEnum = `
		CREATE TYPE discount_type AS ENUM ('PERCENT', 'FIXED');
	`

	createTableCatalog = `
		CREATE TABLE IF NOT EXISTS catalog
		(
			id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			description VARCHAR(255) NOT NULL UNIQUE,
			reward INT NOT NULL,
			reward_type discount_type NOT NULL,
			CONSTRAINT reward_percent_check 
				CHECK (
					(reward_type = 'PERCENT' AND reward <= 100 AND reward > 0) OR
					(reward_type = 'FIXED' AND reward > 0)
				)
		);
	`
	createTableOrders = `
		CREATE TABLE IF NOT EXISTS orders
		(
			id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			description VARCHAR(255) NOT NULL UNIQUE,
			price NUMERIC(10, 2),
			order_number VARCHAR(255) NOT NULL UNIQUE
		);
	`
	dropTableCatalog = `
		DROP TABLE IF EXISTS catalog;
	`
	dropTableOrders = `
		DROP TABLE IF EXISTS orders;
	`

	dropEnum = `
		DROP TYPE IF EXISTS discount_type;
	`
)