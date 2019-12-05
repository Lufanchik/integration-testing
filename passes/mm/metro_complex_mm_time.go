package mm

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
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
		//{
		//	N: "2. ММ - МЦК / Втечение 90 мин",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType:      test.PaymentTypePayment,
		//			Carrier:          carriers.Carrier_MM,
		//			SubCarrier:       carriers.SubCarrier_MM_SUB,
		//			Now:              test.NowFullDate(2019, 11, 01, 23, 50, 00),
		//			IsComplexTimeout: true,
		//			ExpectedSum:      4200,
		//		},
		//		&test.Pass{
		//			PaymentType:      test.PaymentTypeFree,
		//			Carrier:          carriers.Carrier_MM,
		//			SubCarrier:       carriers.SubCarrier_MCK_SUB,
		//			Now:              test.NowFullDate(2019, 11, 01, 23, 55, 00),
		//			IsComplexTimeout: true,
		//			Parent:           1,
		//		},
		//	},
		//},
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
		//{
		//	N: "4. ММ - ММТС / Втечение 90 мин",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         test.NowFullDate(2019, 12, 01, 10, 30, 00),
		//			ExpectedSum: 4200,
		//			IsComplexTimeout: true,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypeFree,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Now:         test.NowFullDate(2019, 12, 01, 11, 59, 59),
		//			Parent:      1,
		//			IsComplexTimeout: true,
		//		},
		//	},
		//},
		//{
		//	N: "5. ММ - ММТС / После 90 мин",
		//	T: test.T{
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MM_SUB,
		//			Now:         test.NowFullDate(2019, 12, 05, 10, 30, 00),
		//			ExpectedSum: 4200,
		//			IsComplexTimeout: true,
		//		},
		//		&test.Pass{
		//			PaymentType: test.PaymentTypePayment,
		//			Carrier:     carriers.Carrier_MM,
		//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
		//			Now:         test.NowFullDate(2019, 12, 05, 12, 00, 01),
		//			ExpectedSum: 4200,
		//			IsComplexTimeout: true,
		//		},
		//	},
		//},
	}
)
