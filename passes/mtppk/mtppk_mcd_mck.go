package mtppk

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesMTPPK_MCD_MSK = test.Cases{
	{
		N:          "1. МТППК - МЦД-МСК1 VISA",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001270", //ОКРУЖНАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8800,
			},
			&test.Complete{
				StartPass: 3,
				Passes: []int{
					4, 5,
				},
				Sum: 8800,
			},
		},
	},
	{
		N:          "2. МТППК - МЦД-МСК2 VISA",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001840", //ТРИКОТАЖНАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001365", //ПЕРЕРВА
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 100,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 100,
			},
			&test.Complete{
				StartPass: 3,
				Passes: []int{
					4,
				},
				Sum: 100,
			},
		},
	},
}
