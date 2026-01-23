package models

type Wallet struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
}

var EmptyWallet = Wallet{}
