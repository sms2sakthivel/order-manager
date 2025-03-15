package model

// UserResponse is the response structure returned for a user
type UserResponse struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// ProductResponse is the response structure returned for a product
type ProductResponse struct {
    ID          uint   `json:"product_id"`
    Name        string `json:"product_name"`
    Brand       string `json:"brand"`
    Category    string `json:"category"`
    SubCategory string `json:"sub_category"`
    Price       uint   `json:"price"`
}
