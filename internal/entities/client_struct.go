package entities

type Client struct {
	Client_id int    `json:"client_id" gorm:"client_id"`
	Nom       string `json:"nom" gomr:"nom"`
	Prenom    string `json:"prenom" gorm:"prenom"`
	Email     string `json:"email" gorm:"email"`
	Telephone string `json:"telephone" gorm:"telephone"`
	Adresse   string `json:"adresse" gorm:"adresse"`
}

func (Client) TableName() string {
	return "Client"
}
