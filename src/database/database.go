package database

import (
	"fmt"
	"log"
	"scratch/src/database/schemas"
	"scratch/src/config"

	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
	*log.Logger
}

func NewDatabase(conf config.PostgresConfig) (Database, error) {
	var db Database
	ds := fmt.Sprintf("user=%s password=%s host=%s port=%d sslmode=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.SSL)
	fmt.Println(ds)
	db.DB = NewEngine("postgres", ds, conf.Database) //创建数据库连接
	db.DB.Debug()                                    // gorm 日志模式
	err := db.initDb()
	if err != nil {
		return Database{}, err
	}

	return db, nil
}
func (d *Database) initDb() error {
	d.initTable()
	return nil
}

func (d *Database) initTable() error {
	//初始化表结构
	d.DB.AutoMigrate(
		&schemas.SensorType{}, &schemas.Sensors{})

	return nil
}
func (d *Database) Close() {
	// d.DB.Close()
}
