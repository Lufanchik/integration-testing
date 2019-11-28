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
			N: "МЦД МО/МСК 1 - ММ - МЦД МСК 2 - МЦД_МСК/МО_1", //46
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ 2000009
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ 2000115
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},
		//{
		//	N: "MCD egress after complex period",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000055",
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:         NowCustom(10, 30),
		//			ExpectedSum: 4900,
		//			TimeToWait:  35 * time.Minute,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155",
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//			Now:     NowCustom(11, 30),
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         NowCustom(12, 01),
		//			Parent:      1,
		//		},
		//	},
		//},
	}
)
