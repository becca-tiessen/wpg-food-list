package migrations

import (
	"database/sql"

	goose "github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateRestaurantTable, downCreateRestaurantTable)
}

func upCreateRestaurantTable(tx *sql.Tx) error {
	sq := `CREATE TABLE IF NOT EXISTS restaurants (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) COLLATE utf8_unicode_ci,
		description TEXT COLLATE utf8_unicode_ci,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
		deleted_at DATETIME DEFAULT NULL
	)ENGINE=INNODB;`

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is applied.
	return nil
}

func downCreateRestaurantTable(tx *sql.Tx) error {
	sq := "DROP TABLE restaurants;"

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is rolled back.
	return nil
}
