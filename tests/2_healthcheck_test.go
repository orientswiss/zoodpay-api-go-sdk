package tests

import (
	"testing"
)

func TestHealthCheck(t *testing.T) {
	Test_initConfig(t)

	healthCheckResponse, err := merchant.Healthcheck()

	if err != nil {
		t.Errorf("%+v", err)
	} else {
		version := "0.0"
		if healthCheckResponse != "\"OK "+version+"\"" {
			t.Error("FAIL - API is down. Please try after sometime.\n")
		} else {
			t.Log("PASS - API is up and running.\n")
		}
	}
}
