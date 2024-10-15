package domain

type Tier struct {
	Tier          string `json:"tier" gorm:"primaryKey;type:varchar(100)"`
	TierPhotoLink string `json:"tier_photo_link"`
}
