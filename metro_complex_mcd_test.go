package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

func TestComplexMCD(t *testing.T) {
	Run(t, casesComplexMCD)
}

var (
	casesComplexMCD = Cases{
		{
			N: "MM-MO-MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2002952",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      4,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM-MO-MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      4,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MO-MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
			},
		},
		{
			N: "MM-MSK-MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MSK-MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_EGRESS,
					},
					ExpectedSum: 700,
					Ingress:     1,
				},
			},
		},
		{
			N: "MO-MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
			},
		},
		{
			N: "MSK-MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
			},
		},
		{
			N: "COMPLEX-MSK-MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		//gusmanov test case
		{
			N: "MCK - MM - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "MCK - MM - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "MCK - MM - MMTS - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					ExpectedSum: 1,
				},
			},
		},

		{
			N: "MCK - MM - MMTS - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "MCK - MM - MMTS - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "MCK - MM - MMTS - MCD_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MCD,
					SubCarrier: carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000275", //СЕТУНЬ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MCD,
					SubCarrier: carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
			},
		},

		{
			N: "MCK - MM - MMTS - MCD_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MMTS_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MCD,
					SubCarrier: carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2002077", //марк
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MCD,
					SubCarrier: carriers.SubCarrier_MCD1_MO,
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
			N: "MCK - MM - MCD_MSK - MM - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MCD,
					SubCarrier: carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2000245", //рабочий поселок
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MCD,
					SubCarrier: carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station: "2001270", //ОКРУЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier: carriers.Carrier_MM,
					SubCarrier: carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

	}
)
