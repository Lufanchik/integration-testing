package mck

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMetroComplexMCK3 = test.Cases{ //gusmanov test case
		{
			N: "51. MCK -  MCD1_MSK - MM - MCD2_MSK - MCD1_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
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
					Ingress: 7,
				},
			},
		},

		{
			N: "52. MCK -  MCD1_MSK - MM - MCD2_MSK - MCD1_MO",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000455", //ДЕГУНИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					ExpectedSum: 700,
					Ingress: 7,
				},
			},
		},

		{
			N: "53. MCK -  MCD1_MSK - MM - MCD1_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
						Station: "2002077", //МАРК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},

		{
			N: "54. MCK - MCD_MSK - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
						Station: "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
			N: "55. MCK - MCD_MSK - MMTC - MM",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
						Station: "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "56. MCK - MCD_MSK - MMTC - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "57. MCK - MCD_MSK - MMTC - MMTC",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
			N: "58. MCK - MCD_MSK - MMTC - MCD_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
					Ingress: 5,
				},
			},
		},

		{
			N: "59. MCK - MCD1_MSK - MMTC - MCD1_MO",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
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
					Ingress: 2,
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
						Station: "2001060", //БЕГОВАЯ
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
					Ingress: 5,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "60. MCK - MCD1_MSK - MCD2_MSK - MM - MM",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "61. MCK - MCD1_MSK - MCD2_MSK - MM - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "62. MCK - MCD1_MSK - MCD2_MSK - MM - MMTS",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "63. MCK - MCD1_MSK - MCD2_MSK - MM - MCD1_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
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
			N: "64. MCK - MCD1_MSK - MCD2_MSK - MM - MCD1_MO",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
						Station: "2001060", //БЕГОВАЯ
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
			N: "65. MCK - MCD1_MSK - MCD2_MSK - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000275", //СЕТУНЬ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2003965", //ТИМИРЯЗЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001410", //ГРАЖДАНСКАЯ
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
					Ingress: 4,
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
			N: "66. MCK - MCD1_MSK - MCD2_MSK - MMTS - MM",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "67. MCK - MCD1_MSK - MCD2_MSK - MMTS - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2002780", //ДЕПО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000045", //ТЕКСТИЛЬЩИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "68. MCK - MCD1_MSK - MCD2_MSK - MMTS - MMTS",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
						Station: "2000630", //БУТОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "69. MCK - MCD1_MSK - MCD2_MSK - MMTS - MCD1_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
						Station: "2000630", //БУТОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
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
			N: "70. MCK - MCD1_MSK - MCD2_MSK - MMTS - MCD1_MO",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
						Station: "2000630", //БУТОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
						Station: "2002077", //МАРК
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
					Ingress: 7,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "71. MCK - MCD1_MSK - MCD2_MSK - MCD1_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
						Station: "2000630", //БУТОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},

		{
			N: "72. MCK - MCD1_MSK - MCD2_MSK - MCD2_MSK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
						Station: "2000630", //БУТОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},

		{
			N: "73. MCK - MCD1_MSK - MCD2_MO - MM",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000225", //ЦАРИЦЫНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station: "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
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
			N: "74. MCK - MCD1_MSK - MCD2_MO - MCK",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000225", //ЦАРИЦЫНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station: "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
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
			N: "75. MCK - MCD1_MSK - MCD2_MO - MMTS",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000225", //ЦАРИЦЫНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station: "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
	}

