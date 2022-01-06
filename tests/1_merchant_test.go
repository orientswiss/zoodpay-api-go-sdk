package tests

import (
	"github.com/orientswiss/zoodpay-api-go-sdk/requests"
	"testing"
)

var merchant *requests.Merchant
var randomTransRefNo, transactionID, refundID string

func Test_initConfig(t *testing.T) {
	r := requests.InitConfig()
	if r == false {
		t.Error("FAIL - Error in Init Config.\n")
	} else {
		randomTransRefNo = requests.RandStringBytes(15)

		merchant = requests.NewClient()
		t.Log("Merchant Detailed Initiated")
	}

}
