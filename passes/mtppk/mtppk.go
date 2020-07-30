package mtppk

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesMTPPKPasses = test.Cases{
	{
		N:          "1.MTPPK,VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4400,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 4400,
			},
		},
	},
	{
		N:          "2. MTPPK, MASTERCARD",
		CardSystem: processing.CardSystem_MASTERCARD,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 2,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 2,
			},
		},
	},
	{
		N:          "3. MTPPK, FAIL",
		CardSystem: processing.CardSystem_MASTERCARD,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeIncorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 4,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 4,
			},
		},
	},
	{
		N: "4. MTPPK - MЦK - MTPPK - MМТС - MTPPK - ММ - MTPPK	- COMPLETE",
		CardSystem: processing.CardSystem_MASTERCARD,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 5,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 5,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 5,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 5,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 6, 8,
				},
				Sum: 5,
			},
		},
	},
	{
		N: "5. MTPPK - MCK - MTPPK - MMTS - MTPPK - MM - MTPPK	- COMPLETE",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
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
				SubCarrier:  carriers.SubCarrier_MM_SUB,
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
	{
		N: "6. MTPPK - MCK - MTPPK - MMTS - MTPPK - MM - MTPPK	- COMPLETE",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
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
				SubCarrier:  carriers.SubCarrier_MM_SUB,
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
