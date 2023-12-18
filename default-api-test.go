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
func NewApiTestDefault(t *testing.T, h *model.MainHandler, add_modules ...*model.Module) (at *ApiTest, err string) {

	if h.AppInfo.App_name == "" {
		h.App_name = "testApp"
		h.Business_name = "Business Testing"
		h.Business_address = "Street 54 New York"
		h.Business_phone = "555-4255-455"
	}
	h.ProductionMode = false

	if h.FileRootFolder == "" {
		h.FileRootFolder = "./test_folder"
	}

	if h.SessionBackendAdapter == nil {
		h.SessionBackendAdapter = AuthTest{}
	}

	if h.Logger == nil {
		h.Logger = logserver.AddLoggerAdapter()
	}

	h.MainHandlerAddModules(add_modules...)

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

	if h.BackendBootDataUser == nil {
		h.BackendBootDataUser = AuthTest{}
	}

	conf, err := api.Add(h)
	if err != "" {
		return nil, err
	}

	mux := conf.ServeMuxAndRoutes()

	new := &ApiTest{
		T:           t,
		Server:      httptest.NewServer(mux),
		MainHandler: h,
	}

	return new, ""

}
