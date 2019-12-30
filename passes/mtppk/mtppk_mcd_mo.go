package mtppk

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMTPPKMCDMO = test.Cases{
	{
		N:          "1. МТППК - МСД-МО1 VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000115", //ЛОБНЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
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
}
