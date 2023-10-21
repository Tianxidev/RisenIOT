package router

type Group struct {
	BaseRouter
	DeviceRouter
}

var GroupApp = new(Group)
