package config

import (
	"fmt"
	"os"
)

func NewMysqlConfigFromEnv(prefix string) (*DB, error) {
	conf := &DB{}
	var err error

	conf.Host, err = GetConfigValue(prefix + "MYSQL_HOST")
	if err == nil {
		conf.Port, err = GetConfigValue(prefix + "MYSQL_PORT")
	}
	if err == nil {
		conf.User, err = GetConfigValue(prefix + "MYSQL_USER")
	}
	if err == nil {
		conf.DBName, err = GetConfigValue(prefix + "MYSQL_DATABASE_NAME")
	}
	if err == nil {
		conf.DBPass, err = GetConfigValue(prefix + "MYSQL_PASSWORD")
	}

	if err != nil {
		return nil, err
	}

	return conf, nil
}

type DB struct {
	Host   string
	Port   string
	User   string
	DBName string
	DBPass string
}

func (d *DB) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", d.User, d.DBPass, d.Host, d.Port, d.DBName)
}

func GetConfigValue(key string) (string, error) {
	value := os.Getenv(key)
	return value, nil
}
