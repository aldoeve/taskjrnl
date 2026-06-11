package queries

const (
	InsertSingleTaskSQL = `
		INSERT INTO Tasks (name, tag, priority, importance_variance)
		VALUES(?, ?, ?, ?);
	`
	SelectPositionsDataFromTasksSQL = `
		SELECT 
			id, 
			date_created, 
			priority, 
			importance_variance
		FROM Tasks;
	`
	SelectRelavantOrderedListInfoSQL = `
		SELECT 
			T.name, 
			T.tag, 
			T.date_created, 
			T.priority 
		FROM Tasks AS T
		INNER JOIN Positions AS P
			ON T.id = P.task_id
		ORDER BY P.position;
	`
	DeleteTaskGivenTaskIdSQL = `
		DELETE FROM Tasks
		WHERE id = ?;
	`
	SelectTaskInfoGivenTaskIdSQL = `
		SELECT 
			name, 
			tag, 
			date_created, 
			priority, 
			importance_variance
		FROM Tasks
		Where id = ?;
	`
	UpdateTaskAllColumnsSQL = `
		UPDATE Tasks
		SET 
			name = ?,
			tag  = ?,
			priority = ?,
			importance_variance = importance_variance
		WHERE id = ?;
	`
	UpdateTaskWeightSQL = `
		UPDATE Tasks
		SET
			importance_variance = importance_variance + ?
		WHERE id = ?;
	`
)
