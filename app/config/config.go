package config

// Config is a struct
type Config struct {
	Email         string // TODO: change to User struct ( will needed to send emails... )
	DB            *DBConfig
	TimeFormat    string
	ContentLength uint
}

// DBConfig is a struct
type DBConfig struct {
	DatabasePath string
}

// GetConfig is a config settings
func GetConfig() *Config {
	// TODO: Viper configuration plugin
	return &Config{
		DB: &DBConfig{
			DatabasePath: "notes.sqlite3",
		},
		Email:         "test@test.com",
		TimeFormat:    "2006-01-02",
		ContentLength: 35,
	}
}
