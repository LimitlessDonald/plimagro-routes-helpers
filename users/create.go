package usersRouteHelper

type create struct {
	Adduser string
}

func Create() create {
	route := create{}

	/*
		START LOGGED IN USER ROUTES "/me/*"
	*/
	route.Adduser = "/add"
	/*
		END LOGGED IN USER ROUTES
	*/

	return route
}
