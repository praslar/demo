package utils

import (
	config "development-kit/config"
	"errors"
	v6 "github.com/caarlos0/env/v6"
	gorm "github.com/jinzhu/gorm"
)

var conf config.AppConfig

func IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
func LoadEnv() {
	_ = v6.Parse(&conf)
}
func GetEnv() config.AppConfig {
	return conf
}
