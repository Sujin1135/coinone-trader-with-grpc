package provider

type RequestHelper struct {
	accessToken string
	nonce       string
}

func (r *RequestHelper) Get() {

}

func (r *RequestHelper) Post() {

}

func newRequestHelper() *RequestHelper {
	return &RequestHelper{accessToken: "", nonce: ""}
}
