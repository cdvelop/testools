module github.com/cdvelop/testools

go 1.20

require (
	github.com/cdvelop/api v0.0.96
	github.com/cdvelop/cutkey v1.0.12
	github.com/cdvelop/fetchserver v0.0.24
	github.com/cdvelop/fileserver v0.0.53
	github.com/cdvelop/logserver v0.0.32
	github.com/cdvelop/model v0.0.106
)

require (
	github.com/cdvelop/filehandler v0.0.32 // indirect
	github.com/cdvelop/input v0.0.77 // indirect
	github.com/cdvelop/maps v0.0.8 // indirect
	github.com/cdvelop/object v0.0.65 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/cdvelop/timetools v0.0.32 // indirect
	github.com/cdvelop/unixid v0.0.46 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.19.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/logserver => ../logserver

replace github.com/cdvelop/api => ../api
