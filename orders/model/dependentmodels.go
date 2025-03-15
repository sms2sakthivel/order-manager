package model

// UserResponse model used for user details.
type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
 
// ProductResponse model used for product details.
type ProductResponse struct {
	ID          uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Price       uint   `json:"price"`
}
