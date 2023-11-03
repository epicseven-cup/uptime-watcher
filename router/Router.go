package main

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	path     []string
	app      echo.Echo
	database *Database
}

func NewRouter() *Router {
	return &Router{}
}

// Adds the path to router's path which checks if the path is avaaible to the user
// this should return the success of the method
func (*Router) addPath() bool {
	return true
}
