package requests

import (
	"fmt"
)

func (m *Merchant) GenerateSignatureCreateTransaction(OR OrderRequest) string {
	s := fmt.Sprintf("%v", OR.Amount)
	return Sha512Encrypt(m.MerchantKey + "|" + OR.MerchantReferenceNo + "|" + s + "|" + OR.Currency + "|" + m.MarketCode + "|" + m.Salt)
}
