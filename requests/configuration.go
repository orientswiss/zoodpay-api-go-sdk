package requests

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
)

// GetConfiguration returns configuration of the Merchant
func (m *Merchant) GetConfiguration(cr ConfigurationRequest) (ConfigurationResponse, error) {
	url := m.Host + "/" + m.Version + "/configuration"
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	jsonValue, _ := json.Marshal(cr)
	reqBody := bytes.NewBuffer(jsonValue)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, reqBody)
	request.Header.Add("Authorization", "Basic "+token)
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	resp := ConfigurationResponse{}

	if err != nil {
		return resp, Error{
			Message: err.Error(),
			Code:    response.StatusCode,
		}
	}

	if response.StatusCode == 401 {
		return resp, Error{
			Message: "You are not authenticated to perform the requested action.",
			Code:    401,
		}
	} else if response.StatusCode == 400 {
		errorMsg := GetHTTPStatusCode400ErrorMessage(response)

		return resp, Error{
			Message: errorMsg,
			Code:    400,
		}
	} else {
		json.NewDecoder(response.Body).Decode(&resp)
	}

	return resp, err
}
