package services

import (
	"database/sql"
)

// サービス構造体
type MyAppService struct {
	db *sql.DB
}

// コンストラクタ
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
