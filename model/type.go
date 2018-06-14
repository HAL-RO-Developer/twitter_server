package model

type Tweet struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Body     string `json:"body" binding:"required"`
	Image    string `json:"image"`
	Favorite string `json:"favorite" binding:"required"`
}

type Tweets []Tweet
