package webapi

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var (
	CasesWEBAPI = test.Cases{
		{
			N: "1. ММ - ММ",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,

					AuthType: test.AuthTypeCorrect,
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

					AuthType: test.AuthTypeIncorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2},
				},
			},
		},
		{
			N: "2. МЦК - МЦК",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,

					AuthType: test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MCK_SUB,
					},
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,

					AuthType: test.AuthTypeIncorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MCK_SUB,
					},
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2},
				},
			},
		},
		{
			N: "3. ММТС - ММТС",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,

					AuthType: test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MMTS_SUB,
					},
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,

					AuthType: test.AuthTypeIncorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MMTS_SUB,
					},
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2},
				},
			},
		},
		{
			N: "4. МО-МСК1 - ММТС",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 5100,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
					AuthType:    test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MMTS_SUB,
					},
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3},
				},
			},
		},
		{
			N: "5. МО-МО2 - МЦК",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 5100,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000600", //НОВОДАЧНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,

					AuthType: test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MCK_SUB,
					},
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3},
				},
			},
		},
		{
			N: "6. МСК-МСК1 - ММ",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4600,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
					AuthType:    test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3},
				},
			},
		},
		{
			N: "7. МСК-МСК - МЦК - МСК-МСК2 - ММ - МСК-МСК",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4600,
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
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4600,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3, 4, 5, 6, 7, 8},
				},
			},
		},
		{
			N: "8. MЦК - МСК-МСК2 - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000235", //ДМИТРОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780", //ДЕПО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3, 4},
				},
			},
		},
		{
			N: "9. МГТ - МО-МСК1 - МГТ - МЦК - МГТ",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MGT,
					//ExpectedSum: 4600,
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
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MGT,
					//ExpectedSum: 4600,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MGT,
					//ExpectedSum: 4600,
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3, 4, 5, 6},
				},
			},
		},
		{
			N: "10. ММТС - МЦК - МЦК",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
				},
				&test.WebAPIPasses{
					Passes: []int{1, 2, 3},
				},
			},
		},
	}
)
