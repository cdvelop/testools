package testools

import (
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/api"
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/fetchserver"
	"github.com/cdvelop/logserver"
	"github.com/cdvelop/model"
)

func NewApiTestDefault(t *testing.T, f model.FileApi, add_objects ...*model.Object) (*ApiTest, error) {

	h := &model.Handlers{
		AuthAdapter: AuthTest{},
		Logger:      logserver.Add(),
		FileApi:     f,
	}
	h.AddObjects(add_objects...)

	cutkey.AddDataConverter(h)

	err := fetchserver.AddFetchAdapter(h)
	if err != nil {
		return nil, err
	}

	conf, err := api.Add(h)
	if err != nil {
		return nil, err
	}

	mux := conf.ServeMuxAndRoutes()

	new := &ApiTest{
		T:        t,
		Server:   httptest.NewServer(mux),
		Handlers: h,
	}

	return new, nil

}
