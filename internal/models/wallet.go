package models

type Wallet struct {
	Id      string
	Created string
	Updated string
	Name    string
}

var EmptyWallet = Wallet{}
