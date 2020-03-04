package mmts

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var (
	CasesComplexTimeMMTS = test.Cases{
		{
			N: "1.ММТС-ММТС",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 10, 30, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 12, 00, 01),

					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "2.ММТС-МЦК", //Втечение 90 мин
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 11, 01, 23, 50, 00),
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType:      test.PaymentTypeFree,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MCK_SUB,
					Now:              test.NowFullDate(2019, 11, 01, 23, 55, 00),
					IsComplexTimeout: true,
					Parent:           1,
				},
			},
		},
		{
			N: "3.ММТС-МЦК", //После90мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 8, 30, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 10, 00, 01),

					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "4.ММТС-ММ", //Втечение 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 10, 30, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType:      test.PaymentTypeFree,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 11, 59, 59),
					Parent:           1,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "5.ММТС-ММ", //После 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 05, 17, 30, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 05, 19, 00, 01),

					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "6.ММТС-МСК-МСК", //Втечение 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 2, 8, 15, 04),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:           1,
					Now:              test.NowFullDate(2019, 12, 2, 8, 33, 00),
					TimeToWait:       5 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:          2,
					Now:              test.NowFullDate(2019, 12, 2, 9, 00, 00),
					ExpectedSum:      700,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "7.ММТС-МСК-МСК", //После90мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 2, 6, 15, 04),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum:      4200,
					Now:              test.NowFullDate(2019, 12, 2, 8, 33, 00),
					TimeToWait:       5 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:          2,
					Now:              test.NowFullDate(2019, 12, 2, 9, 00, 00),
					ExpectedSum:      700,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "8.ММТС-МСК-МСК2", //Втечение 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 10, 30, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000605", //РИЖСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:           1,
					Now:              test.NowFullDate(2019, 12, 01, 10, 45, 00),
					TimeToWait:       15 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:          2,
					Now:              test.NowFullDate(2019, 12, 01, 11, 18, 00),
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "9.ММТС-МСК-МСК2", //После 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 9, 00, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000605", //РИЖСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:           10,
					Now:              test.NowFullDate(2019, 12, 01, 10, 45, 00),
					TimeToWait:       15 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:          11,
					Now:              test.NowFullDate(2019, 12, 01, 11, 18, 00),
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "10.ММТС-МСК-МО", //Втечение 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 05, 20, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:           1,
					Now:              test.NowFullDate(2019, 12, 01, 05, 35, 00),
					TimeToWait:       10 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:          2,
					Now:              test.NowFullDate(2019, 12, 01, 06, 07, 00),
					ExpectedSum:      700,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "11.ММТС-МСК-МО", //После 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 01, 03, 30, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum:      4200,
					Now:              test.NowFullDate(2019, 12, 01, 05, 35, 00),
					TimeToWait:       10 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001101", //ИННОВАЦИОННЫЙЦЕНТР
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:          4,
					Now:              test.NowFullDate(2019, 12, 01, 06, 07, 00),
					ExpectedSum:      700,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "10.ММТС-МСК-МО2", //Втечение 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 31, 00, 05, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840", //ТРИКОТАЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:           1,
					Now:              test.NowFullDate(2019, 12, 31, 00, 40, 45),
					TimeToWait:       5 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Now:              test.NowFullDate(2019, 12, 31, 01, 8, 45),
					Ingress:          2,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "11.ММТС-МСК-МО2", //После 90 мин
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Now:         test.NowFullDate(2019, 12, 31, 00, 00, 00),

					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001840", //ТРИКОТАЖНАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent:           1,
					Now:              test.NowFullDate(2019, 12, 31, 01, 40, 45),
					TimeToWait:       5 * time.Minute,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000460", //НАХАБИНО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Now:              test.NowFullDate(2019, 12, 31, 02, 8, 45),
					Ingress:          2,
					IsComplexTimeout: true,
				},
			},
		},
	}
)
