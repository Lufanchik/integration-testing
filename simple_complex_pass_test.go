package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

var (
	casesSimpleComplexPass = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			//&Pass{
			//	PaymentType: PaymentTypeFree,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MM_SUB,
			//	Parent:      1,
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

func TestSimpleComplexPass(t *testing.T) {
	Run(t, casesSimpleComplexPass)
}
