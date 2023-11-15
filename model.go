package testools

import (
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/model"
)

type Request struct {
	Method   string //ej:POST, GET
	Endpoint string // url
	Object   string //ej: create/files delete/x

	Data any

	Expected any

	model.DataConverter
	model.FetchAdapter
	*httptest.Server
}

type ApiTest struct {
	*testing.T
	*httptest.Server
	*model.Handlers
}
