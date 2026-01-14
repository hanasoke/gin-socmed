package dto

type PostResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"-"`
	User       User   `gorm:"foreignKey:UserID" json:"user"`
	Tweet      string `json:"tweet"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
	UpdateAt   string `json:"updated_at"`
}

type PostRequest struct {
}
