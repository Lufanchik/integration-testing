package mmts

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	_ "lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	_ "time"
)

var CasesComplexTimeMMTS = test.Cases{
	{
		N: "1. MMTS - MMTS",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Now:              test.NowFullDate(2019, 12, 31, 23, 59, 00),
				ExpectedSum:      4200,
				IsComplexTimeout: false,

			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Now:              test.NowFullDate(2020, 01, 01, 00, 30, 00),
				ExpectedSum:      4200,
				IsComplexTimeout: false,

			},
		},
	},

}
