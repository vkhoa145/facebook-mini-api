package transaction

import "database/sql"

type TransactionManager struct {
	DB *sql.DB
}

type SqlTransaction struct {
	Tx *sql.Tx
}

func (tm TransactionManager) NewTransaction() (*SqlTransaction, error) {
	tx, err := tm.DB.Begin()
	if err != nil {
		return nil, err
	}

	return &SqlTransaction{
		Tx: tx,
	}, nil
}
