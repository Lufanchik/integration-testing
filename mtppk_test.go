package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

var (
	mtppkPasses = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				Carrier:     carriers.Carrier_MTPPK,
				Auth: &processing.Auth{
					Sum:  0,
					Type: processing.Auth_AGGREGATE,
				},
				ExpectedSum: 0,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				Carrier:     carriers.Carrier_MTPPK,
				Auth: &processing.Auth{
					Sum:  3000,
					Type: processing.Auth_AGGREGATE,
				},
				ExpectedSum: 3000,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				Carrier:     carriers.Carrier_MTPPK,
				Auth: &processing.Auth{
					Sum:  0,
					Type: processing.Auth_AGGREGATE,
				},
				ExpectedSum: 0,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				Carrier:     carriers.Carrier_MTPPK,
				Auth: &processing.Auth{
					Sum:  0,
					Type: processing.Auth_AGGREGATE,
				},
				ExpectedSum: 0,
			},
		},
	}
)

func TestMTPPK(t *testing.T) {
	Run(t, mtppkPasses)
}
