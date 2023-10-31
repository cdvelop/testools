package testools

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cdvelop/model"
)

func (r *Request) SendRequest(body []byte) ([]model.Response, int, error) {

	// req, err := http.NewRequest(r.Method, r.Server.URL+r.Endpoint, bytes.NewBuffer(body))
	req, err := http.NewRequest(r.Method, r.Server.URL+r.Endpoint, nil)
	if err != nil {
		return nil, 0, err
	}

	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}
	// fmt.Println("SOLICITUD MÉTODO: ", req.Method)

	if r.ContentType != "" {
		req.Header.Set("Content-Type", r.ContentType)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	// fmt.Printf("RESPUESTA BODY: %v\n", resp)

	if body != nil {
		return r.DecodeResponses(resp), res.StatusCode, nil

	}

	return nil, res.StatusCode, nil

}
