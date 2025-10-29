package entities

type LoginStruct struct {
	Id        int     `json:"id"`
	Login     string  `json:"login"`
	Password  string  `json:"password"`
	Name      string  `json:"name"`
	Lastname  string  `json:"lastname"`
	Type      string  `json:"type"`
	Contact   string  `json:"contact"`
	Mail      string  `json:"mail"`
	Adresse   string  `json:"adresse"`
	Latitude  float64 `json:"latitude" gorm:"latitude"`
	Longitude float64 `json:"longitude" gorm:"longitude"`
}

func (LoginStruct) TableName() string {
	return "user"
}

type RequestLogin struct {
	Login string `json:"username"`
	Mdp   string `json:"password"`
}
type RefreshToken struct {
	Id int `json:"id"`
}

// Modèle à tester
type User struct {
	Id           int     `gorm:"primaryKey;column:id" json:"id"`
	Login        string  `gorm:"column:login;size:100;not null" json:"login"`
	Password     string  `gorm:"column:password;size:255;not null" json:"-"`
	Name         string  `gorm:"column:name;size:100;not null" json:"name"`
	LastName     string  `gorm:"column:lastname;size:100;not null" json:"lastname"`
	Type         string  `gorm:"column:type;size:50;not null" json:"type"`
	Contact      string  `gorm:"column:contact;size:20;not null" json:"contact"`
	Mail         string  `gorm:"column:mail;size:150;not null" json:"mail"`
	Adresse      string  `gorm:"column:adresse;type:text;not null" json:"adresse"`
	Latitude     *string `gorm:"column:latitude;size:45" json:"latitude,omitempty"`
	Longitude    *string `gorm:"column:longitude;size:45" json:"longitude,omitempty"`
	CommercantID *int    `gorm:"column:commercant_id" json:"commercant_id,omitempty"` // Optionnel

	// Relation GORM
	Commercant *Commercant `gorm:"foreignKey:CommercantID" json:"commercant,omitempty"`
}

func (User) TableName() string {
	return "user"
}

type SessionData struct {
	User LoginStruct `json:"user"`
}

type UserResponse struct {
	Id        int     `json:"id"`
	Login     string  `json:"login"`
	Name      string  `json:"name"`
	Lastname  string  `json:"lastname"`
	Type      string  `json:"type"`
	Contact   string  `json:"contact"`
	Mail      string  `json:"mail"`
	Adresse   string  `json:"adresse"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
