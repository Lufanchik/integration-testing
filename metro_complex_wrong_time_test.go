package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
	"time"
)

var (
	casesWrongTimeComplexPass = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Date(
						2019, 11, 13, 21, 34, 58, 651387237, time.UTC).UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				Now: func() uint64 {
					return uint64(time.Date(
						2019, 11, 13, 21, 39, 58, 651387237, time.UTC).UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Date(
						2019, 11, 13, 21, 36, 56, 651387237, time.UTC).UnixNano())
				},
			},
			//&PassCheck{
			//	Target:      2,
			//	PaymentType: PaymentTypeFree,
			//	ExpectedSum: 0,
			//	Parent:      3,
			//},
			//&Pass{
			//	PaymentType: PaymentTypeFree,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			//	Parent:      1,
			//},
			//&Pass{
			//	PaymentType: PaymentTypeFullPayment,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//	ExpectedSum: 4200,
			//},
		},
	}
)

func TestWrongTimeComplexPass(t *testing.T) {
	Run(t, casesWrongTimeComplexPass)
}
