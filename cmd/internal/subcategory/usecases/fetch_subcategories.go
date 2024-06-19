package subcategory_usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zatinmohan007/ugaoo/cmd/database"
	subcategory_model "github.com/zatinmohan007/ugaoo/cmd/internal/subcategory/models"
	"github.com/zatinmohan007/ugaoo/cmd/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAllSubcategories(w http.ResponseWriter, r *http.Request) {
	var subcategories []subcategory_model.SubcategoryModel
	ctx := context.TODO()

	collection, err := database.GetCollection(os.Getenv("SUBCATEGORY_COLLECTION"))

	if err != nil {
		panic(err)
	}

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		defer cursor.Close(ctx)
		panic(err)
	}

	for cursor.Next(ctx) {
		var subcategory subcategory_model.SubcategoryModel
		err := cursor.Decode(&subcategory)

		if err != nil {
			log.Printf("Error while fetching all subcategories: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			errorResult := response.GetFinalResponse(http.StatusInternalServerError, response.Error(http.StatusInternalServerError), subcategories)
			json.NewEncoder(w).Encode(errorResult)
		}

		subcategories = append(subcategories, subcategory)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := response.GetFinalResponse(http.StatusOK, fmt.Sprintf("%d Subcategories Found", len(subcategories)), subcategories)
	json.NewEncoder(w).Encode(result)

}
