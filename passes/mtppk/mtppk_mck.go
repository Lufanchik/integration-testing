package mtppk

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesMTPPK_MCK = test.Cases{
	{
		N:          "1.MTPPK-MCK-MM-MCK-MTPPK",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 8400,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					5,
				},
				Sum: 8400,
			},
		},
	},

	{
		N: "2.MCK-MM-MTPPK-MTPPK-MCK-MM",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8400,
			},
			&test.Complete{
				StartPass: 3,
				Passes: []int{
					4,
				},
				Sum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      6,
			},
		},
	},

	{
		N: "3.MTPPK-MCD1_MO-MCK-MM-MMTS-MTPPK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //????????????????
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 5100,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001060", //??????????????
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 8400,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					7,
				},
				Sum: 8400,
			},
		},
	},

	{
		N: "4.MTPPK-MGT-MCK-MTPPK-MMTS",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      3,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					4,
				},
				Sum: 8400,
			},
		},
	},

	{
		N: "5.MTPPk-MCK-MTPPK-MCK-MTPPK-MCK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 12600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 12600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 12600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					3, 5,
				},
				Sum: 12600,
			},
		},
	},

	{
		N: "6.MCK-MTPPK-MM-MTPPK-MCD_MSK-MTPPK-MMTS-MTPPK-MMTS",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
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
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 16800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //????????
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001270", //????????????????
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 5,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 16800,
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
				ExpectedSum: 16800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 7, 9,
				},
				Sum: 16800,
			},
		},
	},

	{
		N: "7.MTPPK-MCK-MMTS-MCK-MTPPK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 8400,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					5,
				},
				Sum: 8400,
			},
		},
	},
}
