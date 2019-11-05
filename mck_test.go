package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
	"time"
)

func TestComplexPassMCKMCK(t *testing.T) {
	Passes(t, casesComplexPassMCKMCK)
}

var (
	// MCK - MCK
	casesComplexPassMCKMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMMM(t *testing.T) {
	Passes(t, casesComplexPassMCKMMMM)
}

var (
	// MM - MM
	casesComplexPassMCKMMMM = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMMCK(t *testing.T) {
	Passes(t, casesComplexPassMCKMMMCK)
}

var (
	// MM - MM
	casesComplexPassMCKMMMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMTSMM(t *testing.T) {
	Passes(t, casesComplexPassMCKMMTSMM)
}

var (
	// MM - MM
	casesComplexPassMCKMMTSMM = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMTSMCK(t *testing.T) {
	Passes(t, casesComplexPassMCKMMTSMCK)
}

var (
	// MM - MM
	casesComplexPassMCKMMTSMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMTSMMTS(t *testing.T) {
	Passes(t, casesComplexPassMCKMMTSMMTS)
}

var (
	// MM - MM
	casesComplexPassMCKMMTSMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMMMTSMM(t *testing.T) {
	Passes(t, casesComplexPassMCKMMMMTSMM)
}

var (
	// MM - MM
	casesComplexPassMCKMMMMTSMM = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMMMTSMCK(t *testing.T) {
	Passes(t, casesComplexPassMCKMMMMTSMCK)
}

var (
	// MM - MM
	casesComplexPassMCKMMMMTSMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMCKMMMMTSMMTS(t *testing.T) {
	Passes(t, casesComplexPassMCKMMMMTSMMTS)
}

var (
	// MM - MM
	casesComplexPassMCKMMMMTSMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)
