package initialize

import (
	"github.com/wordpress-plus/kit-common/gormx/gconfig"
	"github.com/wordpress-plus/kit-common/gormx/initialize/internal"
	"github.com/wordpress-plus/kit-common/gormx/migration"
	"github.com/wordpress-plus/kit-common/kg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSQL 初始化 Postgresql 数据库
func GormPgSQL(doInit bool) *gorm.DB {
	p := kg.C.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular))
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	if doInit {
		migration.InitializePgsql(db)
	}

	return db
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func GormPgSqlByConfig(p gconfig.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular))
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	return db
}
