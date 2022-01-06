package requests

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
)

// GetRefundStatus Used to get the status of refund
func (m *Merchant) GetRefundStatus(rsr RefundStatusRequest) (Refund, error) {
	url := m.Host + "/" + m.Version + "/refunds/" + rsr.RefundID
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Authorization", "Basic "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	resp := Refund{}

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

// CreateRefund used to create a refund application for a transaction
func (m *Merchant) CreateRefund(rr RefundRequest) (Refund, error) {
	url := m.Host + "/" + m.Version + "/refunds"
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	jsonValue, _ := json.Marshal(rr)
	reqBody := bytes.NewBuffer(jsonValue)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, reqBody)
	request.Header.Add("Authorization", "Basic "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	resp := Refund{}

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
