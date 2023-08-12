package usersRoutesHelper

type get struct {
	Me getMe
}

type getMe struct {
	Profile            string
	Settings           string
	ProfileAndSettings string
}

func Get() get {
	route := get{}

	/*
		START LOGGED IN USER ROUTES "/me/*"
	*/
	route.Me.Profile = "/me/profile"
	route.Me.Settings = "/me/settings"
	route.Me.ProfileAndSettings = "/me/profile-and-settings"
	/*
		END LOGGED IN USER ROUTES
	*/

	return route
}
