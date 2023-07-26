package datasource

import (
	"database/sql"
	"fmt"

	"github.com/ka1i/cli/internal/pkg/config"
	"github.com/ka1i/cli/internal/pkg/handlers"
	"github.com/ka1i/cli/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlClient struct {
	client *gorm.DB
}

func (m *mysqlClient) GetClient() *gorm.DB {
	return m.client
}

// Mysql
func (m *mysqlClient) openMysql() {
	options := config.Cfg.Get()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		options.Mysql.User,
		options.Mysql.Passwd,
		options.Mysql.Addr,
		options.Mysql.Port,
		options.Mysql.Database,
		options.Mysql.Options,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// if MaxOpenConns != 0; keep MaxOpenConns > MaxIdleConns
	sqlDB.SetMaxOpenConns(options.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(options.Mysql.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(options.Mysql.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(options.Mysql.ConnMaxLifetime)

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		PrepareStmt: true,
		Logger:      handlers.GormLogger(),
	})

	if err != nil {
		panic(err)
	}

	logger.Printf("Mysql Use %s:%s [%s]\n", options.Mysql.Addr, options.Mysql.Port, "OK")

	m.client = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")
}
