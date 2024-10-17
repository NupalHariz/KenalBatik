package domain

type Province struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	IslandID int `json:"island_id"`
	Name string `json:"name"`
}

type ProvinceParams struct {
	ID int
}
