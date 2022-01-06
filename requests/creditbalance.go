package requests

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
)

func (m *Merchant) GetCreditBalance(um string) (CB, error) {
	resp := CB{}
	url := m.Host + "/" + m.Version + "/customer/credit/balance"
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	cb := CBUser{
		CustomerMobile: um,
		MarketCode:     m.MarketCode,
	}

	jsonValue, _ := json.Marshal(cb)
	reqBody := bytes.NewBuffer(jsonValue)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, reqBody)
	request.Header.Add("Authorization", "Basic "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
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
