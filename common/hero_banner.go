package common

type HeroBannerResponse struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Position    uint   `json:"position"`
	Is_active   bool   `json:"is_active"`
}

