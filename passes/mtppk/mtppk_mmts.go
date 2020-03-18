package mtppk

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMTPPK_single = test.Cases{
	{
		N:          "1. ММТС - MTППК - ММТС - MTPPK - MMTS - MTPPK - COMPLETE",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				RequestType: test.RequestTypeOffline,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				//ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 6, 8,
				},
				Sum: 6,
			},
		},
	},
}
