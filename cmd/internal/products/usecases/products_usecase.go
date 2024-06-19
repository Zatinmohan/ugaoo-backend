package product_usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zatinmohan007/ugaoo/cmd/database"
	product_model "github.com/zatinmohan007/ugaoo/cmd/internal/products/models"
	"github.com/zatinmohan007/ugaoo/cmd/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []product_model.Product
	collection, err := database.GetCollection(os.Getenv("PRODUCT_COLLECTION"))
	if err != nil {
		panic(err)
	}
	cursor, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		defer cursor.Close(context.TODO())
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var singleProduct product_model.Product
		err := cursor.Decode(&singleProduct)
		if err != nil {
			log.Printf("Error while fetching all proucts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			errorResult := response.GetFinalResponse(http.StatusInternalServerError, response.Error(http.StatusInternalServerError), products)
			json.NewEncoder(w).Encode(errorResult)
		}

		products = append(products, singleProduct)
	}

	w.Header().Add("Content-Type", "application/json")
	result := response.GetFinalResponse(http.StatusOK, fmt.Sprintf("%d Products Found", len(products)), products)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
