package mcd_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing"
	"testing"
)

func TestMetroComplexMCDMSK4(t *testing.T) {
	http.Run(t, casesMetroComplexMCDMSK4)
}

var (
	casesMetroComplexMCDMSK4 = http.Cases{
		{
			N: "76. МСК-МСК - МЦК - МСК-МСК2 - ММ - МСК-МСК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
		{
			N: "77. МСК-МСК - МЦК - МСК-МСК2 - ММ - МСК-МО",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
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
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000685", //БАКОВКА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "78. МСК-МСК - МЦК - МСК-МСК2 - МЦК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "79. МСК-МСК - МЦК - МСК-МСК2 - ММТС - ММ",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "80. МСК-МСК - МЦК - МСК-МСК2 - ММТС - МЦК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "81. МСК-МСК - МЦК - МСК-МСК2 - ММТС - ММТС",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "82. МСК-МСК - МЦК - МСК-МСК2 - ММТС - МСК-МСК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
		{
			N: "83. МСК-МСК - МЦК - МСК-МСК2 - ММТС - МСК-МО",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
			N: "84. МСК-МСК - МЦК - МСК-МСК2 - МСК-МСК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
						Station:   "2000275", //СЕТУНЬ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "85. МСК-МСК - МЦК - МСК-МСК2 - МСК-МО",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
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
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "86. МСК-МСК - МЦК - МСК-МО2 - ММ",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
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
			N: "87. МСК-МСК - МЦК - МСК-МО2 - МЦК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
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
			N: "88. МСК-МСК - МЦК - МСК-МО2 - ММТС",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
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
			N: "89. МСК-МСК - МЦК - МСК-МО2 - МСК-МСК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "90. МСК-МСК - МЦК - МСК-МО2 - МСК-МО",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
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
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "91. МСК-МСК - ММТС - ММ",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
			N: "92. МСК-МСК - ММТС - МЦК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
			N: "93. МСК-МСК - ММТС - ММТС",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
			N: "94. МСК-МСК - ММТС - МСК-МСК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 4,
				},
			},
		},
		{
			N: "95. МСК-МСК - ММТС - МСК-МО",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
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
					PaymentType: http.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "96. МСК-МСК - МСК-МСК2 - ММ - ММ",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
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
			N: "97. МСК-МСК - МСК-МСК2 - ММ - МЦК - ММ",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "98. МСК-МСК - МСК-МСК2 - ММ - МЦК - МЦК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
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
			N: "99. МСК-МСК - МСК-МСК2 - ММ - МЦК - ММТС",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
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
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "100. МСК-МСК - МСК-МСК2 - ММ - МЦК - МСК-МСК",
			T: http.T{
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001250", //МОСКВА ТОВАРНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&http.Pass{
					PaymentType: http.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000595", //ПОКРОВСКОЕ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},
	}
)
