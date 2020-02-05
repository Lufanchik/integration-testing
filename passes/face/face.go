package face

import (
	"github.com/google/uuid"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesAuthWithFace = test.Cases{
	//{
	//	N:          "1. Face Auth Test / mgt",
	//	CardSystem: processing.CardSystem_VISA,
	//	FaceId:     uuid.New().String(),
	//	T: test.T{
	//		&test.RegisterFaceId{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			Carrier:     carriers.Carrier_MGT,
	//			ExpectedSum: 4200,
	//		},
	//		&test.FacePass{
	//			CreateLinks: []int{1},
	//			Passes:      []int{2},
	//		},
	//	},
	//},

	// Фейс работает только на ММ
	{
		N:          "2. Face Auth Test / mm",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	//{
	//	N:          "3. Face Auth Test / mck",
	//	CardSystem: processing.CardSystem_VISA,
	//	FaceId:     uuid.New().String(),
	//	T: test.T{
	//		&test.RegisterFaceId{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
	//			ExpectedSum: 4200,
	//		},
	//		&test.FacePass{
	//			CreateLinks: []int{1},
	//			Passes:      []int{2},
	//		},
	//	},
	//},
	//{
	//	N:          "4. Face Auth Test / mmts",
	//	CardSystem: processing.CardSystem_VISA,
	//	FaceId:     uuid.New().String(),
	//	T: test.T{
	//		&test.RegisterFaceId{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
	//			ExpectedSum: 4200,
	//		},
	//		&test.FacePass{
	//			CreateLinks: []int{1},
	//			Passes:      []int{2},
	//		},
	//	},
	//},
	//{
	//	N:          "4. Face Auth Test / mcd_mo",
	//	CardSystem: processing.CardSystem_VISA,
	//	FaceId:     uuid.New().String(),
	//	T: test.T{
	//		&test.RegisterFaceId{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			Carrier:     carriers.Carrier_MCD,
	//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
	//			Terminal: &processing.Terminal{
	//				Station:   "2000055", //ОДИНЦОВО
	//				Direction: processing.TerminalDirection_INGRESS,
	//			},
	//			ExpectedSum: 4900,
	//		},
	//		&test.FacePass{
	//			CreateLinks: []int{1},
	//			Passes:      []int{2},
	//		},
	//	},
	//},
}
