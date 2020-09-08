package extend

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)
type Orm struct {
	config map[string]interface{}
	*sql.DB
	Database
}


var O *Orm
var OrmGroup  map[string]*sql.DB
var err error

func init() {
	O = &Orm{
		config: map[string]interface{}{},
	}
	O.DbGroup()
}

func (o *Orm) Init(m map[string]interface{}) {
	o.config = m
}

type DbConfig struct {
	DriverName string
	DataSourceName string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime time.Duration
}

func  (o *Orm) DbGroup(){
	//需要知道o.config 数据类型 从哪里取

	//初始化所有数据库组连接
	for dbname,group := range o.config{

		dbConfig := &DbConfig{
			DriverName:     group["driverName"],
			DataSourceName: group["dataSourceName"],
			MaxOpenConn:    group["maxOpenConn"],
			MaxIdleConn:    group["maxIdleConn"],
			MaxLifetime:    time.Hour,
		}

		OrmGroup[dbname] = registerOrm(dbConfig)
	}
}

func registerOrm(conn DbConfig) *sql.DB {

	sqlDB, err := gorm.Open(conn.DriverName, conn.DataSourceName)
	if err != nil {
		log.Fatalf("DB connect faild err: %v", err)
	}
	defer sqlDB.Close()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.DB().SetMaxIdleConns(conn.MaxIdleConn)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.DB().SetMaxOpenConns(conn.MaxOpenConn)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.DB().SetConnMaxLifetime(conn.MaxLifetime)
	return sqlDB
}





