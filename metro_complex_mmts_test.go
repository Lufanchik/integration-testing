package http_test

import (
	"testing"
)

func TestComplexPassMMTSMMTS(t *testing.T) {
	Run(t, casesComplexPassMMTSMMTS)
}

var (
	// MMTS - MMTS
	casesComplexPassMMTSMMTS = Cases{

		{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
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
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
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
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
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
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      1,
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
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      1,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		Parent:      1,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		RequestType: RequestTypeOnline,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      1,
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
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      1,
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
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      1,
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
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      1,
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
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//},
			//
			////  MMTS - MMTS (not money) - MCK (если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//		AuthType:    AuthTypeIncorrect,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      2,
			//	},
			//},
			//// Кейсы включающие транзакции с некорректной авторизацией
			//// "MMTS - MMTS (AuthTypeIncorrect) - MM - MCK -MM" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      2,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		Parent:      2,
			//	},
			//},
			//// "MMTS - MMTS (AuthTypeIncorrect) - MM - MCK -MM" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			//
			//{
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//		ExpectedSum: 4200,
			//		AuthType:    AuthTypeIncorrect,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//		Parent:      2,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFree,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		Parent:      2,
			//	},
			//	&Pass{
			//		PaymentType: PaymentTypeFullPayment,
			//		Carrier:     carriers.Carrier_MM,
			//		SubCarrier:  carriers.SubCarrier_MM_SUB,
			//		ExpectedSum: 4200,
			//	},
		},
	}
)
