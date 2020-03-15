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

	vegetableController := controllers.NewVegetableController()
	http.HandleFunc("/vegetable/create", vegetableController.InsertDB)
	http.HandleFunc("/vegetable/update", vegetableController.UpdateDB)
	http.HandleFunc("/vegetable/delete", vegetableController.DeleteDB)
	http.HandleFunc("/vegetable/selectById", vegetableController.SelectById)

	merchantController := controllers.NewMerchantController()
	http.HandleFunc("/merchant/create", merchantController.InsertDB)
	http.HandleFunc("/merchant/update", merchantController.UpdateDB)
	http.HandleFunc("/merchant/delete", merchantController.DeleteDB)
	http.HandleFunc("/merchant/selectById", merchantController.SelectById)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
