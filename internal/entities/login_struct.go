package entities

type LoginStruct struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Type     string `json:"type"`
	Contact  string `json:"contact"`
	Mail     string `json:"mail"`
	Adresse  string `json:"adresse"`
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
	ID    uint
	Name  string
	Email string
}

type SessionData struct {
	User LoginStruct `json:"user"`
}

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
