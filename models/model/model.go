package model

const (
	// SvHttp indicate http schema
	SvHttp = "http"

	// SvHttps indicate https schema
	SvHttps = "https"

	// DbNone indicate unknown database type
	DbNone = ""

	// DbMySQL indicate MySQL type
	DbMySQL = "mysql"

	// DbPostgreSQL indicate PostgreSQL type
	DbPostgreSQL = "postgres"

	// DbSqlite3 indicate Sqlite3 type
	DbSqlite3 = "sqlite3"

	// EnvConfigFile indicate environment value for config file
	EnvConfigFile = "CHIMUSIC_CONFIG_FILE"

	// EnvCertFile indicate environment value for cert file
	EnvCertFile = "CHIMUSIC_CERT_FILE"

	// EnvKeyFile indicate environment value for key file
	EnvKeyFile = "CHIMUSIC_KEY_FILE"
)
