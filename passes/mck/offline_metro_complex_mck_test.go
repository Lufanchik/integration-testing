package mck_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestOfflineMetroComplexMCK(t *testing.T) {
	Run(t, casesOfflineMetroComplexMCK)
}

var (
	casesOfflineMetroComplexMCK = Cases{
		{
			N: "MCK - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MM - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MM - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MMTS - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MMTS - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "MCK - MMTS - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MM - MMTS - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "MCK - MM -MMTS -MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MM - MMTS - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		//
		//// Кейсы включающие транзакции с некорректной авторизацией
		{
			N: "MCK - MCK (AuthTypeIncorrect) - MM - MMTS -MM", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      2,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      2,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCK - MCK (AuthTypeIncorrect) - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      2,
				},
			},
		},
		{
			N: "MCK (AuthTypeIncorrect) - MM - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
			},
		},
	}
)
