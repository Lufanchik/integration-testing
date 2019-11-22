package http_test

import (
	"testing"
)

var (
	mtppkPasses = Cases{
		{
			//&Pass{
			//	PaymentType: PaymentTypePayment,
			//	Carrier:     carriers.Carrier_MTPPK,
			//	Auth: &processing.Auth{
			//		Sum:  0,
			//		Type: processing.Auth_AGGREGATE,
			//	},
			//	ExpectedSum: 0,
			//},
			//&Pass{
			//	PaymentType: PaymentTypePayment,
			//	Carrier:     carriers.Carrier_MTPPK,
			//	Auth: &processing.Auth{
			//		Sum:  3000,
			//		Type: processing.Auth_AGGREGATE,
			//	},
			//	ExpectedSum: 3000,
			//},
			//&Pass{
			//	PaymentType: PaymentTypePayment,
			//	Carrier:     carriers.Carrier_MTPPK,
			//	Auth: &processing.Auth{
			//		Sum:  0,
			//		Type: processing.Auth_AGGREGATE,
			//	},
			//	ExpectedSum: 0,
			//},
			//&Pass{
			//	PaymentType: PaymentTypePayment,
			//	Carrier:     carriers.Carrier_MTPPK,
			//	Auth: &processing.Auth{
			//		Sum:  0,
			//		Type: processing.Auth_AGGREGATE,
			//	},
			//	ExpectedSum: 0,
			//},
		},
	}
)

func TestMTPPK(t *testing.T) {
	Run(t, mtppkPasses)
}
