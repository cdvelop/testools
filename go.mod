module github.com/cdvelop/testools

go 1.20

require (
	github.com/cdvelop/api v0.0.54
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/fetchserver v0.0.2
	github.com/cdvelop/logserver v0.0.3
	github.com/cdvelop/model v0.0.68
)

require (
	github.com/cdvelop/gotools v0.0.56 // indirect
	github.com/cdvelop/input v0.0.51 // indirect
	github.com/cdvelop/output v0.0.15 // indirect
	github.com/cdvelop/timetools v0.0.17 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/logserver => ../logserver
