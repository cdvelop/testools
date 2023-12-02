package testools

import (
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/api"
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/fetchserver"
	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/logserver"
	"github.com/cdvelop/model"
)

// default:
// h.FileRootFolder = "./test_folder"
// h.App_name = "testApp"
func NewApiTestDefault(t *testing.T, h *model.Handlers, add_objects ...*model.Object) (at *ApiTest, err string) {

	if h.AppInfo.App_name == "" {
		h.App_name = "testApp"
		h.App_version = "0.0.1"
		h.Business_name = "Business Testing"
		h.Business_address = "Street 54 New York"
		h.Business_phone = "555-4255-455"
	}
	h.ProductionMode = false

	if h.FileRootFolder == "" {
		h.FileRootFolder = "./test_folder"
	}

	if h.AuthAdapter == nil {
		h.AuthAdapter = AuthTest{}
	}

	if h.Logger == nil {
		h.Logger = logserver.Add()
	}

	h.AddObjects(add_objects...)

	cutkey.AddDataConverter(h)

	err = fetchserver.AddFetchAdapter(h)
	if err != "" {
		return nil, err
	}

	if h.FileApi == nil {
		_, err := fileserver.AddFileApi(h)
		if err != "" {
			return nil, err
		}
	}

	conf, err := api.Add(h)
	if err != "" {
		return nil, err
	}

	mux := conf.ServeMuxAndRoutes()

	new := &ApiTest{
		T:        t,
		Server:   httptest.NewServer(mux),
		Handlers: h,
	}

	return new, ""

}
