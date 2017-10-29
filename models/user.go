package main

import "time"

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	EmailPrimary string `json:"emailPrimary"`
	EmailRecovery string `json:"emailRecovery"`
	Password string `json:"password"`
	Secret2FA string `json:"secret2FA"`
	AddressShip string `json:"addressShip"`
	AddressBill string `json:"addressBill"`
	LastLogin time.Time `json:"lastLogin"`
	LoginAttempts int `json:"loginAttempts"`
	AcctLocked bool `json:"acctLocked"`
}
