package mck

import (
	"lab.siroccotechnology.ru/tp/calculator-service/calculator/rates"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesOfflineMetroComplexMCK = test.Cases{
	{
		N: "MCK - MCK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MM - MM",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MM - MCK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MMTS - MM",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MMTS - MCK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},

	{
		N: "MCK - MMTS - MMTS",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MM - MMTS - MM",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},

	{
		N: "MCK - MM -MMTS -MCK",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MM - MMTS - MMTS",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	//
	//// Кейсы включающие транзакции с некорректной авторизацией
	{
		N: "MCK - MCK (test.AuthTypeIncorrect) - MM - MMTS -MM", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
				AuthType:    test.AuthTypeIncorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: rates.MmSubRate,
			},
		},
	},
	{
		N: "MCK - MCK (test.AuthTypeIncorrect) - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
				AuthType:    test.AuthTypeIncorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
		},
	},
	{
		N: "MCK (test.AuthTypeIncorrect) - MM - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: rates.MmSubRate,
				AuthType:    test.AuthTypeIncorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
		},
	},
}
