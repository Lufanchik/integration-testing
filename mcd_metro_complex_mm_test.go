package http_test

import (
	"testing"
)

func TestComplexPassMMM(t *testing.T) {
	Run(t, casesComplexPassMm)
}

// "MM - MM" (Две платные поездки)
var (
	casesComplexPassMm = Cases{
		{
			//&Pass{
			//	PaymentType: PaymentTypeFullPayment,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MM_SUB,
			//	ExpectedSum: 4200,
			//},
			//&Pass{
			//	PaymentType: PaymentTypeFullPayment,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MM_SUB,
			//	ExpectedSum: 4200,
			//},
		},
	}
)
