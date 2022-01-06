# GO SDK for ZoodPay API

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)

### ZoodPay API

ZoodPay wants to provide its payment solution to every online business who may be interested in it. ZoodPay API v0 is
the latest version which offers our latest features.

[ZoodPay API Documentation](https://apidocs.zoodpay.com/)

[ZoodPay API API Simulator](https://apidocs.zoodpay.com/docs)

zoodpay-api-sdk is a Go SDK library for accessing [ZoodPay API][].

Currently, zoodpay-api-sdk requires Go version 1.13 or greater.

## Installation

zoodpay-api-sdk is compatible with modern Go releases in module mode, with Go installed:

```bash
$ go get github.com/orientswiss/zoodpay-api-go-sdk
```

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/orientswiss/zoodpay-api-go-sdk"
```

## Usage
The application configuration is represented in `./config/config.go`. When the package imported before execution of each request,
it loads the configuration from a configuration file. The path to the configuration
file is specified via the `-config` command line argument which defaults to `config.yml`. (Copy config.yml.sample to config.yml and fill with necessary information)


```go
import "github.com/orientswiss/zoodpay-api-go-sdk/requests" // with go modules disabled

func main() {

//Init Merchant
merchant = requests.NewClient()

//Health-check Endpoint 
healthCheckResponse, err := merchant.Healthcheck()

//Configuration Endpoint
configurationsResponse, err := merchant.GetConfiguration(requests.ConfigurationRequest{
MarketCode: merchant.MarketCode,
})

configurations := configurationsResponse.Configurations



//Order Type
or := requests.OrderRequest{
	
    ServiceCode:         "",
    Amount:              0,
    MarketCode:          "",
    Currency:            "",
    MerchantReferenceNo: "",
    DiscountAmount:      0,
    ShippingAmount:      0,
    TaxAmount:           0,
    Language:            "",
    Signature:           "",
    
}
//Create Signature for Transaction	
or.Signature = merchant.GenerateSignatureCreateTransaction(or)
//Transaction Endpoint
transaction, err := merchant.CreateTransaction(requests.TransactionRequest{
	//Customer Type
    Customer: requests.CustomerRequest{
    FirstName:     "",
    LastName:      "",
    CustomerEmail: "",
    CustomerPhone: "",
    CustomerDOB:   "",
    CustomerPID:   0,
},
    //Billing Type
    Billing: requests.ContactRequest{
    Name:         "",
    AddressLine1: "",
    AddressLine2: "",
    City:         "",
    State:        "",
    Zipcode:      "",
    CountryCode:  "",
    PhoneNumber:  "",
},
    //Shipping
    Shipping: requests.ContactRequest{
    Name:         "",
    AddressLine1: "",
    AddressLine2: "",
    City:         "",
    State:        "",
    Zipcode:      "",
    CountryCode:  "",
    PhoneNumber:  "",
},
    // ShippingService Type
   ShippingService: requests.ShippingServiceRequest{{
    Name:      "",
    ShippedAt: "",
    Tracking:  "",
    Priority:  "",
},
    //Items Type
    Items: []requests.ItemRequest{
    {
    Name:           "",
    Sku:            "",
    Price:          0,
    Quantity:       0,
    DiscountAmount: 0,
    TaxAmount:      0,
    CurrencyCode:   "",
    Categories: [][]string{
        {
        "",
        },
    },
 },
},
    Order: or,
})
    

//Get transaction Status from API
transactionStatus, err := merchant.GetTransactionStatus(requests.TransactionStatusRequest{
TransactionID: "",
})

//Set Delivery Date
delivery, err := merchant.AddDelivery(
requests.TransactionStatusRequest{
TransactionID: "",
},
requests.DeliveryRequest{
DeliveredAt:        "",
FinalCaptureAmount: 0,
},

//Create Refund for Paid Transaction
refund, err := merchant.CreateRefund(requests.RefundRequest{
TransactionID:           "",
Amount:                  0,
Reason:                  "",
RequestID:               "",
MerchantRefundReference: randomTransRefNo,
})

//Get Customer Credit Balance
balance, err := merchant.GetCreditBalance("")


}



```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

## Support

For any inquiry write to integration@zoodpay.com with a detailed description of the issue.

## Credits

- [ZoodPay](https://github.com/orientswiss)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.


[ZoodPay API]: https://apidocs.zoodpay.com/