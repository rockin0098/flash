package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/rockin0098/flash/base/logger"
)

type TSQLLogger struct{}

func (slog TSQLLogger) Print(values ...interface{}) {
	vals := gorm.LogFormatter(values...)
	Log.SqlDebug(vals...)
}

type DataSourceManager struct {
	configs  []DataSource
	masterDB *gorm.DB
	slaveDB  *gorm.DB
	models   []interface{}
}

var datasourceManager = &DataSourceManager{}

type DataSource struct {
	URL         string
	IdleSize    int
	MaxSize     int
	MaxLifeTime int64
	SqlDebug    int
}

func DataSourceInit(configs []DataSource, models ...interface{}) {
	datasourceManager.configs = configs
	datasourceManager.models = models

	Log.Infof("datasourceManager = %+v", datasourceManager)

	if len(configs) == 0 {
		Log.Warnf("db config not found...")
		return
	}

	if len(configs) == 1 {
		datasourceManager.masterDB = datasourceManager.openDBConn(configs[0])
		datasourceManager.slaveDB = datasourceManager.masterDB
	} else {
		datasourceManager.masterDB = datasourceManager.openDBConn(configs[0])
		datasourceManager.slaveDB = datasourceManager.openDBConn(configs[1])
	}
}

func (d *DataSourceManager) openDBConn(ds DataSource) *gorm.DB {
	db, err := gorm.Open("mysql", ds.URL)
	if err != nil {
		Log.Error("connect to mysql failed, err = %v", err)
		return nil
	}

	db.DB().SetMaxIdleConns(ds.IdleSize)
	db.DB().SetMaxOpenConns(ds.MaxSize)
	db.DB().SetConnMaxLifetime(time.Duration(ds.MaxLifeTime) * time.Second)

	// 设置字符编码
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	db.SingularTable(true)
	if ds.SqlDebug == 1 {
		db.LogMode(true)
		db.SetLogger(TSQLLogger{})
	}

	for _, m := range d.models {
		if !db.HasTable(m) {
			err := db.CreateTable(m).Error
			if err != nil {
				Log.Error(err)
			}
		}
	}
	db.AutoMigrate(d.models...)

	return db
}

func (d *DataSourceManager) Master() *gorm.DB {
	return d.masterDB
}

func (d *DataSourceManager) Slave() *gorm.DB {
	return d.slaveDB
}

func DataSourceInstance() *DataSourceManager {
	return datasourceManager
}
