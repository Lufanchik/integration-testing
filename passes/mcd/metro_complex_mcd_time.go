package mcd

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var CasesComplexTimeMCD = test.Cases{
	{
		N: "MCD egress after complex period",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055",
					Direction: processing.TerminalDirection_INGRESS,
				},
				Now:         test.NowCustom(10, 30),
				ExpectedSum: 4900,
				TimeToWait:  35 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001140",
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowCustom(11, 22),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Now:         test.NowCustom(12, 35),
				Parent:      1,
			},
		},
	},
	{
		N: "MCD egress after complex period",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000065",
					Direction: processing.TerminalDirection_INGRESS,
				},
				Now:         test.NowCustom(10, 30),
				ExpectedSum: 4900,
				TimeToWait:  360 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2002780",
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowCustom(17, 9),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Now:         test.NowCustom(17, 31),
				Parent:      1,
			},
		},
	},
	{
		N: "MCD egress after complex period",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Now:         test.NowCustom(10, 00),
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Now:         test.NowCustom(11, 31),
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Now:         test.NowCustom(13, 02),
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000700", //ТЕСТОВСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
				Now:         test.NowCustom(14, 33),
				TimeToWait:  5 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     4,
				Now:         test.NowCustom(15, 00),
				ExpectedSum: 700,
			},
		},
	},
}
