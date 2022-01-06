package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//ConfigurationRequest - Accepted inputs for configuration
type ConfigurationRequest struct {
	MarketCode string `json:"market_code"`
}

// Configuration - Struct for configuration
type Configuration struct {
	MinLimit    json.Number `json:"min_limit"`
	MaxLimit    json.Number `json:"max_limit"`
	PaymentType string      `json:"service_name"`
	Description string      `json:"description"`
	ServiceCode string      `json:"service_code"`
	Instalments int         `json:"instalments,omitempty"`
}

// ConfigurationResponse - Final response formtat of Configuration api
type ConfigurationResponse struct {
	Configurations []Configuration `json:"configuration"`
}

// TransactionStatusRequest - Accepted inputs to fetch status of transaction
type TransactionStatusRequest struct {
	TransactionID string `json:"transaction_id"`
}

// TransactionStatus - Final response format of transaction status api
type TransactionStatus struct {
	TransactionID string      `json:"transaction_id"`
	PaymentStatus string      `json:"status"`
	Amount        json.Number `json:"amount"`
	CreatedAt     time.Time   `json:"created_at"`
}

// CustomerRequest - Accepted input format of customer data
type CustomerRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	CustomerEmail string `json:"customer_email"`
	CustomerPhone string `json:"customer_phone"`
	CustomerDOB   string `json:"customer_dob"`
	CustomerPID   int    `json:"customer_pid,omitempty"`
}

// OrderRequest - Accepted input format of order data
type OrderRequest struct {
	ServiceCode         string  `json:"service_code"`
	Amount              float32 `json:"amount"`
	MarketCode          string  `json:"market_code"`
	Currency            string  `json:"currency"`
	MerchantReferenceNo string  `json:"merchant_reference_no"`
	DiscountAmount      int     `json:"discount_amount"`
	ShippingAmount      int     `json:"shipping_amount"`
	TaxAmount           int     `json:"tax_amount"`
	Language            string  `json:"lang"`
	Signature           string  `json:"signature"`
}

// ItemRequest - Accepted input format of item data
type ItemRequest struct {
	Name           string     `json:"name"`
	Sku            string     `json:"sku"`
	Price          float32    `json:"price"`
	Quantity       int        `json:"quantity"`
	DiscountAmount int        `json:"discount_amount"`
	TaxAmount      int        `json:"tax_amount"`
	CurrencyCode   string     `json:"currency_code"`
	Categories     [][]string `json:"categories"`
}

// ContactRequest - Accepted input format of contact data
type ContactRequest struct {
	Name         string `json:"name"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zipcode      string `json:"zipcode"`
	CountryCode  string `json:"country_code"`
	PhoneNumber  string `json:"phone_number"`
}

// ShippingServiceRequest - Accepted input format of Shipping Service data
type ShippingServiceRequest struct {
	Name      string `json:"name"`
	ShippedAt string `json:"shipped_at"`
	Tracking  string `json:"tracking"`
	Priority  string `json:"priority"`
}

// TransactionRequest - Accepted input format of transaction requests
type TransactionRequest struct {
	Customer        CustomerRequest        `json:"customer"`
	Order           OrderRequest           `json:"order"`
	Items           []ItemRequest          `json:"items"`
	Billing         ContactRequest         `json:"billing"`
	Shipping        ContactRequest         `json:"shipping"`
	ShippingService ShippingServiceRequest `json:"shipping_service"`
}

// Transaction - Final response format of transaction api
type Transaction struct {
	Token           string    `json:"session_token"`
	TransactionID   string    `json:"transaction_id"`
	TokenExpiryTime time.Time `json:"expiry_time"`
	PaymentURL      string    `json:"payment_url"`
	Signature       string    `json:"signature"`
}

// DeliveryRequest - Accepted input format of delivery requests
type DeliveryRequest struct {
	DeliveredAt        string `json:"delivered_at"`
	FinalCaptureAmount int    `json:"final_capture_amount,omitempty"`
}

// Delivery - Final response format of add delivery api
type Delivery struct {
	TransactionID      string `json:"transaction_id"`
	Status             string `json:"status"`
	OriginalAmount     int    `json:"original_amount"`
	DeliveredAt        string `json:"delivered_at"`
	FinalCaptureAmount int    `json:"final_capture_amount"`
}

// RefundStatusRequest - Accepted input format of refund status api
type RefundStatusRequest struct {
	RefundID string `json:"refund_id"`
}

// Refund - Final response format of create refund and refund status api requests
type Refund struct {
	RefundID                string      `json:"refund_id"`
	Amount                  json.Number `json:"refund_amount"`
	Currency                string      `json:"currency"`
	TransactionID           string      `json:"transaction_id"`
	Status                  string      `json:"status"`
	RequestID               string      `json:"request_id"`
	MerchantRefundReference string      `json:"merchant_refund_reference"`
	DeclinedReason          string      `json:"declined_reason"`
	CreatedAt               time.Time   `json:"created_at"`
	RefundedAt              *time.Time  `json:"refunded_at,omitempty"`
}

// RefundRequest - Accepted input format of create refund api
type RefundRequest struct {
	TransactionID           string `json:"transaction_id"`
	Amount                  int    `json:"refund_amount"`
	Reason                  string `json:"reason"`
	RequestID               string `json:"request_id"`
	MerchantRefundReference string `json:"merchant_refund_reference"`
}

// Details - Final format of error details
type Details struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

// HTTPStatusCode400 Struct for HTTP Status code 400
type HTTPStatusCode400 struct {
	StatusCode int       `json:"status"`
	Message    string    `json:"message"`
	Details    []Details `json:"details"`
}

// Error defines an error received when making a requests to the API.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Error returns a string representing the error, satisfying the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("Merchant: %s Status Code: (%d)", e.Message, e.Code)
}

// Merchant defines the Merchant client.
type Merchant struct {
	MerchantKey    string
	MerchantSecret string
	Salt           string
	MarketCode     string
	Host           string
	Version        string
	Timeout        time.Duration
	Transport      http.RoundTripper
}

type CB struct {
	CreditBalance []struct {
		Amount      string `json:"amount"`
		Currency    string `json:"currency"`
		ServiceCode string `json:"service_code"`
	} `json:"credit_balance"`
}
type CBUser struct {
	CustomerMobile string `json:"customer_mobile"`
	MarketCode     string `json:"market_code"`
}
