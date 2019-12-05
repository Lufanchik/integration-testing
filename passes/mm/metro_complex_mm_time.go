package mm

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var (
	CasesComplexTimeMM = test.Cases{
		{
			N: "1. ММ - ММ - ММ",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowFullDate(2019, 12, 05, 12, 25, 00),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowFullDate(2019, 12, 05, 12, 40, 00),
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowFullDate(2019, 12, 05, 13, 10, 00),
					ExpectedSum: 4200,
				},
			},
		},
	}
)
