package entities

type PayloadAssignProfil struct {
	IDUser   int `json:"idUser" gorm:"column:idUser"`
	IDProfil int `json:"idProfil" gorm:"column:idProfil"`
}

func (PayloadAssignProfil) TableName() string {
	return "userProfil"
}
