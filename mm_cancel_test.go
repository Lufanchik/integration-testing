package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestMetroComplexCancel(t *testing.T) {
	Run(t, casesMetroComplexCancel)
}

var (
	casesMetroComplexCancel = Cases{

		{
			N: "MCK (AuthTypeIncorrect) - MM - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
			},
		},

		{
			N: "MM - MCK - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
			},
		},

		{
			N: "MCK - MM - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
			},
		},


	}
)
