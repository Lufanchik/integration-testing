package mcd_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing"
	"testing"
)

func TestMetroComplexMCDMSK2(t *testing.T) {
	http.Run(t, casesMetroComplexMCDMSK2)
}

var (
	casesMetroComplexMCDMSK2 = http.Cases{
		{
			N: "26. МСК-МСК - ММ - ММТС - МСК-МСК",
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},
		{
			N: "27. МСК-МСК - ММ - ММТС - МСК-МО",
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},
		{
			N: "28. МСК-МСК - ММ - МСК-МСК2 - ММ - ММ",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "29. МСК-МСК - ММ - МСК-МСК2 - ММ - МЦК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "30. МСК-МСК - ММ - МСК-МСК2 - ММ - ММТС",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "31. МСК-МСК - ММ - МСК-МСК2 - ММ - МСК-МСК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "32. МСК-МСК - ММ - МСК-МСК2 - ММ - МСК-МО",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "33. МСК-МСК - ММ - МСК-МСК2 - МЦК - ММ",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "34. МСК-МСК - ММ - МСК-МСК2 - МЦК - МЦК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "35. МСК-МСК - ММ - МСК-МСК2 - МЦК - ММТС",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "36. МСК-МСК - ММ - МСК-МСК2 - МЦК - МСК-МСК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
		{
			N: "37. МСК-МСК - ММ - МСК-МСК2 - МЦК - МСК-МО",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "39. МСК-МСК - ММ - МСК-МСК2 - ММТС - МЦК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "40. МСК-МСК - ММ - МСК-МСК2 - ММТС - ММТС",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "41. МСК-МСК - ММ - МСК-МСК2 - ММТС - МСК-МСК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
			N: "42. МСК-МСК - ММ - МСК-МСК2 - ММТС - МСК-МО",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "43. МСК-МСК - ММ - МСК-МСК2 - МСК-МСК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2000185", //ЛИАНОЗОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},
		{
			N: "44. МСК-МСК - ММ - МСК-МСК2 - МСК-МО",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "45. МСК-МСК - ММ - МСК-МО2 - ММ",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2002952", //СИЛИКАТНАЯ
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
			N: "46. МСК-МСК - ММ - МСК-МО2 - МЦК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2002952", //СИЛИКАТНАЯ
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
			N: "47. МСК-МСК - ММ - МСК-МО2 - ММТС",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2002952", //СИЛИКАТНАЯ
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
			N: "48. МСК-МСК - ММ - МСК-МО2 - МСК-МСК",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2002952", //СИЛИКАТНАЯ
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
			N: "49. МСК-МСК - ММ - МСК-МО2 - МСК-МО",
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
					SubCarrier:  carriers.SubCarrier_MM_SUB,
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
						Station:   "2002952", //СИЛИКАТНАЯ
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
						Station:   "2000600", //НОВОДАЧНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "50. МСК-МСК - МЦК - ММ - ММ",
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
	}
)
