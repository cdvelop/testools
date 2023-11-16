package testools

import (
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/model"
)

type Request struct {
	TestName string

	Method   string //ej:POST, GET
	Endpoint string // ej: upload,create,read,update,delete,file
	Object   string //ej: create/files delete/x

	Data     any
	Expected any

	Analysis func(rq *Request, resp []map[string]string, err error)

	*ApiTest
}

type ApiTest struct {
	*testing.T
	*model.Handlers
	*httptest.Server
}
