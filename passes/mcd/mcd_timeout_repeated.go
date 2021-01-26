package mcd

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var (
	CasesMcdTimeout = test.Cases{
		{
			N: "1_mcd_timeout_repeated_passes",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 5100,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 5100,
				},
			},
		},
	}
)
