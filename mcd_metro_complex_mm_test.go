package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

func TestComplexMCD(t *testing.T) {
	Run(t, casesComplexMCD)
}

// "MM - MM" (Две платные поездки)
var (
	casesComplexMCD = Cases{
		{
			N: "MCD",
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					Terminal: &processing.Terminal{
						Station:   "2002780", //мцд2мск
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					Terminal: &processing.Terminal{
						Station:   "2000200", //мцд2мск
						Direction: processing.TerminalDirection_EGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
	}
)
