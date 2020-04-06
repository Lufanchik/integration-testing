package face

import (
	"github.com/google/uuid"
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var MetroComplexFaceID = test.Cases{
	{
		N:          "1. Face/MM-MCK-MMTS",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
		},
	},
	{
		N:          "2. Face/MMTS-MCK-MM",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
		},
	},
	{
		N:          "3. Face/MCK-MM-MMTS",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
		},
	},
	{
		N:          "4. Face/MCD_MO-MCD_MSK-MM-MCK-MMTS",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000185", //ЛИАНОЗОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      3,
			},
		},
	},
	{
		N:          "5. Face/MCD_MO-MCD_MO-MMTS-MCK",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000600", //
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 3,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      9,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      9,
			},
		},
	},
	{
		N:          "6. Face/MGT-MMTS-MCK-MGT",
		CardSystem: processing.CardSystem_VISA,
		FaceId:     uuid.New().String(),
		T: test.T{
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      6,
			},
			&test.RegisterFaceId{},
			&test.FaceIdRegistrationStatus{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
		},
	},
}
