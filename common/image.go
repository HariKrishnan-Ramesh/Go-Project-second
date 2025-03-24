package common

import (

)

type SetFeaturedImage struct {
	ImageURL string `json:"imageURL" binding:"required"`
}
