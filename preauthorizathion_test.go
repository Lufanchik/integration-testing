package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestPreauthorizaton(t *testing.T) {
	Run(t, casesPreauthorizaton)
}

var (
	casesPreauthorizaton = Cases{
		{
			N: "MM_successful_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
			},
		},
	}
)
