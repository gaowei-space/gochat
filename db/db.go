/**
 * Created by lock
 * Date: 2019-09-22
 * Time: 22:37
 */
package db

import (
	"gochat/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var dbMap = map[string]*gorm.DB{}
var syncLock sync.Mutex
var dbName = "gochat"

func init() {
	initMySQL(dbName, config.Conf.Common.CommonMysql.Dsn)
}

func initMySQL(dbName string, dsn string) {
	var e error
	syncLock.Lock()
	dbMap[dbName], e = gorm.Open("mysql", dsn)
	if config.GetMode() == "dev" {
		dbMap[dbName].LogMode(true)
	}
	syncLock.Unlock()
	if e != nil {
		logrus.Error("connect db fail:%s", e.Error())
	}
}

func GetDb(dbName string) (db *gorm.DB) {
	if db, ok := dbMap[dbName]; ok {
		return db
	} else {
		return nil
	}
}

type DbGoChat struct {
}

func (*DbGoChat) GetDbName() string {
	return dbName
}
