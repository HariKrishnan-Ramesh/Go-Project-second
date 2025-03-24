package common

type HeroBannerCreationInput struct{
	Title	string	`json:"title" binding:"required"`
	Description	string	`json:"description"`
}

type SetHeroBannerImage struct {
	ImageURL string `json:"imageURL" binding:"required"` 
}

type HeroBannerResponse struct{
	Id  uint  `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Image []HeroBannerImageResponse `json:"images"`
}

type HeroBannerImageResponse struct{
	 ImageURL string `json:"iamge_url"`
}