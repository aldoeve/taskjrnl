package queries

const (
	SelectPageIDsFromTaskIdSQL = `
		SELECT page_id
		FROM Jrnls
		WHERE task_id = ?;
	`
	DeleteJrnlsFromTaskIdSQL = `
		DELETE FROM Jrnls
		WHERE task_id = ?;
	`
	CreateJrnlSQL = `
		INSERT INTO Jrnls (task_id, page_id) VALUES (?,?);
	`
)
