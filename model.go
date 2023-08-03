package testools

import (
	"net/http/httptest"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

type Request struct {
	Endpoint    string //ej: create/files delete/x
	Method      string //ej: "PUT","GET"
	ContentType string //ej: multipart/form-data, application/json

	*model.Object
	Data []map[string]string

	ExpectedCode int
	*httptest.Server
	*cutkey.Cut
}
