package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"vegetablesMarket/repo"
)

type VegetableCategoryController struct {
}

func (a *VegetableCategoryController) InsertDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	categoryId := 0
	categoryName := r.PostForm["categoryName"][0]
	parentId, err := strconv.Atoi(r.PostForm["parentId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetableCategory := repo.NewVegetableCategory(categoryId, categoryName, parentId)
	err = vegetableCategory.CreateDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetableCategory)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *VegetableCategoryController) UpdateDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	categoryId, err := strconv.Atoi(r.PostForm["categoryId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	categoryName := r.PostForm["categoryName"][0]
	parentId, err := strconv.Atoi(r.PostForm["parentId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetableCategory := repo.NewVegetableCategory(categoryId, categoryName, parentId)
	err = vegetableCategory.UpdateDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetableCategory)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *VegetableCategoryController) DeleteDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	categoryId, err := strconv.Atoi(r.PostForm["categoryId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetableCategory := repo.NewVegetableCategory(categoryId, "", 0)
	err = vegetableCategory.DeleteDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetableCategory)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *VegetableCategoryController) SelectById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	categoryId, err := strconv.Atoi(r.PostForm["categoryId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetableCategory := repo.NewVegetableCategory(categoryId, "", 0)
	err = vegetableCategory.SelectDBById()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetableCategory)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
