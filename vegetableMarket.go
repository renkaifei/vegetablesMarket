package main

import (
	"log"
	"net/http"
	"vegetablesMarket/controllers"
)

func main() {
	vegetableCategoryController := &controllers.VegetableCategoryController{}
	http.HandleFunc("/vegetableCategory/create", vegetableCategoryController.InsertDB)
	http.HandleFunc("/vegetableCategory/update", vegetableCategoryController.UpdateDB)
	http.HandleFunc("/vegetableCategory/delete", vegetableCategoryController.DeleteDB)
	http.HandleFunc("/vegetableCategory/selectById", vegetableCategoryController.SelectById)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
