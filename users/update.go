package usersRouteHelper

type update struct {
	Me updateMe
}

type updateMe struct {
	Profile  string
	Settings string
	Password string
	Email    string
}

func Update() update {
	route := update{}

	/*
		START LOGGED IN USER ROUTES "/me/*"
	*/
	route.Me.Profile = "/me/profile"
	route.Me.Settings = "/me/settings"
	route.Me.Password = "/me/change-password"
	route.Me.Email = "/me/change-email"
	/*
		END LOGGED IN USER ROUTES
	*/

	return route
}
