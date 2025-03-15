package model


type UserResponse struct {
	ID       uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username  string `json:"username"`
}


type ProductResponse struct {
	ID          uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Price       uint   `json:"price"`
}
