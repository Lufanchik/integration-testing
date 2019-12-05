package mck

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMetroComplexMCK1 = test.Cases{ //gusmanov test case
		{
			N: "1. MCK - MM - MM", //1
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "2. MCK - MM - MCK", //2
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "3. MCK - MM - MMTS - MM", //3
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "4. MCK - MM - MMTS - MCK", //4
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "5. MCK - MM - MMTS - MMTS", //5
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "6. MCK - MM - MMTS - MCD_MSK", //6
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000275", //СЕТУНЬ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
			},
		},
		{
			N: "7. MCK - MM - MMTS - MCD_MO", //7
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2002077", //марк
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000115", //лобня
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "8. MCK - MM - MCD_MSK - MM - MM", //8
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000245", //рабочий поселок
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "9. MCK - MM - MCD_MSK - MM - MCK", //9
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001140", //Кунцевская (бывш. КУНЦЕВО 1)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000455", //ДЕГУНИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "10. MCK - MM - MCD_MSK - MM - MMTS", //10
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType:test. PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2003965", //ТИМИРЯЗЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "11. MCK - MM - MCD_MSK - MM - MCD_MSK", //11
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "12. MCK - MM - MCD_MSK - MM - MCD_MO", //12
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "13. MCK - MM - MCD_MSK - MCK", //13
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "14. MCK - MM - MCD_MSK - MMTS - MM", //14
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:    carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "15. MCK - MM - MCD_MSK - MMTS - MCK", //15
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:    carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "16. MCK - MM - MCD_MSK - MMTS - MMTS", //16
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:    carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "17. MCK - MM - MCD_MSK - MMTS - MCD_MSK", //17
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2003965", //ТИМИРЯЗЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "18. MCK - MM - MCD_MSK - MMTS - MCD_MO", //18
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001240", //БЕСКУДНИКОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2001775", //ШЕРЕМЕТЬЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "19. MCK - MM - MCD_MSK1 - MCD_MSK2 - MM", //19
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000235", //ДМИТРОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001425", //КРАСНЫЙ БАЛТИЕЦ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "20. MCK - MM - MCD_MSK1 - MCD_MSK2 - MCK", //20
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000235", //ДМИТРОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001010", //Новохохловская
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "21. MCK - MM - MCD_MSK1 - MCD_MSK2 - MMTS", //21
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000235", //ДМИТРОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001010", //Новохохловская
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "22. MCK - MM - MCD_MSK1 - MCD_MSK2 - MCD_MSK1", //22
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000235", //ДМИТРОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001010", //Новохохловская
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000006", //Белорусская (бывш. Белорусский вокзал)
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},

		{
			N: "23. MCK - MM - MCD_MSK1 - MCD_MSK2 - MCD_MO", //23
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000235", //ДМИТРОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001010", //Новохохловская
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2001775", //ШЕРЕМЕТЬЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "24. MCK - MM - MCD_MSK1 - MCD_MSK1", //24
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2002077", //МАРК
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},
	}