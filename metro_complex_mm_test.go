package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestMetroComplexMM(t *testing.T) {
	Run(t, casesMetroComplexMM)
}

// "MM - MM" (Две платные поездки)
var (
	casesMetroComplexMM = Cases{
		{
			N: "1. MМ - MМ",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		//{
		//	N: "2. MM - MCK - MМ", //(Бесплатная пересадка c MM на МЦК, взимание денежных средств за пересадку с МЦК на МЦК)
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//	},
		//},
		//{
		//	N: "3. MM - MЦK - MЦК	", //(Бесплатная пересадка c MM на МЦК, взимание денежных средств за пересадку с МЦК на МЦК)
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//	},
		//},
		//{
		//	N: "4. MM - MЦK - MМТС - ММ	", //(Бесплатная пересадка c MM на МЦК на ММТС
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "5. MM - МЦК - MMTS -MЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "6. MM - МЦК - MMTS - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "7. MM - МЦК - MMTS - MCK-МСК", //(Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на МЦК (ММТС может быть только в начале или конце))
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000185", //ЛИАНОЗОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//	},
		//},
		//{
		//	N: "8. MM - МЦК - MMTS - MCK-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001340", //ВОДНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "9. MM - МЦК - MСК-МСК - ММ - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "10. MM - МЦК - MСК-МСК - ММ - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "11. MM - МЦК - MСК-МСК - ММ - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "12. MM - МЦК - MСК-МСК - ММ - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000006", //БЕЛОРУССКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001240", //БЕСКУДНИКОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//		},
		//	},
		//},
		//{
		//	N: "13. MM - МЦК - MСК-МСК - ММ - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000006", //БЕЛОРУССКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000055", //БЕСКУДНИКОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "14. MM - МЦК - MСК-МСК - ММ - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "15. MM - МЦК - MСК-МСК - ММТС - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "16. MM - МЦК - MСК-МСК - ММТС - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "17. MM - МЦК - MСК-МСК - ММТС - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "18. MM - МЦК - MСК-МСК - ММТС - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//		},
		//	},
		//},
		//{
		//	N: "19. MM - МЦК - MСК-МСК - ММТС - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "20. MM - МЦК - MСК-МСК - МСК-МСК2 - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum:      4200,
		//		},
		//	},
		//},
		//{
		//	N: "21. MM - МЦК - MСК-МСК - МСК-МСК2 - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum:      4200,
		//		},
		//	},
		//},
		//{
		//	N: "22. MM - МЦК - MСК-МСК - МСК-МСК2 - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001410", //ГРАЖДАНСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000045", //ТЕКСТИЛЬЩИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum:      4200,
		//		},
		//	},
		//},
		//{
		//	N: "23. MM - МЦК - MСК-МСК - МСК-МСК2 - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001410", //ГРАЖДАНСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000045", //ТЕКСТИЛЬЩИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001410", //ГРАЖДАНСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000045", //ТЕКСТИЛЬЩИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "24. MM - МЦК - MСК-МСК - МСК-МСК2 - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001410", //ГРАЖДАНСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000045", //ТЕКСТИЛЬЩИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000009", //САВЕЛОВСКИЙ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000620", //ХЛЕБНИКОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "25. MM - МЦК - MСК-МСК - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 	 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//	},
		//},
		//{
		//	N: "26. MM - МЦК - MСК-МСК - МСК-МО2 - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001194", //АНИКЕЕВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "27. MM - МЦК - MСК-МСК - МСК-МО2 - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001194", //АНИКЕЕВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "28. MM - МЦК - MСК-МСК - МСК-МО2 - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001194", //АНИКЕЕВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "29. MM - МЦК - MСК-МСК - МСК-МО2 - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001194", //АНИКЕЕВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "30. MM - МЦК - MСК-МСК - МСК-МО2 - МСК-МО2",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001194", //АНИКЕЕВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "31. MM - МЦК - MСК-МО - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000155", //ФИЛИ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "32. MM - МЦК - MСК-МО - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000185", //ЛИАНОЗОВО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000600", //НОВОДАЧНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "33. MM - МЦК - MСК-МО - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000185", //ЛИАНОЗОВО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000600", //НОВОДАЧНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "34. MM - МЦК - MСК-МО - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000185", //ЛИАНОЗОВО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000600", //НОВОДАЧНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//	},
		//},
		//{
		//	N: "35. MM - МЦК - MСК-МО - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000185", //ЛИАНОЗОВО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000600", //НОВОДАЧНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000100", //ДОЛГОПРУДНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "36. MM - ММТС - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//	},
		//},
		//{
		//	N: "37. MM - ММТС - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "38. MM - ММТС - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "39. MM - ММТС - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000275", //СЕТУНЬ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//		},
		//	},
		//},
		//{
		//	N: "40. MM - ММТС - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000100", //ДОЛГОПРУДНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "41. MM - MCK-МСК - MM - MM",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "42. MM - MCK-МСК - MM - MM",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "43. MM - MCK-МСК - MM - MЦК - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "44. MM - MCK-МСК - MM - MЦК - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "45. MM - MCK-МСК - MM - MЦК - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//		},
		//	},
		//},
		//{
		//	N: "46. MM - MCK-МСК - MM - MЦК - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002910", //НЕМЧИНОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "47. MM - MCK-МСК -MM - MМТС - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001250", //МОСКВА ТОВАРНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "48. MM - MCK-МСК -MM - MМТС - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001250", //МОСКВА ТОВАРНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "49. MM - MCK-МСК -MM - MМТС - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001250", //МОСКВА ТОВАРНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "50. MM - MCK-МСК -MM - MМТС - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001250", //МОСКВА ТОВАРНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001250", //МОСКВА ТОВАРНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//		},
		//	},
		//},
		//{
		//	N: "51. MM - MCK-МСК -MM - MМТС - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001250", //МОСКВА ТОВАРНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000720", //ОПАЛИХА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "52. MM - MCK-МСК -MM - МСК-МСК2 - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002077", //МАРК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000720", //ОПАЛИХА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "53. MM - MCK-МСК -MM - МСК-МСК2 - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002077", //МАРК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000720", //ОПАЛИХА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "54. MM - MCK-МСК -MM - МСК-МСК2 - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002077", //МАРК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000720", //ОПАЛИХА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "55. MM - MCK-МСК -MM - МСК-МСК2 - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002077", //МАРК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000720", //ОПАЛИХА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002077", //МАРК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//	},
		//},
		//{
		//	N: "56. MM - MCK-МСК -MM - МСК-МСК2 - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002077", //МАРК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000720", //ОПАЛИХА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "57. MM - MCK-МСК - MM - MСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//	},
		//},
		//{
		//	N: "58. MM - MCK-МСК - MM - MСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "59. MM - MCK-МСК - MЦК - MМ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "60. MM - MCK-МСК - MЦК - MЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "61. MM - MCK-МСК - MЦК - MMTS - MМ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "62. MM - MCK-МСК - MЦК - MMTS - MЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "63. MM - MCK-МСК - MЦК - MMTS - MМТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "64. MM - MCK-МСК - MЦК - MMTS - MСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001270", //ОКРУЖНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//		},
		//	},
		//},
		//{
		//	N: "65. MM - MCK-МСК - MЦК - MMTS - MСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "66. MM - MCK-МСК - MЦК - MСК-МСК2 - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "67. MM - MCK-МСК - MЦК - MСК-МСК2 - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "68. MM - MCK-МСК - MЦК - MСК-МСК2 - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001775", //ШЕРЕМЕТЬЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "69. MM - MCK-МСК - MЦК - MСК-МСК2 - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "70. MM - MCK-МСК - MЦК - MСК-МСК2 - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000235", //ДМИТРОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000200", //КАЛИТНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000055", //ОДИНЦОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "71. MM - MCK-МСК - MЦК - MСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//	},
		//},
		//{
		//	N: "72. MM - MCK-МСК - MЦК - MСК-МО2 - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000630", //БУТОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "73. MM - MCK-МСК - MЦК - MСК-МО2 - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000630", //БУТОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "74. MM - MCK-МСК - MЦК - MСК-МО2 - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000630", //БУТОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "75. MM - MCK-МСК - MЦК - MСК-МО2 - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000630", //БУТОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "76. MM - MCK-МСК - MЦК - MСК-МО2 - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000630", //БУТОВО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002910", //НЕМЧИНОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "77. MM - MCK-МСК - MMТС - MM",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "78. MM - MCK-МСК - MMТС - MЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "79. MM - MCK-МСК - MMТС - MМТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "80. MM - MCK-МСК - MMТС - MСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//	},
		//},
		//{
		//	N: "81. MM - MCK-МСК - MMТС - MСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001060", //БЕГОВАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent:      1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 5,
		//		},
		//	},
		//},
		//{
		//	N: "82. MM - MCK-МСК - МСК-МСК2 - MM", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств при пересадке с ММТС на ММ)
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "83. MM - MCK-МСК - МСК-МСК2 - МЦК - MM",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "84. MM - MCK-МСК - МСК-МСК2 - МЦК - MЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "85. MM - MCK-МСК - МСК-МСК2 - МЦК - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "86. MM - MCK-МСК - МСК-МСК2 - МЦК - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "87. MM - MCK-МСК - МСК-МСК2 - МЦК - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2002910", //НЕМЧИНОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "88. MM - MCK-МСК - МСК-МСК2 - ММТС - MM",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "89. MM - MCK-МСК - МСК-МСК2 - ММТС - MЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "90. MM - MCK-МСК - МСК-МСК2 - ММТС - MМТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "91. MM - MCK-МСК - МСК-МСК2 - ММТС - MСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//		},
		//	},
		//},
		//{
		//	N: "92. MM - MCK-МСК - МСК-МСК2 - ММТС - MСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 7,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "93. MM - MCK-МСК - МСК-МСК2 - MСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//		},
		//	},
		//},
		//{
		//	N: "94. MM - MCK-МСК - МСК-МСК2 - MСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001005", //МОСКВА КАЛАНЧЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 6,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "95. MM - MCK-МСК - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 	 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal:    &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//	},
		//},
		//{
		//	N: "96. MM - MCK-МСК - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2002910", //НЕМЧИНОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "97. MM - MCK-МО - ММ",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "98. MM - MCK-МО - МЦК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "99. MM - MCK-МО - ММТС",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			ExpectedSum: 4200,
		//		},
		//	},
		//},
		//{
		//	N: "100. MM - MCK-МО - МСК-МСК",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//		},
		//	},
		//},
		//{
		//	N: "101. MM - MCK-МО - МСК-МО",
		//	T: T{
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent: 1,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000685", //БАКОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 2,
		//			ExpectedSum: 700,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//		},
		//		&Pass{
		//			PaymentType: PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2002910", //НЕМЧИНОВКА
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 4,
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
	}
)
