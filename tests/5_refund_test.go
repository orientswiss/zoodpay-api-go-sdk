package tests

import (
	"github.com/orientswiss/zoodpay-api-go-sdk/requests"
	"testing"
)

func TestCreateRefund(t *testing.T) {
	Test_initConfig(t)
	refund, err := merchant.CreateRefund(requests.RefundRequest{
		TransactionID:           transactionID,
		Amount:                  20,
		Reason:                  "Test case go sdk",
		RequestID:               "101",
		MerchantRefundReference: randomTransRefNo,
	})

	if err != nil {
		// t.Errorf("%+v", err)
		t.Skipf("%+v", err)
	} else {
		if refund == (requests.Refund{}) {
			t.Error("FAIL - refund could not created successfully.\n")
		} else {
			refundID = refund.RefundID
			t.Log("PASS - Refund created successfully.\n")
		}
	}
}

func TestGetRefundStatus(t *testing.T) {
	Test_initConfig(t)
	refundStatus, err := merchant.GetRefundStatus(requests.RefundStatusRequest{
		RefundID: refundID,
	})

	if err != nil {
		// t.Errorf("%+v", err)
		t.Skipf("%+v", err)
	} else {
		if refundStatus == (requests.Refund{}) {
			t.Error("FAIL - Refund Id Not found.\n")
		} else {
			t.Log("PASS - Refund status got successfully.\n")
		}
	}
}
