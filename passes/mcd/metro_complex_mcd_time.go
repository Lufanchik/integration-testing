package mcd

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var (
	CasesComplexTimeMCD = test.Cases{
		//{
		//	N: "1. МО-МСК - ММ",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000055", //ОДИНЦОВО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:         test.NowFullDate(12, 2, 10, 30,00),
		//			ExpectedSum: 4900,
		//			TimeToWait:  35 * time.Minute,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001140", //КУНЦЕВСКАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//			Now:     test.NowFullDate(12, 2, 11, 22,00),
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         test.NowFullDate(12, 2, 12, 35,00),
		//			Parent:      1,
		//		},
		//	},
		//},
		//{
		//	N: "2. МО-МСК2 - ММ",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000065", //ПОДОЛЬСК
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:         test.NowFullDate(12, 2, 10, 30,00),
		//			ExpectedSum: 4900,
		//			TimeToWait:  60 * time.Minute,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2002780", //ДЕПО
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//			Now:     test.NowFullDate(12, 2, 12, 9,00),
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         test.NowFullDate(12, 2, 12, 31,00),
		//			Parent:      1,
		//		},
		//	},
		//},
		//{
		//	N: "3. ММ - МЦК - ММТС -МСК-МО",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         test.NowFullDate(12, 2, 10, 00,00),
		//			ExpectedSum: 4200,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
		//			Now:         test.NowFullDate(12, 2, 11, 31,00),
		//			ExpectedSum: 4200,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Now:         test.NowFullDate(12, 2, 13, 02,00),
		//			ExpectedSum: 4200,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			ExpectedSum: 4200,
		//			Now:         test.NowFullDate(12, 2, 14, 33,00),
		//			TimeToWait:  5 * time.Minute,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress:     4,
		//			Now:         test.NowFullDate(12, 2, 15, 00,00),
		//			ExpectedSum: 700,
		//		},
		//	},
		//},
		//{
		//	N: "4. МО-МСК2 - МСК-МО",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000460", //НАХАБИНО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:         test.NowFullDate(12, 2, 9, 00,00),
		//			ExpectedSum: 4900,
		//			TimeToWait:  60 * time.Minute,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001840", //ТРИКОТАЖНАЯ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//			Now:     test.NowFullDate(12, 2, 10, 23,00),
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000700", //ТЕСТОВСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:        test.NowFullDate(12, 2, 11, 00,00),
		//			TimeToWait: 5 * time.Minute,
		//			Parent:     1,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD1_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001340", //ВОДНИКИ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 3,
		//			Now:     test.NowFullDate(12, 2, 11, 55,00),
		//		},
		//	},
		//},
		//{
		//	N: "5. МО-МСК2 - МСК-МО2", //Поездка на границе нового года, месяца, дня
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000460", //НАХАБИНО
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Now:         test.NowFullDate(12, 31, 23, 00,00),
		//			ExpectedSum: 4900,
		//			TimeToWait:  30 * time.Minute,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001425", //КРАСНЫЙ БАЛТИЕЦ
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Ingress: 1,
		//			Now:     test.NowFullDate(13, 01, 24, 8, 00),
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MSK,
		//			Terminal: &processing.Terminal{
		//				Station:   "2001410", //ГРАЖДАНСКАЯ
		//				Direction: processing.TerminalDirection_INGRESS,
		//			},
		//			Parent:     1,
		//			Now:        test.NowFullDate(13, 01, 24, 20,00),
		//			TimeToWait: 5 * time.Minute,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MCD,
		//			SubCarrier:  carriers.SubCarrier_MCD2_MO,
		//			Terminal: &processing.Terminal{
		//				Station:   "2000065", //ПОДОЛЬСК
		//				Direction: processing.TerminalDirection_EGRESS,
		//			},
		//			Now:     test.NowFullDate(13, 01, 25, 43,00),
		//			Ingress: 3,
		//		},
		//	},
		//},
		{
			N: "6. ММ - МЦК - ММТС - МСК-МО - МО-МО2 - МСК-МСК - ММТС - МСК-МСК2 - ММ,",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowFullDate(01, 01, 01, 01, 00),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Now:         test.NowFullDate(01, 01, 02, 32, 00),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(01, 01, 04, 03, 00),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
					Now:         test.NowFullDate(01, 01, 05, 35, 00),
					TimeToWait:  10 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					Now:         test.NowFullDate(01, 01, 06, 07, 00),
					ExpectedSum: 700,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
					Now:         test.NowFullDate(01, 01, 06, 07, 00),
					TimeToWait:  10 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960", //ПАВШИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
					Now:     test.NowFullDate(1, 1, 7, 58, 00),
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
					Now:         test.NowFullDate(1, 1, 8, 30, 00),
					TimeToWait:  10 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 8,
					Now:     test.NowFullDate(1, 1, 8, 53, 00),
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(1, 1, 10, 11, 00),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000605", //РИЖСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:     10,
					Now:        test.NowFullDate(1, 1, 10, 45, 00),
					TimeToWait: 15 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 11,
					Now:     test.NowFullDate(1, 1, 11, 18, 00),
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowFullDate(01, 01, 11, 31, 00),
					Parent:      10,
				},
			},
		},
		{
			N: "7. МО-МСК - МСК-МО2",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
					Now:         test.NowFullDate(12, 31, 23, 45, 00),
					TimeToWait:  30 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Now:     test.NowFullDate(12, 32, 25, 05, 00),
					Ingress: 1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840", //ТРИКОТАЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:     1,
					Now:        test.NowFullDate(12, 32, 25, 25, 15),
					TimeToWait: 5 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Now:     test.NowFullDate(12, 32, 25, 53, 59),
					Ingress: 3,
				},
			},
		},
	}
)
