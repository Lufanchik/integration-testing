package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

var (
	mtppkPasses = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 0,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MTPPK,
				Sum:         3000,
				ExpectedSum: 3000,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 0,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 0,
			},
		},
	}
)

func TestMTPPK(t *testing.T) {
	Passes(t, mtppkPasses)
}
