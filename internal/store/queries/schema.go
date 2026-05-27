package queries

const (
	SchemaSQL = `
		CREATE TABLE IF NOT EXISTS Tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			tag TEXT,
			date_created TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
			priority CHAR(1) NOT NULL CHECK(priority IN('L', 'M','H')),
			importance_variance INTEGER NOT NULL
		);
	
		CREATE TABLE IF NOT EXISTS Positions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task_id INTEGER NOT NULL,
			position INTEGER NOT NULL,
			FOREIGN KEY (task_id) REFERENCES Tasks(id)
		)
	`
)
