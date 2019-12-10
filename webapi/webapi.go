package webapi

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var (
	CasesWEBAPI = test.Cases{
		{
			N: "Web API test",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    test.AuthTypeIncorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&test.WebAPICheck{
					Passes: []int{1, 2},
				},
			},
		},
	}
)
