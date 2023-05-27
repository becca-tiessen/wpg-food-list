package migrations

import (
	"database/sql"

	goose "github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateFactorsTable, downCreateFactorsTable)
}

func upCreateFactorsTable(tx *sql.Tx) error {
	sq := `CREATE TABLE IF NOT EXISTS factors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR (255) COLLATE utf8_unicode_ci
	)ENGINE=INNODB;`

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is applied.
	return nil
}

func downCreateFactorsTable(tx *sql.Tx) error {
	sq := "DROP TABLE factors;"

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is rolled back.
	return nil
}
