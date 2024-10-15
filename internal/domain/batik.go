package domain

type Batik struct {
	ID          int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string   `json:"name"`
	JavaName    string   `json:"java_name"`
	Description string   `json:"description"`
	City        string   `json:"city"`
	ProvinceID  int      `json:"province_id"`
	IslandID    int      `json:"island_id"`
	Link_Image  string   `json:"link_image"`
	Province    Province `json:"province" gorm:"foreignKey:ProvinceID"`
	Island      Island   `json:"island" gorm:"foreignKey:IslandID"`
}

type BatikParams struct {
	ID       int    `json:"id"`
	ProvinceID string `json:"province"`
	IslandID   string `json:"island"`
	Page     int    `json:"page" gorm:"-"`
}

type BatikResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	JavaName    string `json:"java_name"`
	Description string `json:"description"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Island      string `json:"island"`
	Link_Image  string `json:"link_image"`
}
