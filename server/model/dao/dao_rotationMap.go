package dao

type RotationMap struct {
	BID int    `json:"bid" gorm:"column:bid"` // id
	URL string `json:"url" gorm:"column:url"`
}

func (r *RotationMap) TableName() string {
	return "rotationMap"
}
