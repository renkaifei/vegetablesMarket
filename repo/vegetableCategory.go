package repo

import (
	"database/sql"
	"errors"
)

type VegetableCategory struct {
	CategoryId   int
	CategoryName string
	ParentId     int
}

func NewVegetableCategory(id int, name string, parentid int) *VegetableCategory {
	return &VegetableCategory{CategoryId: id, CategoryName: name, ParentId: parentid}
}

func (a *VegetableCategory) CreateDB() error {
	var (
		categoryId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(" select categoryId from VegetableCategory where categoryName = ? and parentId = ? ", a.CategoryName, a.ParentId)
	err = row.Scan(&categoryId)
	if err != sql.ErrNoRows {
		return errors.New("蔬菜[" + a.CategoryName + "]类型已经存在！")
	}

	result, err := tx.Exec(" insert into vegetableCategory(categoryName,parentId)values(?,?) ", a.CategoryName, a.ParentId)
	if err != nil {
		return err
	}
	categoryId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	a.CategoryId = int(categoryId)
	return nil
}

func (a *VegetableCategory) UpdateDB() error {
	var (
		categoryId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(" select categoryId from vegetableCategory where CategoryName = ? and ParentId = ?  ", a.CategoryName, a.ParentId)
	err = row.Scan(&categoryId)
	if err != sql.ErrNoRows {
		return errors.New("蔬菜类型[" + a.CategoryName + "]已经存在")
	}
	_, err = tx.Exec(" update vegetableCategory set CategoryName = ? ,ParentId = ? where CategoryId = ?  ", a.CategoryName, a.ParentId, a.CategoryId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil

}

func (a *VegetableCategory) DeleteDB() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(" delete from VegetableCategory where categoryId = ? ", a.CategoryId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil

}

func (a *VegetableCategory) SelectDBById() error {
	var (
		categoryId   int64
		categoryName string
		parentId     int64
	)
	row := mySqlDB.QueryRow(" select CategoryId,CategoryName,ParentId from VegetableCategory where CategoryId = ? ", a.CategoryId)
	err := row.Scan(&categoryId, &categoryName, &parentId)
	if err != nil {
		return err
	}
	a.CategoryId = int(categoryId)
	a.CategoryName = categoryName
	a.ParentId = int(parentId)
	return nil
}
