package config

const ROOT_FOLDER_VAR = "PROJECT_ROOT_FOLDER"
const PPE = "p5aNLsDH3r$YDQQ5"

// routes constant
const (
	//swager
	SwaggerPath = "/swagger/*any"

	//adminPath
	AdminPath = "/admin"

	//dashboardPaath
	DashPath = "/dash"

	// authentification
	AuthPath = "/auth"
	Login    = "/login"
	Logout   = "/logout"
	Refresh  = "/refresh"

	// user
	UserPath = "/user"
	Creat    = "/creatUser"
	Delete   = "/:idUser/delete"
	Update   = "/updateUser"

	//profil
	ProfilPath   = "/profil"
	GetProfil    = "/list"
	AssignProfil = "/assignProfil"

	//article
	ArticlePath = "/article"
	ListArticle = "/list"

	//categorie
	CategoriePath = "/categorie"
	ListCategorie = "/list"
)
