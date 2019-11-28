package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
	"time"
)

func TestComplexTimeMCD(t *testing.T) {
	Run(t, casesComplexTimeMCD)
}

var (
	casesComplexTimeMCD = Cases{
		{
			N: "MCD egress after complex period",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         NowCustom(10, 30),
					ExpectedSum: 4900,
					TimeToWait:  35 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001140",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     NowCustom(11, 22),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         NowCustom(12, 35),
					Parent:      1,
				},
			},
		},
		//{
		//	N: "MCD egress after complex period",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000065",
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:         NowCustom(10, 30),
		//			ExpectedSum: 4900,
		//			TimeToWait:  360 * time.Minute,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2002780",
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//			Now:     NowCustom(17, 9),
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         NowCustom(17, 31),
		//			Parent:      1,
		//		},
		//	},
		//},
		//{
		//	N: "MCD egress after complex period",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         NowCustom(10, 00),
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Now:         NowCustom(11, 31),
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Now:         NowCustom(13, 02),
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//			Now:		 NowCustom(14, 33),
		//			TimeToWait:  5 * time.Minute,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//			Now:		 NowCustom(15, 00),
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
	}
)
