package mm

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var (
	CasesComplexTimeMM = test.Cases{
		{
			N: "1. ММ - ММ",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 10, 30, 00),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 11, 00, 00),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "2. ММ - МЦК / Втечение 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 11, 01, 23, 50, 00),
					IsComplexTimeout: true,
					ExpectedSum:      4200,
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
			N: "3. ММ - МЦК / После 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 10, 30, 00),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MCK_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 12, 00, 01),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "4. ММ - ММТС / Втечение 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 10, 30, 00),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType:      test.PaymentTypeFree,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 11, 59, 59),
					Parent:           1,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "5. ММ - ММТС / После 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MM_SUB,
					Now:              test.NowFullDate(2019, 12, 05, 10, 30, 00),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 05, 12, 00, 01),
					ExpectedSum:      4200,
					IsComplexTimeout: true,
				},
			},
		},
		{
			N: "8. ММТС - МСК-МСК2 / Втечение 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 10, 30, 00),
					ExpectedSum:      4200,
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
			N: "9. ММТС - МСК-МСК2 / После 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 9, 00, 00),
					ExpectedSum:      4200,
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
			N: "10. ММТС - МСК-МО / Втечение 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 05, 20, 00),
					ExpectedSum:      4200,
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
						Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
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
			N: "11. ММТС - МСК-МО / После 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 01, 03, 30, 00),
					ExpectedSum:      4200,
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
						Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
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
			N: "10. ММТС - МСК-МО2 / Втечение 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 31, 00, 05, 00),
					ExpectedSum:      4200,
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
			N: "11. ММТС - МСК-МО2 / После 90 мин",
			T: test.T{
				&test.Pass{
					PaymentType:      test.PaymentTypePayment,
					Carrier:          carriers.Carrier_MM,
					SubCarrier:       carriers.SubCarrier_MMTS_SUB,
					Now:              test.NowFullDate(2019, 12, 31, 00, 00, 00),
					ExpectedSum:      4200,
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
