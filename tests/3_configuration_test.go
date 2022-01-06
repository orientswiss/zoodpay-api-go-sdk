package tests

import (
	"github.com/orientswiss/zoodpay-api-go-sdk/requests"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	Test_initConfig(t)
	TestHealthCheck(t)
	configurationsResponse, err := merchant.GetConfiguration(requests.ConfigurationRequest{
		MarketCode: merchant.MarketCode,
	})

	configurations := configurationsResponse.Configurations

	if err != nil {
		t.Errorf("%+v", err)
	} else {
		if len(configurations) == 0 {
			t.Error("FAIL - No Configuration Fetched.\n")
		} else {
			t.Log("PASS - Configuration Fetched.\n")
		}
	}
}
