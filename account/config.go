package main

// ConfigModel configuration model.
type ConfigModel struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	SvcPort    string
}

// Config service configuration object
var (
	Config = &ConfigModel{
		DBName:     "simpos_account",
		DBUser:     "simpos_account",
		DBPassword: "p@ssw0rd",
		DBHost:     "localhost",
		DBPort:     "3306",
		// SvcPort This service's port.
		SvcPort: ":6070"}
)
