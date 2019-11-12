package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
	"time"
)

func TestComplexPassMMM(t *testing.T) {
	Run(t, casesComplexPassMMМ)
}

// "MM - MM" (Две платные поездки)
var (
	casesComplexPassMMМ = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)
