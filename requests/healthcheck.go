package requests

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Healthcheck used to check the availability of the api
func (m *Merchant) Healthcheck() (string, error) {
	url := m.Host + "/healthcheck"
	response, err := http.Get(url)
	resp := ""

	if err != nil {
		return resp, Error{
			Message: err.Error(),
			Code:    response.StatusCode,
		}
	}

	if response.StatusCode == 404 {
		return resp, Error{
			Message: "The requested resource was not found.",
			Code:    404,
		}
	} else if response.StatusCode == 400 {
		errorMsg := GetHTTPStatusCode400ErrorMessage(response)

		return resp, Error{
			Message: errorMsg,
			Code:    400,
		}
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		resp = strings.TrimSuffix(string(data), "\n")
	}

	return resp, err
}
