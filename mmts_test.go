package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
	"time"
)

var (
	// MMTS - MMTS
	casesComplexPassMMTSMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

var (
	// MMTS - MM - MM
	casesComplexPassMMTSMMMM = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
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

var (
	// MMTS - MCK - MCK
	casesComplexPassMMTSMCKMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
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

var (
	// MMTS - MM - MMTS
	casesComplexPassMMTSMMMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

var (
	// MMTS - MM - MCK - MCK
	casesComplexPassMMTSMMMCKMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
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

var (
	// MMTS - MM - MCK - MMTS
	casesComplexPassMMTSMMMCKMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

var (
	// MMTS - MCK - MM - MM
	casesComplexPassMMTSMCKMMMM = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
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

var (
	// MMTS - MCK - MM - MCK
	casesComplexPassMMTSMCKMMMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
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

var (
	// MMTS - MCK - MM - MMTS
	casesComplexPassMMTSMCKMMMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

var (
	// MMTS - MM -MCK - MM -MM
	casesComplexPassMMTSMMMCKMMMM = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
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

var (
	// MMTS - MM -MCK - MM -MCK
	casesComplexPassMMTSMMMCKMMMCK = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
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

var (
	// MMTS - MM -MCK - MM -MMTS
	casesComplexPassMMTSMMMCKMMMMTS = Cases{

		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				Now: func() uint64 {
					return uint64(time.Now().UnixNano())
				},
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
		},
	}
)

func TestComplexPassMMMMTSAuthMCK(t *testing.T) {
	Passes(t, casesComplexPassMMMMTSAuthMCK)
}

//  MMTS - MMTS (not money) - MCK (если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMMMTSAuthMCK = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      2,
			},
		},
	}
)

// Кейсы включающие транзакции с некорректной авторизацией

func TestComplexPassMMTSMMTSAuthMMMCKMM(t *testing.T) {
	Passes(t, casesComplexPassMMTSMMTSAuthMMMCKMM)
}

// "MMTS - MMTS (AuthTypeIncorrect) - MM - MCK -MM" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMTSMMTSAuthMMMCKMM = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      2,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
			},
		},
	}
)

func TestComplexPassMMTSMMTSAuthMCKMM(t *testing.T) {
	Passes(t, casesComplexPassMMTSMMTSAuthMCKMM)
}

// "MMTS - MMTS (AuthTypeIncorrect) - MM - MCK -MM" (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMTSMMTSAuthMCKMM = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      2,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
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

func TestComplexPassMMTSMMTS(t *testing.T) {
	Run(t, casesComplexPassMMTSMMTS)
}

func TestComplexPassMMTSMMMM(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMM)
}

func TestComplexPassMMTSMCKMCK(t *testing.T) {
	Run(t, casesComplexPassMMTSMCKMCK)
}

func TestComplexPassMMTSMMMMTS(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMMTS)
}

func TestComplexPassMMTSMMMCKMCK(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMCKMCK)
}

func TestComplexPassMMTSMMMCKMMTS(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMCKMMTS)
}

func TestComplexPassMMTSMCKMMMM(t *testing.T) {
	Run(t, casesComplexPassMMTSMCKMMMM)
}

func TestComplexPassMMTSMCKMMMCK(t *testing.T) {
	Run(t, casesComplexPassMMTSMCKMMMCK)
}

func TestComplexPassMMTSMCKMMMMTS(t *testing.T) {
	Run(t, casesComplexPassMMTSMCKMMMMTS)
}

func TestComplexPassMMTSMMMCKMMMM(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMCKMMMM)
}

func TestComplexPassMMTSMMMCKMMMCK(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMCKMMMCK)
}

func TestComplexPassMMTSMMMCKMMMMTS(t *testing.T) {
	Run(t, casesComplexPassMMTSMMMCKMMMMTS)
}
