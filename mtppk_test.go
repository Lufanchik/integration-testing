package http_test

import (
	"testing"
)

var (
	mtppkPasses = Cases{
		{
			//Pass{
			//	PaymentType: PaymentTypeFullPayment,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MTPPK,
			//},
			//Pass{
			//	PaymentType: PaymentTypeFree,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MTPPK,
			//	Sum:         3000,
			//},
			//Pass{
			//	PaymentType: PaymentTypeFree,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MTPPK,
			//},
			//Pass{
			//	PaymentType: PaymentTypeFullPayment,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MTPPK,
			//},
		},
	}
)

func TestMTPPK(t *testing.T) {
	Passes(t, mtppkPasses)
}
