package repo

import (
	"database/sql"
	"errors"
)

type Vegetable struct {
	VegetableId   int
	VegetableName string
	CategoryId    int
}

func NewVegetable(id int, name string, categoryId int) *Vegetable {
	return &Vegetable{VegetableId: id, VegetableName: name, CategoryId: categoryId}
}

func (a *Vegetable) CreateDB() error {
	var (
		vegetableId int64
	)

	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(" select vegetableId from vegetable where vegetableName = ? and categoryId = ? ",
		a.VegetableName, a.CategoryId)
	err = row.Scan(&vegetableId)
	if err != sql.ErrNoRows {
		return errors.New("蔬菜[" + a.VegetableName + " ]已经存在")
	}
	result, err := tx.Exec(" insert into vegetable(vegetableName,categoryid)values(?,?) ", a.VegetableName, a.CategoryId)
	if err != nil {
		return err
	}
	vegetableId, err = result.LastInsertId()
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	a.VegetableId = int(vegetableId)
	return nil
}

func (a *Vegetable) UpdateDB() error {
	var (
		vegetableId int64
	)
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(" select vegetableid from vegetable where vegetableName = ? and categoryid = ? ", a.VegetableName, a.CategoryId)
	err = row.Scan(&vegetableId)
	if err != sql.ErrNoRows {
		return errors.New("蔬菜[" + a.VegetableName + "]已经存在")
	}
	_, err = tx.Exec(" update vegetable set vegetableName = ? ,categoryId = ? where vegetableid = ?  ",
		a.VegetableName, a.CategoryId, a.VegetableId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *Vegetable) DeleteDB() error {
	tx, err := mySqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(" delete from vegetable where vegetableid = ? ", a.VegetableId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (a *Vegetable) SelectDBById() error {
	var (
		vegetableId   int64
		vegetableName string
		categoryId    int64
	)
	row := mySqlDB.QueryRow(" select vegetableid,vegetablename,categoryid from vegetable where vegetableid = ?  ")
	err := row.Scan(&vegetableId, &vegetableName, &categoryId)
	if err != nil {
		return err
	}
	a.VegetableId = int(vegetableId)
	a.VegetableName = vegetableName
	a.CategoryId = int(categoryId)
	return nil
}
