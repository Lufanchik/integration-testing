package mt

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/pass"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMetroComplexMT = test.Cases{
	{
		N:        "1.MМ-MМ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePrepayed,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePrepayed,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
		},
	},
}
