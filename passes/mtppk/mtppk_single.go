package mtppk

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMTPPK_single = test.Cases{
	{
		N:          "99. Just one pass",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 1000,
			},
			&test.Complete{
				StartPass: 1,
				Sum:       1000,
			},
		},
	},
}
