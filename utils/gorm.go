package utils

import (
	"sync"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var (
	gormDB   *gorm.DB
	gormOnce sync.Once
)

// GetDB ...
func GetDB() *gorm.DB {
	return gormDB
}

// InitGORM ...
func InitGORM(dbPath string) (*gorm.DB, error) {
	if gormDB != nil {
		return gormDB, nil
	}
	var err error
	gormOnce.Do(func() {
		// github.com/mattn/go-sqlite3
		gormDB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
