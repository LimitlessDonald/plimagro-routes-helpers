package usersRoutesHelper

type validate struct {
	Me   validateMe
	Auth validateAuth
}

type validateMe struct {
	Subscription validateSubscription
}

type validateAuth struct {
	Login       string
	VerifyToken string
}
type validateSubscription struct {
	Verify string
}

func Validate() validate {
	route := validate{}

	/*
		START AUTH VALIDATION ROUTES
	*/
	route.Auth.Login = "/auth/login"
	route.Auth.VerifyToken = "/auth/verify-token"

	/*
		END AUTH VALIDATION ROUTES
	*/

	/*
		START LOGGED IN USER ROUTES "/me/*"
	*/
	route.Me.Subscription.Verify = "/me/verify-sub"

	/*
		END LOGGED IN USER ROUTES
	*/

	return route
}
