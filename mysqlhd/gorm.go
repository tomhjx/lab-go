package mysqlhd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GORMHandle struct {
	db *gorm.DB
}

func (h *GORMHandle) Query(sql string, param ...any) (any, error) {
	dest := []map[string]any{}
	r := h.db.Raw(sql, param...).Scan(&dest)
	if r.Error != nil {
		return nil, r.Error
	}
	return dest, nil
}

func (h *GORMHandle) Exec(sql string, param ...any) error {
	r := h.db.Exec(sql, param...)
	return r.Error
}

func NewGORMHandle(dsn string) (DBHandler, error) {
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	PrepareStmt: true,
	// })
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &GORMHandle{db: db}, nil
}
