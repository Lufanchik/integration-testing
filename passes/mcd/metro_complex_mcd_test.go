package mcd_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing"
	"testing"
)

func TestComplexMCD(t *testing.T) {
	http.Run(t, casesComplexMCD)
}

var (
	casesComplexMCD = http.Cases{
		{
			N: "mm-mcd1-mo",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000600",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
				},
			},
		},
		{
			N: "MM-MCD1MO-MCD2",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2002952",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2002910",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},
		{
			N: "MM-MO-MO",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000285", //ЩЕРБИНКА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM-MO-MSK",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MO-MSK",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
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
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MSK-MO",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
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
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
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
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
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
			N: "ММ - МСК - ММ - МЦД МСК 2",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCD MO - MM - MM",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ - МЦК -	ММ - MM",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - ММ - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - ММ - МMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - ММ - МЦД МО",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS - MM",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS - MCK",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS - MMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MCD MO - MM - MM",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		//{
		//	N: "MCD MSK1 - MCD MSK2 - MCD MO2",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000275",
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270",
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000075",
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000460",
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress:     5,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		{
			N: "МЦД МО - ММ - МЦК -	ММ - MM",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - ММ - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - ММ - МMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1  - ММ -  МЦК - ММ - МЦД МО 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS - MM",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS - MCK",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО - ММ -  МЦК - МMTS - MMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО1 - ММ -  МЦК - МMTS - MЦД МСК1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000275",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002077",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "МЦД МО1 - ММ -  МЦК - МMTS - MЦД МО1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000275",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "МЦД МО1 - ММ -  МЦК - МЦД МСК 2 - ММ",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО1 - ММ -  МЦК - МЦД МСК 2 - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО1 - ММ -  МЦК - МЦД МСК 2 - МMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840", //ТРИКОТАЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780", //ДЕПО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ -  МЦК - МЦД МСК 2 - МЦД МСК 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000455", //ДЕГУНИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
		{
			N: "МЦД МО/МСК 1 - ММ -  МЦК - МЦД МСК/МСК 2 - МЦД МСК/МО 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000225", //ЦАРИЦЫНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					ExpectedSum: 700,
					Ingress:     7,
				},
			},
		},
		{
			N: "МЦД МО/МСК 1 - ММ -  МЦК - МЦД МСК/МО 2 - ММ",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО/МСК 1 - ММ -  МЦК - МЦД МСК/МО 2 - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО/МСК 1 - ММ -  МЦК - МЦД МСК/МО 2 - ММТС",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},

				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО/МСК 1 - ММ -  МЦК - МЦД МСК/МO 2 - МЦД МЦК 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000275", //СЕТУНЬ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000455", //ДЕГУНИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
		{
			N: "МЦД МО/МСК 1 - ММ -  МЦК - МЦД МСК/МO 2 - МЦД МСК/МО 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000275", //СЕТУНЬ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ -  ММТС - ММ",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ -  ММТС - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ -  ММТС - ММТС",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ -  ММТС - МЦД МЦК 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ -  ММТС - МЦД МО 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -ММ - ММ",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -ММ - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -ММ - МMTS",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -ММ - МЦД МСК 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -ММ - МЦД МО 1",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -МЦК - ММ",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -МЦК - МЦК",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "МЦД МО 1 - ММ - МЦД МСК 2 -МЦК - ММТС",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000075", //ТУШИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
	}
)
