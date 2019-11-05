package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
	"time"
)

func TestComplexPassMMMM(t *testing.T) {
	Passes(t, casesComplexPassMMММ)
}

// "MM - MM" (Две платные поездки)
var (
	casesComplexPassMMММ = Cases{
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

func TestComplexPassMMMCKMCK(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMCK)
}

// "MM - MCK - MCK" (Бесплатная пересадка c MM на МЦК, взимание денежных средств за пересадку с МЦК на МЦК)
var (
	casesComplexPassMMMCKMCK = Cases{
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

func TestComplexPassMMMMTSMM(t *testing.T) {
	Passes(t, casesComplexPassMMMMTSMM)
}

// "MM - MMTS -MM" (Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на ММ (ММТС может быть только в начале или конце, поэтому правило промежуточного звена для ММ здесь не работает))
var (
	casesComplexPassMMMMTSMM = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMMTSMCK(t *testing.T) {
	Passes(t, casesComplexPassMMMMTSMCK)
}

// "MM - MMTS -MCK" (Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на МЦК (ММТС может быть только в начале или конце))
var (
	casesComplexPassMMMMTSMCK = Cases{
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMMTSMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMMMTSMMTS)
}

// "MM - MMTS -MMTS" (Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на ММТС)
var (
	casesComplexPassMMMMTSMMTS = Cases{
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMMM(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMMM)
}

// "MM - MCK - MM - MM" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММ, взимание денежных средств за пересадку с ММ на ММ)
var (
	casesComplexPassMMMCKMMMM = Cases{
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

func TestComplexPassMMMCKMMMCK(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMMCK)
}

// "MM - MCK -MM - MCK" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММ, взимание денежных средств за пересадку с ММ на МЦК)
var (
	casesComplexPassMMMCKMMMCK = Cases{
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
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMMMTSMM(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMMMTSMM)
}

// "MM - MCK -MM - MMTS - MM" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММ (Действует правило промежуточного звена для ММ), с ММ на ММТС, взимание денежых средств за пересадку с ММТС на ММ)
var (
	casesComplexPassMMMCKMMMMTSMM = Cases{
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMMMTSMCK(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMMMTSMCK)
}

// "MM - MCK - MM - MMTS - MCK" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММ (Действует правило промежуточного звена для ММ), с ММ на ММТС, взимание денежых средств за пересадку с ММТС на МЦК)
var (
	casesComplexPassMMMCKMMMMTSMCK = Cases{
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMMMTSMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMMMTSMMTS)
}

// "MM - MCK - MM - MMTS - MMTS" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММ (Действует правило промежуточного звена для ММ), с ММ на ММТС, взимание денежых средств за пересадку с ММТС на ММТС)
var (
	casesComplexPassMMMCKMMMMTSMMTS = Cases{
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMTSMM(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMTSMM)
}

// "MM - MCK - MMTS - MM" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств при пересадке с ММТС на ММ)
var (
	casesComplexPassMMMCKMMTSMM = Cases{
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
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMTSMCK(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMTSMCK)
}

// "MM - MCK - MMTS - MCK" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств за пересадку с ММТС на МЦК)
var (
	casesComplexPassMMMCKMMTSMCK = Cases{
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
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

func TestComplexPassMMMCKMMTSMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMMCKMMTSMMTS)
}

// "MM - MCK - MMTS - MMTS" (Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств за пересадку с ММТС на ММТС)
var (
	casesComplexPassMMMCKMMTSMMTS = Cases{
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
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
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

// Кейсы включающие транзакции с некорректной авторизацией

func TestComplexPassMMMMAuthMCKMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMMMAuthMCKMMTS)
}

// "MM - MM (AuthTypeIncorrect) - MCK - MMTS" (если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMMMAuthMCKMMTS = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
		},
	}
)

func TestComplexPassMMMMAuthMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMMMAuthMMTS)
}

//  "MM - MM (AuthTypeIncorrect) - MMTS"  (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMMMAuthMMTS = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
		},
	}
)

func TestComplexPassMMAuthMMAuthMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMAuthMMAuthMMTS)
}

//  "MM - MM (AuthTypeIncorrect) - MMTS"  (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMAuthMMAuthMMTS = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFree,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      2,
			},
		},
	}
)

func TestComplexPassMMAuthMCKAuthMMTS(t *testing.T) {
	Passes(t, casesComplexPassMMAuthMCKAuthMMTS)
}

//  "MM (AuthTypeIncorrect) - MCK  - MMTS"  (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMAuthMCKAuthMMTS = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
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
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
			},
		},
	}
)

func TestComplexPassMMAuthMCKAuthMCKMM(t *testing.T) {
	Passes(t, casesComplexPassMMAuthMCKAuthMCKMM)
}

//  "MM - MM (AuthTypeIncorrect) - MCK - MM"  (Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
var (
	casesComplexPassMMAuthMCKAuthMCKMM = Cases{
		{
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    AuthTypeIncorrect,
			},
			&Pass{
				PaymentType: PaymentTypeFullPayment,
				RequestType: RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
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
		},
	}
)
