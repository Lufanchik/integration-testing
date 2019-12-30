package face

import (
	"github.com/google/uuid"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesAuthWithFace = test.Cases{
	{
		N:          "Face Auth Test",
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
			&test.WebAPIPasses{
				Passes: []int{2},
			},
		},
	},
}
