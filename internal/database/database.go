package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
	migrations := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,

		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			email TEXT UNIQUE NOT NULL,
			password_hash TEXT,
			name TEXT,
			default_currency TEXT(3) DEFAULT 'USD',
			avatar_url TEXT,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS groups (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			owner_id UUID REFERENCES users(id) ON DELETE CASCADE,
			name TEXT NOT NULL,
			description TEXT,
			avatar_url TEXT,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS memberships (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			role TEXT DEFAULT 'member',
			joined_at TIMESTAMPTZ DEFAULT NOW(),
			UNIQUE(group_id, user_id)
		);`,

		`CREATE TABLE IF NOT EXISTS expenses (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
			created_by UUID REFERENCES users(id) ON DELETE CASCADE,
			amount DECIMAL(18,4) NOT NULL,
			currency TEXT(3) NOT NULL,
			description TEXT,
			date DATE NOT NULL,
			category TEXT,
			receipt_url TEXT,
			metadata JSONB,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS expense_shares (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			expense_id UUID REFERENCES expenses(id) ON DELETE CASCADE,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			share_amount DECIMAL(18,4) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS settlements (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
			payer UUID REFERENCES users(id) ON DELETE CASCADE,
			payee UUID REFERENCES users(id) ON DELETE CASCADE,
			amount DECIMAL(18,4) NOT NULL,
			currency TEXT(3) NOT NULL,
			method TEXT DEFAULT 'cash',
			description TEXT,
			metadata JSONB,
			created_at TIMESTAMPTZ DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS events (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			type TEXT NOT NULL,
			user_id UUID REFERENCES users(id),
			group_id UUID REFERENCES groups(id),
			payload JSONB,
			created_at TIMESTAMPTZ DEFAULT NOW()
		);`,

		// Indexes for better performance
		`CREATE INDEX IF NOT EXISTS idx_memberships_user_id ON memberships(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_memberships_group_id ON memberships(group_id);`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_group_id ON expenses(group_id);`,
		`CREATE INDEX IF NOT EXISTS idx_expenses_created_by ON expenses(created_by);`,
		`CREATE INDEX IF NOT EXISTS idx_expense_shares_expense_id ON expense_shares(expense_id);`,
		`CREATE INDEX IF NOT EXISTS idx_expense_shares_user_id ON expense_shares(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_settlements_payer ON settlements(payer);`,
		`CREATE INDEX IF NOT EXISTS idx_settlements_payee ON settlements(payee);`,
		`CREATE INDEX IF NOT EXISTS idx_events_type ON events(type);`,
		`CREATE INDEX IF NOT EXISTS idx_events_user_id ON events(user_id);`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("failed to run migration: %w", err)
		}
	}

	return nil
}
