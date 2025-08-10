package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Initialize(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func RunMigrations(db *sql.DB) error {
	// Enable pgcrypto extension for UUID generation
	if _, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`); err != nil {
		return fmt.Errorf("failed to create pgcrypto extension: %w", err)
	}

	// Helper function to create a trigger for updating the 'updated_at' column
	createUpdatedAtTrigger := `
        CREATE OR REPLACE FUNCTION update_updated_at_column()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.updated_at = NOW();
            RETURN NEW;
        END;
        $$ language 'plpgsql';
    `
	if _, err := db.Exec(createUpdatedAtTrigger); err != nil {
		return fmt.Errorf("failed to create updated_at trigger function: %w", err)
	}

	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255),
			name VARCHAR(255),
			default_currency VARCHAR(3) DEFAULT 'USD',
			avatar_url VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

		`DROP TRIGGER IF EXISTS update_users_updated_at ON users;`,
		`CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();`,

		`CREATE TABLE IF NOT EXISTS groups (
			id VARCHAR(36) PRIMARY KEY,
			owner_id VARCHAR(36),
			name VARCHAR(255) NOT NULL,
			description TEXT,
			avatar_url VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

		`DROP TRIGGER IF EXISTS update_groups_updated_at ON groups;`,
		`CREATE TRIGGER update_groups_updated_at BEFORE UPDATE ON groups FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();`,

		`CREATE TABLE IF NOT EXISTS memberships (
			id VARCHAR(36) PRIMARY KEY,
			group_id VARCHAR(36),
			user_id VARCHAR(36),
			role VARCHAR(50) DEFAULT 'member',
			joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(group_id, user_id),
			FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS expenses (
			id VARCHAR(36) PRIMARY KEY,
			group_id VARCHAR(36),
			created_by VARCHAR(36),
			amount DECIMAL(18,4) NOT NULL,
			currency VARCHAR(3) NOT NULL,
			description TEXT,
			date DATE NOT NULL,
			category VARCHAR(255),
			receipt_url VARCHAR(255),
			metadata JSON,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
			FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
		);`,

		`DROP TRIGGER IF EXISTS update_expenses_updated_at ON expenses;`,
		`CREATE TRIGGER update_expenses_updated_at BEFORE UPDATE ON expenses FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();`,

		`CREATE TABLE IF NOT EXISTS expense_shares (
			id VARCHAR(36) PRIMARY KEY,
			expense_id VARCHAR(36),
			user_id VARCHAR(36),
			share_amount DECIMAL(18,4) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (expense_id) REFERENCES expenses(id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS settlements (
			id VARCHAR(36) PRIMARY KEY,
			group_id VARCHAR(36),
			payer VARCHAR(36),
			payee VARCHAR(36),
			amount DECIMAL(18,4) NOT NULL,
			currency VARCHAR(3) NOT NULL,
			method VARCHAR(50) DEFAULT 'cash',
			description TEXT,
			metadata JSON,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
			FOREIGN KEY (payer) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (payee) REFERENCES users(id) ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS events (
			id VARCHAR(36) PRIMARY KEY,
			type VARCHAR(255) NOT NULL,
			user_id VARCHAR(36),
			group_id VARCHAR(36),
			payload JSON,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (group_id) REFERENCES groups(id)
		);`,

		`CREATE TABLE IF NOT EXISTS whatsapp_users (
			whatsapp_id VARCHAR(255) PRIMARY KEY,
			user_id VARCHAR(36) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS user_consent (
			user_id VARCHAR(36) PRIMARY KEY,
			has_consented BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("failed to run migration: %w", err)
		}
	}

	return nil
}
