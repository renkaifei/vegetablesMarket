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

func (a *Merchant) UpdateDB() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(" update merchant set MerchantName = ? ,merchantAddress = ? where merchantId = ? ", a.MerchantName, a.MerchantAddress, a.MerchantId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return nil
	}
	return nil
}

func (a *Merchant) DeleteDB() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(" delete from merchant where merchantId = ? ", a.MerchantId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *Merchant) SelectDBById() error {
	var (
		merchantId      int64
		merchantName    string
		merchantAddress string
	)
	row := mySqlDB.QueryRow(" select MerchantId,MerchantName,MerchantAddress from merchant where MerchantId = ? ", a.MerchantId)
	err := row.Scan(&merchantId, &merchantName, &merchantAddress)
	if err == sql.ErrNoRows {
		return errors.New("没有找到指定商家")
	}
	if err != nil {
		return err
	}
	a.MerchantId = int(merchantId)
	a.MerchantName = merchantName
	a.MerchantAddress = merchantAddress
	return nil
}
