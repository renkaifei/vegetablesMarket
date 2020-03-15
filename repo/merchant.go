package repo

import (
	"database/sql"
	"errors"
)

type Merchant struct {
	MerchantId      int
	MerchantName    string
	MerchantAddress string
}

func NewMerchant(id int, name string, address string) *Merchant {
	return &Merchant{MerchantId: id, MerchantName: name, MerchantAddress: address}
}

func (a *Merchant) CreateDB() error {
	var (
		merchantId int64
	)

	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(" select MerchantId from Merchant where MerchantName = ? ", a.MerchantName)
	err = row.Scan(&merchantId)
	if err != sql.ErrNoRows {
		return errors.New("供应商[" + a.MerchantName + "]已经存在")
	}
	result, err := tx.Exec(" insert into merchant(merchantName,merchantaddress)values(?,?)", a.MerchantName, a.MerchantAddress)
	if err != nil {
		return err
	}
	merchantId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	a.MerchantId = int(merchantId)
	return nil
}
