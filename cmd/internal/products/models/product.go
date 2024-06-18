package product_model

type Product struct {
	Id            string   `bson:"_id" json:"id"`
	Name          string   `json:"name"`
	CatetoryId    []string `bson:"categoryId" json:"categoryId"`
	SubcategoryId []string `bson:"subcategoryId" json:"subcategoryId"`
	Price         float32  `json:"price"`
	DiscountPrice float32  `json:"discountPrice"`
	Quantity      int8     `json:"quantity"`
	ImageUrl      string   `json:"imageUrl"`
}
