package database

type City struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	CityName string `gorm:"column:cityname"`
}

type Tabler interface {
	TableName() string
}

func (City) TableName() string {
	return "city"
}
