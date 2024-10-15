package transaction

import "gorm.io/gorm"

type TransactionManager struct {
	Tx *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{
		Tx: db,
	}
}

func (t TransactionManager) Begin() *gorm.DB {
	return t.Tx.Begin()
}

func (t TransactionManager) Commit() *gorm.DB {
	return t.Tx.Commit()
}

func (t TransactionManager) Rollback() *gorm.DB {
	return t.Tx.Rollback()
}
