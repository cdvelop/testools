package testools

import (
	"net/http/httptest"

	"github.com/cdvelop/model"
)

type Request struct {
	Method   string //ej:POST, GET
	Endpoint string // url
	Object   string //ej: create/files delete/x

	Data []map[string]string

	Expected []map[string]string

	model.DataConverter
	model.FetchAdapter
	*httptest.Server
}
