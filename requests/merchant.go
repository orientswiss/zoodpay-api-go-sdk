package requests

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/orientswiss/zoodpay-api-go-sdk/config"
)

var configFile = "../config.yml"
var host, version, defaultCode, key, secret, salt, marketCode string

func InitConfig() bool {
	cfg, err := config.Load(configFile)
	if err != nil {
		return false
	}

	// host is the default host of Merchant Api
	host = cfg.Host

	// version of the merchant api
	version = cfg.Version

	// defaultCode is the default error code for failures.
	defaultCode = cfg.DefaultCode

	// Credential Source Type

	switch cfg.CST {
	case "yml":
		key = cfg.YML["merchant_key"]
		secret = cfg.YML["merchant_secret"]
		salt = cfg.YML["merchant_salt"]
		marketCode = cfg.YML["market_code"]
	case "db":
		key, secret, salt, marketCode = getCredentialsFromDB(cfg.DB)
	default:
		key = cfg.YML["merchant_key"]
		secret = cfg.YML["merchant_secret"]
		salt = cfg.YML["merchant_salt"]
		marketCode = cfg.YML["market_code"]
	}
	return true
}

// NewClient returns a new Merchant API client which can be used to make RPC requests.
func NewClient() *Merchant {

	InitConfig()

	return &Merchant{
		MerchantKey:    key,
		MerchantSecret: secret,
		Salt:           salt,
		MarketCode:     marketCode,
		Host:           host,
		Version:        version,
		Timeout:        10 * time.Second,
		Transport:      http.DefaultTransport,
	}
}

func getCredentialsFromDB(dbConfig map[string]string) (string, string, string, string) {
	// Open up our database connection.
	connectionString := dbConfig["user_name"] + ":" + dbConfig["password"] + "@tcp(" + dbConfig["host"] + ")/" + dbConfig["database"]
	db, err := sql.Open("mysql", connectionString)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	selectQuery := "SELECT zoodpay_merchant_key, zoodpay_merchant_secret, zoodpay_merchant_salt, zoodpay_merchant_marketCode FROM " + dbConfig["table"] + " limit 1"

	err = db.QueryRow(selectQuery).Scan(&key, &secret, &salt, &marketCode)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return key, secret, salt, marketCode
}
