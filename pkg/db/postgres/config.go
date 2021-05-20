package postgres

import (
	config "development-kit/config"
	v6 "github.com/caarlos0/env/v6"
	gorm "github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

var Muxtex *sync.RWMutex

const DefaultConnName = "default"

func GetDBInfo() (dbInfo DBInfo, err error) {
	conf := config.AppConfig{}
	_ = v6.Parse(&conf)
	dbInfo = DBInfo{Host: conf.DBHost, Port: conf.DBPort, Name: conf.DBName, User: conf.DBUser, Pass: conf.DBPass, SearchPath: conf.DBSchema}
	return
}
func GetDatabase(aliasName string) (db *gorm.DB, err error) {
	var customerSchema string
	err = nil
	_, errConv := strconv.Atoi(aliasName)
	if errConv == nil {
		customerSchema = aliasName
	} else {
		customerSchema = "default"
	}
	db, err = GetDB(customerSchema)
	if err != nil {
		var dbInfo DBInfo
		dbInfo, err = GetDBInfo()
		if err == nil {
			err = RegisterDataBase(customerSchema, "postgres", CreateDBConnectionString(dbInfo))
			if err == nil {
				db, err = GetDB(customerSchema)
			}
		}
	}
	return
}
