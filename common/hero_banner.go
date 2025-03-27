package common

type HeroBannerResponse struct {
	Id          uint                      `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Images      []HeroBannerImageResponse `json:"images"`
	Position    uint                      `json:"position"`
	Is_active   bool                      `json:"is_active"`
}

type HeroBannerImageResponse struct {
	ImageURL string `json:"image_url"`
}
