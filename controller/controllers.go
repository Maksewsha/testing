package controller

type Controllers struct {
	UserController UserController
	PageController PageController
}

func NewControllers() *Controllers {
	return &Controllers{
		UserController: UserController{},
		PageController: PageController{},
	}
}
