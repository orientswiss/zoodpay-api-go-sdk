package tests

import (
	"testing"
)

func TestCreditBalance(t *testing.T) {
	Test_initConfig(t)
	TestHealthCheck(t)
	balance, err := merchant.GetCreditBalance("998365896609")
	if err != nil {
		t.Error("FAILED to Get Credit Balance.\n")
	} else {

		t.Log("PASS - Obtained Following Credit Balance.\n")
		for _, rs := range balance.CreditBalance {
			t.Log("Service Code : " + rs.ServiceCode + ", Amount: " + rs.Amount + " " + rs.Currency + "\n")
		}

	}

}
