package queries

const (
	DeletePositionsRowsSQL = `
		DELETE FROM Positions;
	`
	InsertSinglePositionRowSQL = `
		INSERT INTO Positions (task_id, position)
		VALUES(?, ?); 
	`
	DeletePositionRowGivenTaskIdSQL = `
		DELETE FROM Positions
		Where task_id = ?;
	`
	SelectTaskIdGivenPositionSQL = `
		SELECT task_id
		FROM Positions
		WHERE position = ?;
	`
)
