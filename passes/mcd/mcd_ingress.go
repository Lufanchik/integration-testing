package mcd

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var (
	CasesMcdIngress = test.Cases{
		{
			N: "1.MCD_INGRESS_UNKNOWN_INGRESS",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					ExpectedSum:        4600,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_INGRESS,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					ExpectedSum:        4600,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_INGRESS,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					Ingress:            3,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_EGRESS,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					ExpectedSum:        4600,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_INGRESS,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					Ingress:            5,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_EGRESS,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					ExpectedSum:        4600,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_INGRESS,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_UNKNOWN,
					},
					Ingress:            7,
					IsReplaceDirection: true,
					RealDirection:      processing.TerminalDirection_EGRESS,
				},
			},
		},
		//{
		//	N: "1.MCD_INGRESS_UNKNOWN_EGRESS",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_CHECK_IN_PATH,
		//			},
		//			ExpectedSum: 4600,
		//		},
		//	},
		//},
	}
)
