package mcd_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing"
	"testing"
	"time"
)

func TestComplex2MCD(t *testing.T) {
	http.Run(t, casesComplexMCD2)
}

var (
	casesComplexMCD2 = http.Cases{
		{
			N: "98. ММТС - МСК-МСК - ММ - МСК-МО - ММ",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001140", //КУНЦЕВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     5,
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
			N: "MCD1-fail-egress-MCD2",
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
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000235",
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460",
						Direction: processing.TerminalDirection_EGRESS,
					},
					ExpectedSum: 700,
					Ingress:     2,
				},
			},
		},
		{
			N: "MCD1-mo-msk-MCD2-msk-mo",
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
						Station:   "2001060",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000235",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
			},
		},
		{
			N: "MCD1-mo-msk-MCD2-msk-mo",
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
						Station:   "2001060",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000235",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
			},
		},
		{
			N: "76. MM - MCK-МСК - MЦК - MСК-МО2 - МСК-МО",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 2,
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
						Station:   "2001120", //МОСКВОРЕЧЬЕ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     5,
					ExpectedSum: 700,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2002910", //НЕМЧИНОВКА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "MCD egress after complex period",
			T: http.T{
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					Now:         http.NowCustom(10, 00),
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:     1,
					Now:        http.NowCustom(10, 30),
					TimeToWait: 95 * time.Minute,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_EGRESS,
					},
					ExpectedSum: 700,
					Ingress:     2,
					Now:         http.NowCustom(12, 30),
				},
			},
		},
	}
)
