module github.com/cdvelop/testools

go 1.20

require (
	github.com/cdvelop/api v0.0.59
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/fetchserver v0.0.3
	github.com/cdvelop/logserver v0.0.4
	github.com/cdvelop/model v0.0.69
)

require (
	github.com/cdvelop/filehandler v0.0.6 // indirect
	github.com/cdvelop/maps v0.0.5 // indirect
	github.com/cdvelop/object v0.0.32 // indirect
	github.com/cdvelop/unixid v0.0.19 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.18.0 // indirect
)

require (
	github.com/cdvelop/fileserver v0.0.26
	github.com/cdvelop/gotools v0.0.58 // indirect
	github.com/cdvelop/input v0.0.53 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.19 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/logserver => ../logserver
