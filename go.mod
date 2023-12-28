module github.com/cdvelop/testools

go 1.20

require (
	github.com/cdvelop/api v0.0.102
	github.com/cdvelop/cutkey v1.0.14
	github.com/cdvelop/fetchserver v0.0.26
	github.com/cdvelop/fileserver v0.0.54
	github.com/cdvelop/logserver v0.0.34
	github.com/cdvelop/model v0.0.107
)

require (
	github.com/cdvelop/filehandler v0.0.37 // indirect
	github.com/cdvelop/input v0.0.82 // indirect
	github.com/cdvelop/maps v0.0.8 // indirect
	github.com/cdvelop/object v0.0.68 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/cdvelop/timetools v0.0.34 // indirect
	github.com/cdvelop/unixid v0.0.49 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.19.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/logserver => ../logserver

replace github.com/cdvelop/api => ../api
