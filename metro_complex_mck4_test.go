package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

var (
	casesComplexMCK4 = Cases{ //gusmanov test case
		{
			N: "76. MCK - MCD1_MSK - MCD2_MO - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000225", //ЦАРИЦЫНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station: "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2002077", //МАРК
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
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
			N: "77. MCK - MCD1_MSK - MCD2_MO - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station: "2000225", //ЦАРИЦЫНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station: "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2002077", //МАРК
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
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
			N: "78. MCK - MCD1_MO - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: 	carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "79. MCK - MCD1_MO - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000100", //ДОЛГОПРУДНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: 	carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "80. MCK - MCD1_MO - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: 	carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "81. MCK - MCD1_MO - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2002077", //МАРК
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
			},
		},

		{
			N: "82. MCK - MCD1_MO - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:	 carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station: "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
					ExpectedSum: 700,
				},
			},
		},
	}
)