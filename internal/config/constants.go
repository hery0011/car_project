package config

const ROOT_FOLDER_VAR = "PROJECT_ROOT_FOLDER"
const PPE = "p5aNLsDH3r$YDQQ5"

// routes constant
const (
	//swager
	SwaggerPath = "/swagger/*any"

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
)
