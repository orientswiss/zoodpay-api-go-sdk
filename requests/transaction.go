package requests

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
)

// GetTransactionStatus used to get the status of a transaction
func (m *Merchant) GetTransactionStatus(tsr TransactionStatusRequest) (TransactionStatus, error) {
	url := m.Host + "/" + m.Version + "/transactions/" + tsr.TransactionID
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	resp := TransactionStatus{}

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

// CreateTransaction used to create a new transaction
func (m *Merchant) CreateTransaction(tr TransactionRequest) (Transaction, error) {
	url := m.Host + "/" + m.Version + "/transactions"
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	jsonValue, _ := json.Marshal(tr)
	reqBody := bytes.NewBuffer(jsonValue)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, reqBody)
	request.Header.Add("Authorization", "Basic "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	resp := Transaction{}

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

// AddDelivery used to add a delivery date to a transaction
func (m *Merchant) AddDelivery(tsr TransactionStatusRequest, dr DeliveryRequest) (Delivery, error) {
	url := m.Host + "/" + m.Version + "/transactions/" + tsr.TransactionID + "/delivery"
	token := b64.StdEncoding.EncodeToString([]byte(m.MerchantKey + ":" + m.MerchantSecret))

	jsonValue, _ := json.Marshal(dr)
	reqBody := bytes.NewBuffer(jsonValue)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPut, url, reqBody)
	request.Header.Add("Authorization", "Basic "+token)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	resp := Delivery{}

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
