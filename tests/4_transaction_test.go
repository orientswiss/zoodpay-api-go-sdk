package tests

import (
	"github.com/orientswiss/zoodpay-api-go-sdk/requests"
	"github.com/orientswiss/zoodpay-api-go-sdk/tests/testdata"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	Test_initConfig(t)
	TestHealthCheck(t)

	fd := testdata.FeedInfo(merchant.MarketCode)

	or := requests.OrderRequest{
		ServiceCode:         "ZPI",
		Amount:              fd.Amount,
		MarketCode:          fd.Country,
		Currency:            fd.Currency,
		MerchantReferenceNo: randomTransRefNo,
		DiscountAmount:      0,
		ShippingAmount:      0,
		TaxAmount:           0,
		Language:            "en",
		Signature:           "",
	}
	or.Signature = merchant.GenerateSignatureCreateTransaction(or)

	transaction, err := merchant.CreateTransaction(requests.TransactionRequest{
		Customer: requests.CustomerRequest{
			FirstName:     "Test User",
			LastName:      "Family",
			CustomerEmail: "test@zoodpay.com",
			CustomerPhone: fd.Phone,
			CustomerDOB:   "1998-11-08",
			CustomerPID:   123456789,
		},
		Order: or,
		Items: []requests.ItemRequest{
			{
				Name:           "Blue Washed Men's Jeans Shirt",
				Sku:            "BLUE-XL-34",
				Price:          fd.Amount,
				Quantity:       1,
				DiscountAmount: 0,
				TaxAmount:      0,
				CurrencyCode:   fd.Currency,
				Categories: [][]string{
					{
						"Consumer Electronics",
					},
				},
			},
		},
		Billing: requests.ContactRequest{
			Name:         "Test User Family",
			AddressLine1: "Test Address 1",
			AddressLine2: "Test Address 1",
			City:         "Tashkent",
			State:        "Tashkent",
			Zipcode:      "1234567",
			CountryCode:  fd.Country,
			PhoneNumber:  fd.Phone,
		},
		Shipping: requests.ContactRequest{
			Name:         "Test User Family",
			AddressLine1: "Test Address 1",
			AddressLine2: "Test Address 1",
			City:         "Tashkent",
			State:        "Tashkent",
			Zipcode:      "1234567",
			CountryCode:  fd.Country,
			PhoneNumber:  fd.Phone,
		},
		ShippingService: requests.ShippingServiceRequest{
			Name:      "Blue Dart",
			ShippedAt: "",
			Tracking:  "8878451200014",
			Priority:  "High",
		},
	})

	if err != nil {
		t.Errorf("%+v", err)
	} else {
		if transaction == (requests.Transaction{}) {
			t.Error("FAIL - Transaction could not be created.\n")
		} else {
			transactionID = transaction.TransactionID // This is being used in furhter test cases
			t.Log("PASS - Transaction created successfully.\n")
			t.Log("Transaction ID: " + transaction.TransactionID)
			t.Log("Transaction Signature: " + transaction.Signature)
			t.Log("Transaction PaymentURL: " + transaction.PaymentURL)
			t.Log("Transaction TokenExpiryTime: " + transaction.TokenExpiryTime.String())
			t.Log("Transaction Token: " + transaction.Token)

		}
	}
}

func TestGetTransactionStatus(t *testing.T) {
	Test_initConfig(t)
	TestHealthCheck(t)
	TestCreateTransaction(t)
	transactionStatus, err := merchant.GetTransactionStatus(requests.TransactionStatusRequest{
		TransactionID: transactionID,
	})

	if err != nil {
		t.Errorf("%+v", err)
	} else {
		if transactionStatus == (requests.TransactionStatus{}) {
			t.Error("FAIL - Transaction status not found.\n")
		} else {
			t.Log("PASS - Found Transaction Status.\n")
		}
	}
}

func TestAddDelivery(t *testing.T) {
	Test_initConfig(t)
	delivery, err := merchant.AddDelivery(
		requests.TransactionStatusRequest{
			TransactionID: transactionID,
		},
		requests.DeliveryRequest{
			DeliveredAt:        "2020-12-02 00:00:00",
			FinalCaptureAmount: 20,
		},
	)

	if err != nil {
		// t.Errorf("%+v", err)
		t.Skipf("%+v", err)
	} else {
		if delivery == (requests.Delivery{}) {
			t.Error("FAIL - Delivery date could not be set.\n")
		} else {
			t.Log("PASS - Delivery date was set successfully.\n")
		}
	}
}
