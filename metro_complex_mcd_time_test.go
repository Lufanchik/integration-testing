package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
	"time"
)

func TestComplexTimeMCD(t *testing.T) {
	Run(t, casesComplexTimeMCD)
}

var (
	casesComplexTimeMCD = Cases{
		{
			N: "1. MCD egress after complex period",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000055",
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         NowCustom(10, 30),
					ExpectedSum: 4900,
					TimeToWait:  35 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001140",
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     NowCustom(11, 22),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         NowCustom(12, 35),
					Parent:      1,
				},
			},
		},
		{
			N: "2. МО-МСК2 - ММ",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         NowCustom(10, 30),
					ExpectedSum: 4900,
					TimeToWait:  60 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2002780", //ДЕПО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     NowCustom(12, 9),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         NowCustom(12, 31),
					Parent:      1,
				},
			},
		},
		{
			N: "3. ММ - МЦК - ММТС -МСК-МО",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         NowCustom(10, 00),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Now:         NowCustom(11, 31),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         NowCustom(13, 02),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
					Now:         NowCustom(14, 33),
					TimeToWait:  5 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					Now:         NowCustom(15, 00),
					ExpectedSum: 700,
				},
			},
		},
		{
			N: "4. МО-МСК2 - МСК-МО",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         NowCustom(9, 00),
					ExpectedSum: 4900,
					TimeToWait:  60 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840", //ТРИКОТАЖНАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     NowCustom(10, 23),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:        NowCustom(11, 00),
					TimeToWait: 5 * time.Minute,
					Parent:     1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
					Now:     NowCustom(11, 55),
				},
			},
		},
		{
			N: "5. МО-МСК2 - МСК-МО2", //Поездка на границе нового года, месяца, дня
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_INGRESS,
					},
					Now:         NowCustomDate(12, 31, 23, 00),
					ExpectedSum: 4900,
					TimeToWait:  30 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001425", //КРАСНЫЙ БАЛТИЕЦ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
					Now:     NowCustomDate(13, 01, 24, 8),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:     1,
					Now:        NowCustomDate(13, 01, 24, 20),
					TimeToWait: 5 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000065", //ПОДОЛЬСК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Now:     NowCustomDate(13, 01, 25, 43),
					Ingress: 3,
				},
			},
		},
		{
			N: "6. ММ - МЦК - ММТС -МСК-МО - МО-МО2 - МСК-МСК - ММТС - МСК-МСК2 - ММ",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         NowCustomDate(01, 01, 01, 01),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Now:         NowCustomDate(01, 01, 02, 32),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         NowCustomDate(01, 01, 04, 03),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
					Now:         NowCustomDate(01, 01, 05, 35),
					TimeToWait:  10 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     4,
					Now:         NowCustomDate(01, 01, 06, 07),
					ExpectedSum: 700,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2002952", //СИЛИКАТНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
					Now:         NowCustomDate(01, 01, 06, 07),
					TimeToWait:  10 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2003960", //ПАВШИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
					Now:     NowCustomDate(1, 1, 7, 58),
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
					Now:         NowCustomDate(1, 1, 8, 30),
					TimeToWait:  10 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000245", //РАБОЧИЙ ПОСЕЛОК
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 8,
					Now:     NowCustomDate(1, 1, 8, 53),
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         NowCustomDate(1, 1, 10, 11),
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000605", //РИЖСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:     10,
					Now:        NowCustomDate(1, 1, 10, 45),
					TimeToWait: 15 * time.Minute,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 11,
					Now:     NowCustomDate(1, 1, 11, 18),
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Now:         NowCustomDate(01, 01, 11, 31),
					Parent:      10,
				},
			},
		},
	}
)
