package config

// Config is a struct
type Config struct {
	Email string // TODO: change to User struct ( will needed to send emails... )
	DB    *DBConfig
}

// DBConfig is a struct
type DBConfig struct {
	DatabasePath string
}

// GetConfig is a config settings
func GetConfig() *Config {
	// TODO: Viper configuration plugin
	return &Config{
		Email: "test@test.com",
		DB: &DBConfig{
			DatabasePath: "notes.sqlite3",
		},
	}
}
