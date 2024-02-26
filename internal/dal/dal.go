package dal

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/config"
	"gitlabee.chehejia.com/k8s/liks-gitops/internal/dal/query"
)

func NewQuery(cfg *config.Config) (*query.Query, error) {
	var gormCfg gorm.Config
	if cfg.Debug {
		gormCfg.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Microsecond, // print every SQL statement
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		})
	}

	if cfg.DB == nil {
		return nil, fmt.Errorf("db: not found configuration")
	}

	gdb, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gormCfg)
	if err != nil {
		return nil, fmt.Errorf("initDb: %w", err)
	}

	if err := gdb.Use(tracing.NewPlugin()); err != nil {
		return nil, fmt.Errorf("use tracing plugin: %w", err)
	}

	if cfg.Debug {
		gdb = gdb.Debug()
	}

	db, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("get raw db: %v", err)
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return query.Use(gdb), nil
}
