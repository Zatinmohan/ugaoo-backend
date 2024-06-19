package category_usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zatinmohan007/ugaoo/cmd/database"
	category_model "github.com/zatinmohan007/ugaoo/cmd/internal/category/models"
	"github.com/zatinmohan007/ugaoo/cmd/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAllCategory(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	var categories []category_model.CategoryModel

	collection, err := database.GetCollection(os.Getenv("CATEGORY_COLLECTION"))

	if err != nil {
		panic(err)
	}

	cursor, err := collection.Find(ctx, bson.D{{}})

	if err != nil {
		defer cursor.Close(ctx)
		panic(err)
	}

	for cursor.Next(ctx) {
		var category category_model.CategoryModel
		err := cursor.Decode(&category)

		if err != nil {
			log.Printf("Error while fetching all proucts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			errorResult := response.GetFinalResponse(http.StatusInternalServerError, response.Error(http.StatusInternalServerError), categories)
			json.NewEncoder(w).Encode(errorResult)
		}

		categories = append(categories, category)
	}

	result := response.GetFinalResponse(http.StatusOK, fmt.Sprintf("%d Categories Found", len(categories)), categories)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
