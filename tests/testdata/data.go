package testdata

func FeedInfo(mc string) FData {
	fd := FData{}
	switch mc {
	case "UZ":
		{
			fd = FData{
				Phone:    "998123456789",
				Currency: "UZS",
				Country:  "UZ",
				Amount:   25000,
			}
		}

	case "KZ":
		{
			fd = FData{
				Phone:    "77777777777",
				Currency: "KZT",
				Country:  "KZ",
				Amount:   25000,
			}
		}
	case "JO":
		{
			fd = FData{
				Phone:    "962123456789",
				Currency: "JOD",
				Country:  "JO",
				Amount:   15,
			}
		}

	}
	return fd
}

type FData struct {
	Phone    string  `json:"phone"`
	Currency string  `json:"currency"`
	Country  string  `json:"country"`
	Amount   float32 `json:"amount"`
}
