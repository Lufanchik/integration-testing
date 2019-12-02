package mcd

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var (
	CasesComplexTimeMCD = test.Cases{
		{
			N: "1. MCD egress after complex period",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055", //ОДИНЦОВО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         test.NowCustom(10, 30),
					ExpectedSum: 4900,
					TimeToWait:  35 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001140",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     test.NowCustom(11, 22),
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowCustom(12, 35),
					Parent:      1,
				},
			},
		},
		{
			N: "2. МО-МСК2 - ММ",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         test.NowCustom(10, 30),
					ExpectedSum: 4900,
					TimeToWait:  60 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780", //ДЕПО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     test.NowCustom(12, 9),
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowCustom(12, 31),
					Parent:      1,
				},
			},
		},
		{
			N: "3. ММ - МЦК - ММТС -МСК-МО",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowCustom(10, 00),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Now:         test.NowCustom(11, 31),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowCustom(13, 02),
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
					Now:         test.NowCustom(14, 33),
					TimeToWait:  5 * time.Minute,
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
					Now:         test.NowCustom(15, 00),
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "4. МО-МСК2 - МСК-МО",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         test.NowCustom(9, 00),
					ExpectedSum: 4900,
					TimeToWait:  60 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840", //ТРИКОТАЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     test.NowCustom(10, 23),
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:        test.NowCustom(11, 00),
					TimeToWait: 5 * time.Minute,
					Parent:     1,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
					Now:     test.NowCustom(11, 55),
				},
			},
		},
		{
			N: "5. МО-МСК2 - МСК-МО2", //Поездка на границе нового года, месяца, дня
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         test.NowCustomDate(12, 31, 23, 00),
					ExpectedSum: 4900,
					TimeToWait:  30 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001425", //КРАСНЫЙ БАЛТИЕЦ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     test.NowCustomDate(13, 01, 24, 8),
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:     1,
					Now:        test.NowCustomDate(13, 01, 24, 20),
					TimeToWait: 5 * time.Minute,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Now:     test.NowCustomDate(13, 01, 25, 43),
					Ingress: 3,
				},
			},
		},
		{
			N: "6. ММ - МЦК - ММТС -МСК-МО - МО-МО2 - МСК-МСК - ММТС - МСК-МСК2 - ММ",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowCustomDate(01, 01, 01, 01),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Now:         test.NowCustomDate(01, 01, 02, 32),
					ExpectedSum: 4200,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowCustomDate(01, 01, 04, 03),
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
					Now:         test.NowCustomDate(01, 01, 05, 35),
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
					Now:         test.NowCustomDate(01, 01, 06, 07),
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
					Now:         test.NowCustomDate(01, 01, 06, 07),
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
					Now:     test.NowCustomDate(1, 1, 7, 58),
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
					Now:         test.NowCustomDate(1, 1, 8, 30),
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
					Now:     test.NowCustomDate(1, 1, 8, 53),
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowCustomDate(1, 1, 10, 11),
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
					Now:        test.NowCustomDate(1, 1, 10, 45),
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
					Now:     test.NowCustomDate(1, 1, 11, 18),
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         test.NowCustomDate(01, 01, 11, 31),
					Parent:      10,
				},
			},
		},
	}
)
