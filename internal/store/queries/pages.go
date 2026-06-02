package queries

const (
	DeletePageFromPageIdSQL = `
		DELETE FROM Pages
		WHERE id = ?;
	`

	CreatePageSQL = `
		INSERT INTO Pages (note) VALUES (?);
	`
)
