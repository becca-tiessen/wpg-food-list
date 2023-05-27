package migrations

import (
	"database/sql"

	goose "github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateVotesTable, downCreateVotesTable)
}

func upCreateVotesTable(tx *sql.Tx) error {
	sq := `CREATE TABLE IF NOT EXISTS votes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		restaurant_id INT NOT NULL,
		factor_id INT NOT NULL,
		factor_rating INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP
	)ENGINE=INNODB;`

	_, err := tx.Exec(sq, nil)
	if err != nil {
		return err
	}
	// This code is executed when the migration is applied.
	return nil
}

func downCreateVotesTable(tx *sql.Tx) error {
	sq := "DROP TABLE votes;"

	_, err := tx.Exec(sq)
	if err != nil {
		return err
	}
	// This code is executed when the migration is rolled back.
	return nil
}
