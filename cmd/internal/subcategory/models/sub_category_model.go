package subcategory_model

type SubcategoryModel struct {
	Name         string   `json:"name"`
	Id           string   `bson:"_id" json:"id"`
	CategoriesId []string `bson:"categoriesId" json:"categoriesId"`
}
