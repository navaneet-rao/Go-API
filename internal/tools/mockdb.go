package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"luke": {
		AuthToken: "123ABC",
		Username:  "luke",
	},
	"navaneet": {
		AuthToken: "098XYZ",
		Username:  "navaneet",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"luke": {
		Coins:    100,
		Username: "luke",
	},
	"navaneet": {
		Coins:    200,
		Username: "navaneet",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
