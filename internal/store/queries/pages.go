package queries

const (
	DeletePageFromPageIdSQL = `
		DELETE FROM Pages
		WHERE id = ?;
	`
	CreatePageSQL = `
		INSERT INTO Pages (note) 
		VALUES (?);
	`
	FetchJrnlPagesFromTaskIdSQL = `
		SELECT 
			P.note, 
			P.date_created
		FROM Pages AS P
		INNER JOIN Jrnls AS J
			ON J.page_id = P.id
		WHERE J.task_id = ?
		ORDER BY P.date_created ASC; 
	`
)
