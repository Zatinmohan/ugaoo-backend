package category_model

type CategoryModel struct {
	Id   string `bson:"_id" json:"id"`
	Name string `json:"name"`
}
