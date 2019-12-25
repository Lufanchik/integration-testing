package mtppk

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMTPPKPasses = test.Cases{
	{
		N:          "1. MTPPK, VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 16800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 16800,
			},
		},
	},
	{
		N:          "1. MTPPK, VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 16800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 16800,
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
				ExpectedSum: 16800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 16800,
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
				ExpectedSum: 16800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 16800,
			},
		},
	},
	{
		N: "7. MM - MTPPK - MЦK - MTPPK - MМТС - MTPPK - ММ - MTPPK	- COMPLETE", //online
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 16800,
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 6, 8,
				},
				Sum: 16800,
			},
		},
	},
	{
		N: "8. MM - MTPPK - MCK - MTPPK - MMTS - MTPPK - MM - MTPPK	- COMPLETE", //offline
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				RequestType: test.RequestTypeOffline,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 16800,
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 6, 8,
				},
				Sum: 16800,
			},
		},
	},
	{
		N: "9. MM - MTPPK - MCK - MTPPK - MCK-MCK - MTPPK - MCK-MCK2 - MTPPK - MCK-MO - MTPPK - COMPLETE",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate, //  2
				RequestType: test.RequestTypeOffline,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 21000,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate, // 4
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 5,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate, // 7
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001410", //ГРАЖДАНСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000045", //ТЕКСТИЛЬЩИКИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 8,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate, //10
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001240", //БЕСКУДНИКОВО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000620", //ХЛЕБНИКОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     11,
				ExpectedSum: 700,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 7, 10, 13,
				},
				Sum: 21000,
			},
		},
	},
}
