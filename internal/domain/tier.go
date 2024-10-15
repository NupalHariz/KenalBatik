package domain

type Tier struct {
	Tier          string `json:"tier" gorm:"primaryKey"`
	TierPhotoLink string `json:"tier_photo_link"`
}
