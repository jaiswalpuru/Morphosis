package models

import "net/http"

type Route struct {
	Name           string
	Method         string
	Pattern        string
	DefaultHandler http.HandlerFunc
	HeaderDefault  HeaderHandleMap
}

type HeaderHandleMap struct {
	HeaderType  Header
	HandlerFunc http.HandlerFunc
}

type Header struct {
	Name  string
	Value string
}
