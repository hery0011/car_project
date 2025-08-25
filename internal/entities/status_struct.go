package entities

type Status struct {
	IdStatus  int    `json:"id_status" gorm:"column:id_status"`
	NomStatus string `json:"nom_status" gorm:"nom_status"`
}

func (Status) TableName() string {
	return "status"
}
