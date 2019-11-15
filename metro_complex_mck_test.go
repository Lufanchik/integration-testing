package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestMetroComplexMCK(t *testing.T) {
	Run(t, casesMetroComplexMCK)
}

var (
	casesMetroComplexMCK = Cases{
		{
			N: "MCK - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
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
					PaymentType: PaymentTypeFullPayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		//// MCK - MM - MCK
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//// MCK - MMTS - MM
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//// MCK - MMTS - MCK
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//// MCK - MMTS - MMTS
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//// MCK - MM - MMTS - MM
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//// MCK - MM -MMTS -MCK
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//
		//// MCK - MM -MMTS -MMTS
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//
		//// Кейсы включающие транзакции с некорректной авторизацией
		//// "MCK - MCK (AuthTypeIncorrect) - MM - MMTS -MM" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//		AuthType:    AuthTypeIncorrect,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		Parent:      2,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      2,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		ExpectedSum: 4200,
		//	},
		//},
		//// "MCK - MCK (AuthTypeIncorrect) - MMTS" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//		AuthType:    AuthTypeIncorrect,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      2,
		//	},
		//},
		//
		//// "MCK (AuthTypeIncorrect) - MM - MMTS" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
		//{
		//	&Pass{
		//		PaymentType: PaymentTypeFullPayment,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//		ExpectedSum: 4200,
		//		AuthType:    AuthTypeIncorrect,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		Parent:      1,
		//	},
		//	&Pass{
		//		PaymentType: PaymentTypeFree,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//		Parent:      1,
		//	},
		//},
	}
)
