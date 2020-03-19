package permissions

type Permission string

const (
	USERSIGNIN      = "UserSignIn"
	USERSIGNUP      = "UserSignUp"
	GETUSERDETAILS  = "GetUserDetails"
	CREATEROLES     = "CreateRoles"
	GETROLES        = "GetRoles"
)

// EventsPermission is a map of all server side events which should be handled TODO: better explanation
var EventsPermission = map[string]string{
	"/api/v1/users/signin":             USERSIGNIN,
	"/api/v1/users/signup":             USERSIGNUP,
	"/api/v1/users/getUserDetails":     GETUSERDETAILS,
	"/api/v1/users/roles":              GETROLES,
	"/api/v1/users/createRoles":        CREATEROLES,
}

const (
	PUBLIC = "PUBLIC"
	PRIVATE = "PRIVATE"
)
var permissions = map[string]string{
	"/api/v1/users/signin" : PUBLIC,
	"/api/v1/users/signup" : PUBLIC,
	"/api/v1/users/getUserDetails" : PRIVATE,
	"/api/v1/users/roles" : PRIVATE,
	"/api/v1/users/createRoles" : PRIVATE,
	"/api/v1/users/migrateDB" : PUBLIC,

}

func GetPermission() map[string]string{
	return permissions
}
