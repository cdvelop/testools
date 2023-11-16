package testools

func (a ApiTest) BuildEndPoint(r Request) string {

	endpoint := a.Server.URL

	if r.Endpoint != "" {
		endpoint += "/" + r.Endpoint
	}
	if r.Object != "" {
		endpoint += "/" + r.Object
	}

	return endpoint

}
