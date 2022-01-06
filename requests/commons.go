package requests

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetHTTPStatusCode400ErrorMessage(response *http.Response) string {
	var errorMsg string
	result := HTTPStatusCode400{}
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		errorMsg = err.Error()
	} else {
		json.Unmarshal(data, &result)

		if len(result.Details) > 0 {
			for _, detail := range result.Details {
				if errorMsg != "" {
					errorMsg = errorMsg + "\n"
				}
				errorMsg = errorMsg + detail.Field + ": " + detail.Error
			}
		} else {
			var m map[string]string
			json.Unmarshal(data, &m)
			errorMsg = errorMsg + m["message"]
		}
	}
	return errorMsg
}

//Sha512Encrypt used to convert string into encrypted string using Sha 512 algorithm
func Sha512Encrypt(input string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(input))
	return fmt.Sprintf("%x", sha512.Sum(nil))
}

// RandStringBytes used to generate random string (For Merchant Reference No)
func RandStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
