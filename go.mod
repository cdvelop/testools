module github.com/cdvelop/testools

go 1.20

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/fetchserver v0.0.9
	github.com/cdvelop/logserver v0.0.10
	github.com/cdvelop/model v0.0.76
)

require (
	github.com/cdvelop/filehandler v0.0.13 // indirect
	github.com/cdvelop/maps v0.0.7 // indirect
	github.com/cdvelop/object v0.0.40 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/unixid v0.0.25 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.19.0 // indirect
)

require (
	github.com/cdvelop/api v0.0.67
	github.com/cdvelop/fileserver v0.0.30
	github.com/cdvelop/input v0.0.59 // indirect
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.25 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/maps => ../maps

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/logserver => ../logserver
