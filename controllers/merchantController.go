package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"vegetablesMarket/repo"
)

type MerchantController struct {
}

func NewMerchantController() *MerchantController {
	return &MerchantController{}
}

func (a *MerchantController) InsertDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	merchantName := r.PostForm["merchantName"][0]
	merchantAddress := r.PostForm["merchantAddress"][0]
	merchant := repo.NewMerchant(0, merchantName, merchantAddress)
	err := merchant.CreateDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchant)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantController) UpdateDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	merchantId, err := strconv.Atoi(r.PostForm["merchantId"][0])
	merchantName := r.PostForm["merchantName"][0]
	merchantAddress := r.PostForm["merchantAddress"][0]
	merchant := repo.NewMerchant(merchantId, merchantName, merchantAddress)
	err = merchant.UpdateDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchant)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantController) DeleteDB(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	merchantId, err := strconv.Atoi(r.PostForm["merchantId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchant := repo.NewMerchant(merchantId, "", "")
	err = merchant.DeleteDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = merchant.DeleteDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchant)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantController) SelectById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	merchantId, err := strconv.Atoi(r.PostForm["merchantId"][0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchant := repo.NewMerchant(merchantId, "", "")
	err = merchant.SelectDBById()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchant)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
