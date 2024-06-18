package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	category_usecase "github.com/zatinmohan007/ugaoo/cmd/internal/category/usecase"
	product_usecase "github.com/zatinmohan007/ugaoo/cmd/internal/products/usecases"
	subcategory_usecase "github.com/zatinmohan007/ugaoo/cmd/internal/subcategory/usecases"
)

func SetupRoutes() *mux.Router {
	log.Println("Setting up Routes")
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(notFound)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ugaoo Backend Server")
		json.NewEncoder(w).Encode("Welcome to Ugaoo Backend Server")
	})

	router.HandleFunc("/fetchAllProducts", product_usecase.GetAllProducts).Methods("GET")
	router.HandleFunc("/fetchAllCategories", category_usecase.FetchAllCategory).Methods("GET")
	router.HandleFunc("/fetchAllSubcategories", subcategory_usecase.FetchAllSubcategories).Methods("GET")

	log.Println("Routes set up")
	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	result := map[string]interface{}{"statusCode": http.StatusNotFound, "message": "Not Found"}
	json.NewEncoder(w).Encode(result)
}
