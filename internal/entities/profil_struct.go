package entities

type Profil struct {
	IDProfil    int    `json:"idProfil" gorm:"column:idProfil;primaryKey;autoIncrement"`
	NomProfil   string `json:"nomProfil" gorm:"column:nomProfil"`
	Description string `json:"description" gorm:"column:description"`
}

func (Profil) TableName() string {
	return "profil"
}
