package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
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
			&Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	}
)

func TestSimplePass(t *testing.T) {
	Run(t, casesSimplePass)
}
