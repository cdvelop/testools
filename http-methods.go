package testools

import (
	"fmt"
	"net/url"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (r *Request) Get(data_in ...map[string]string) ([]model.Response, int, error) {

	url_values := url.Values{}

	for _, data := range data_in {
		for key, value := range data {
			// Agregar cada clave-valor al url_values
			url_values.Add(key, value)
		}
	}

	params := url_values.Encode()

	if params != "" {
		r.Endpoint = fmt.Sprintf("%s?%s", r.Endpoint, params)
	}

	// fmt.Println("ENDPOINT GET: ", r.Endpoint)

	r.Method = "GET"

	return r.SendRequest(nil)
}

func (r *Request) CutPost() ([]model.Response, int, error) {
	body, err := cutkey.Encode(r.Object, r.Data...)
	if err != nil {
		return nil, 0, err
	}

	r.Method = "POST"

	return r.SendRequest(body)
}
