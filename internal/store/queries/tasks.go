package queries

const (
	InsertSingleTaskSQL = `
		INSERT INTO Tasks (name, tag, priority, importance_variance)
		VALUES(?, ?, ?, ?);
	`
	SelectPositionsDataFromTasksSQL = `
		SELECT 
		id, date_created, 
		priority, importance_variance
		FROM Tasks;
	`
	SelectRelavantOrderedListInfoSQL = `
		SELECT T.name, T.tag, T.date_created, T.priority 
		FROM Tasks AS T
		INNER JOIN Positions AS P
			ON T.id = P.task_id
		ORDER BY P.position;
	`
	SelectTaskIdGivenPositionSQL = `
		SELECT T.id
		FROM Tasks T
		RIGHT JOIN Positions P
		ON T.id = P.task_id
		WHERE P.Position = ?;
	`
	DeleteTaskGivenTaskIdSQL = `
		DELETE FROM Tasks
		WHERE id = ?;
	`
)
