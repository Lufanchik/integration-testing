package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestOfflineMetroComplexMM(t *testing.T) {
	Run(t, casesOfflineMetroComplexMM)
}

// "MM - MM" (Две платные поездки)
var (
	casesOfflineMetroComplexMM = Cases{
		{
			N: "MCK - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MCK", //(Бесплатная пересадка c MM на МЦК, взимание денежных средств за пересадку с МЦК на МЦК)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MMTS -MM", //(Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на ММ (ММТС может быть только в начале или конце, поэтому правило промежуточного звена для ММ здесь не работает))
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MMTS -MCK", //(Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на МЦК (ММТС может быть только в начале или конце))
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MMTS -MMTS", //(Бесплатная пересадка с ММ на ММТС, взимание денежных средств за пересадку с ММТС на ММТС)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MM - MM", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММ, взимание денежных средств за пересадку с ММ на ММ)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK -MM - MCK", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММ, взимание денежных средств за пересадку с ММ на МЦК)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK -MM - MMTS - MM", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММ (Действует правило промежуточного звена для ММ), с ММ на ММТС, взимание денежых средств за пересадку с ММТС на ММ)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MM - MMTS - MCK", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММ (Действует правило промежуточного звена для ММ), с ММ на ММТС, взимание денежых средств за пересадку с ММТС на МЦК)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MM - MMTS - MMTS", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММ (Действует правило промежуточного звена для ММ), с ММ на ММТС, взимание денежых средств за пересадку с ММТС на ММТС)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MMTS - MM", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств при пересадке с ММТС на ММ)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MMTS - MCK", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств за пересадку с ММТС на МЦК)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "MM - MCK - MMTS - MMTS", //(Бесплатная пересадка с ММ на МЦК, с МЦК на ММТС, взимание денежных средств за пересадку с ММТС на ММТС)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		//// Кейсы включающие транзакции с некорректной авторизацией
		{
			N: "MM - MM (AuthTypeIncorrect) - MCK - MMTS", //(если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      2,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      2,
				},
			},
		},
		{
			N: "MM - MM (AuthTypeIncorrect) - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      2,
				},
			},
		},
		{
			N: "MM - MM (AuthTypeIncorrect) - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      2,
				},
			},
		},
		{
			N: "MM (AuthTypeIncorrect) - MCK  - MMTS", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
			},
		},
		{
			N: "MM - MM (AuthTypeIncorrect) - MCK - MM", //(Если было две одинаковые поездки и последняя из них неоплачена, комплексная поездка должна создаваться и привязываться к последней)
			T: T{
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFullPayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      2,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      2,
				},
			},
		},
	}
)
