package common

type CategoryResponse struct {
	Id            uint     `json:"id"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt"`
	CategoryName  string   `json:"category_name"`
	URLKey        string   `json:"url_key"`
	Description   string   `json:"description"`
	BannerImage   []string `json:"banner_image"`
	CategoryImage []string `json:"category_image"`
}
