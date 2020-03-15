package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"vegetablesMarket/repo"
)

type VegetableController struct {
}

func NewVegetableController() *VegetableController {
	return &VegetableController{}
}

func (a *VegetableController) InsertDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vegetableName := r.PostForm["vegetableName"][0]
	categoryId, err := strconv.Atoi(r.PostForm["categoryId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetable := repo.NewVegetable(0, vegetableName, categoryId)
	err = vegetable.CreateDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetable)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *VegetableController) UpdateDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vegetableId, err := strconv.Atoi(r.PostForm["vegetableId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetableName := r.PostForm["vegetableName"][0]
	categoryId, err := strconv.Atoi(r.PostForm["categoryId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetable := repo.NewVegetable(vegetableId, vegetableName, categoryId)
	err = vegetable.UpdateDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetable)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *VegetableController) DeleteDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vegetableId, err := strconv.Atoi(r.PostForm["vegetableId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetable := repo.NewVegetable(vegetableId, "", 0)
	err = vegetable.DeleteDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetable)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *VegetableController) SelectById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vegetableId, err := strconv.Atoi(r.PostForm["vegetableId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vegetable := repo.NewVegetable(vegetableId, "", 0)
	err = vegetable.SelectDBById()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(vegetable)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
