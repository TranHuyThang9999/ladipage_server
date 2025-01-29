package adapters

import (
	"fmt"
	"ladipage_server/common/configs"
	"ladipage_server/core/domain"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pgsql struct {
	db *gorm.DB
}

func NewPgsql() *Pgsql {
	return &Pgsql{}
}

func (p *Pgsql) Connect() error {
	dataSource := configs.Get().DataSource

	dataSourceWithParams := fmt.Sprintf("%s connect_timeout=10 statement_timeout=10000 idle_in_transaction_session_timeout=10000",
		dataSource)

	config := &gorm.Config{
		PrepareStmt: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dataSourceWithParams), config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	if err := db.AutoMigrate(&domain.Users{},
		&domain.FileDescriptors{},
		&domain.VehicleCategory{},
	); err != nil {
		return fmt.Errorf("failed to auto migrate schemas: %v", err)
	}

	p.db = db
	return nil
}

func (p *Pgsql) DB() *gorm.DB {
	return p.db
}
