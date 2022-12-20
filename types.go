package goravel

import "database/sql"

type initPath struct {
	rootPath    string
	folderNames []string
}

type cookieConfig struct {
	name     string
	lifeTime string
	persist  string
	secure   string
	domain   string
}

type databaseConfig struct {
	dsn string
	database string
}

type Database struct {
	DatabaseType string
	Pool *sql.DB
}