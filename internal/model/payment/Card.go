package payment

import "time"

type Card struct {
	OwnerId        uint64
	PaymentSystem  string
	Number         string
	HolderName     string
	ExpirationDate time.Time
	CvcCvv         string
}
