package passes

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var CasesWrongTimeComplexPass = test.Cases{
	{
		N: "Wrong time complex",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,

				Now: func() uint64 {
					return uint64(time.Date(
						2019, 11, 13, 21, 34, 58, 651387237, time.UTC).UnixNano())
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				Now: func() uint64 {
					return uint64(time.Date(
						2019, 11, 13, 21, 39, 58, 651387237, time.UTC).UnixNano())
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,

				Now: func() uint64 {
					return uint64(time.Date(
						2019, 11, 13, 21, 36, 56, 651387237, time.UTC).UnixNano())
				},
			},
			&test.PassCheck{
				Target:      2,
				PaymentType: test.PaymentTypeFree,
				ExpectedSum: 0,
				Parent:      3,
			},
			//&Pass{
			//	PaymentType: PaymentTypeFree,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//	Parent:      1,
			//},
			//&Pass{
			//	PaymentType: PaymentTypePayment,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//
			//},
		},
	},
}
