package mck

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var CasesComplexTimeMCK = test.Cases{
	{
		N: "1. MCK -  MCD1_MSK - MM - MCD2_MSK - MCD2_MO",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now:         test.NowCustom(10, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station: "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
				Now:     test.NowCustom(10, 20),
				TimeToWait:  10 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station: "2001270", //ОКРУЖНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 2,
				Now:     test.NowCustom(11, 00),

			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent: 1,
				Now:         test.NowCustom(11, 01),

			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station: "2002780", //ДЕПО
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
				Now:     test.NowCustom(11, 02),
				TimeToWait:  8 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station: "2000045", //ТЕКСТИЛЬЩИКИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 5,
				Now:     test.NowCustom(11, 16),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station: "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
				Now:        test.NowCustom(15, 00),
				TimeToWait: 5 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Now:   	 	 test.NowCustom(15, 33 ),
				Ingress: 7,
				ExpectedSum: 700,

			},
		},
	},

	{
		N: "2. MCD_MO - MCK - MM - MMTS - MCD_MO",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000065", //ПОДОЛЬСК
					Direction: processing.TerminalDirection_INGRESS,
				},
				Now:         test.NowCustom(10, 00),
				ExpectedSum: 4900,
				TimeToWait:  120 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2002780", //ДЕПО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowCustom(12, 39),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Now:         test.NowCustom(12, 43),
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Now:         test.NowCustom(12, 49),
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Now:         test.NowCustom(13, 00),
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station: "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
				Now:        test.NowCustom(15, 00),
				TimeToWait: 5 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Now:   	 	 test.NowCustom(15, 33 ),
				Ingress: 6,
				ExpectedSum: 700,
			},
		},
	},
}