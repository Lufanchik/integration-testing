package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

var (
	casesSimplePass = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestSimplePass(t *testing.T) {
	Run(t, casesSimplePass)
}
