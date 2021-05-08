package handlers

import (
	"cmd/controllers"
	"cmd/models"
)

var header_one = models.Header{Name: "Content-Type", Value: "text/html"}
var header_two = models.Header{Name: "Content-Type", Value: "application/json"}

var Routes = []models.Route{
	{
		Name:           "Index",
		Method:         "GET",
		Pattern:        "/",
		DefaultHandler: controllers.WelcomeHandler,
		HeaderDefault:  models.HeaderHandleMap{HeaderType: header_one, HandlerFunc: nil},
	}, {
		Name:           "Images",
		Method:         "GET",
		Pattern:        "/images",
		DefaultHandler: controllers.InvalidHandler,
		HeaderDefault:  models.HeaderHandleMap{HeaderType: header_two, HandlerFunc: controllers.GetImages},
	},
}
