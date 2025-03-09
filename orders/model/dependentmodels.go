package model

// UserResponse represents the user information returned by the API
type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

// ProductResponse represents the product information returned by the API
type ProductResponse struct {
	ID          uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Price       uint   `json:"price"`
}
