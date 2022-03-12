package config

import "os"

type DbConnection struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	SSLMode  string
}

type ESDbConnection struct {
	IndexName string
}

func DbConfig() DbConnection {
	dbConfig := DbConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	return dbConfig
}
func ESDbConfig() ESDbConnection {
	esDbConfig := ESDbConnection{
		IndexName: os.Getenv("ES_INDEX_NAME"),
	}

	return esDbConfig
}

type AppEnvironment struct {
	Port    string
	AppEnv  string
	AppHost string
}

func AppConfig() AppEnvironment {
	appConfig := AppEnvironment{
		AppHost: os.Getenv("APP_HOST"),
		Port:    os.Getenv("PORT"),
		AppEnv:  os.Getenv("APP_ENV"),
	}

	return appConfig
}
