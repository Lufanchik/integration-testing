package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

func TestComplexTimeMCD(t *testing.T) {
	Run(t, casesComplexTimeMCD)
}

var (
	casesComplexTimeMCD = Cases{
		{
			N: "MCD train fail",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					Now:         NowCustom(10, 00),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
					Now:    NowCustom(10, 30),
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_EGRESS,
					},
					ExpectedSum: 700,
					Ingress:     2,
					Now:         NowCustom(12, 30),
				},
			},
		},
	}
)
