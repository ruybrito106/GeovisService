package router

import (
	"net/http"
)

type Route struct {
	Name    string
	Patter  string
	Methods []string
	Handler http.HandlerFunc
}
