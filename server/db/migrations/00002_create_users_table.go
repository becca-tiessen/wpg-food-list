package migrations

import (
	"database/sql"

	goose "github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(tx *sql.Tx) error {
	sq := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(255) COLLATE utf8_unicode_ci NOT NULL,
		last_name VARCHAR (255) COLLATE utf8_unicode_ci NOT NULL,
		email VARCHAR (255) COLLATE utf8_unicode_ci NOT NULL UNIQUE,
		auth_id VARCHAR (255) NOT NULL UNIQUE, 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	)ENGINE=INNODB;`

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is applied.
	return nil
}

func downCreateUsersTable(tx *sql.Tx) error {
	sq := "DROP TABLE users;"

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is rolled back.
	return nil
}
