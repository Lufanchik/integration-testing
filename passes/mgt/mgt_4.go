package mgt

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesMGTWithEmptyEMV = test.Cases{
	{
		N:                    "1. MGT EMPTY EMV",
		SkipIdempotencyCheck: true,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeIncorrect,
				Carrier:     carriers.Carrier_MGT,
				EmptyEMV:    true,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				Equal:       1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      3,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
		},
	},
}
