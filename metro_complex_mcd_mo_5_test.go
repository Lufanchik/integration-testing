package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

func TestComplexMCDMOPartFife(t *testing.T) {
	Run(t, casesComplexMCDMOPartFife)
}

var (
	casesComplexMCDMOPartFife = Cases{
		{
			N: "100. MCD1_MO - MCD2_MSK - MM - MMTS - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "101. MCD1_MO - MCD2_MSK - MM - MMTS - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "102. MCD1_MO - MCD2_MSK - MM - MMTS - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "103. MCD1_MO - MCD2_MSK - MM - MMTS - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},

		{
			N: "104. MCD1_MO - MCD2_MSK - MM - MMTS - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "105. MCD1_MO - MCD2_MSK - MM - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},

		{
			N: "106. MCD1_MO - MCD2_MSK - MM - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "107. MCD1_MO - MCD2_MSK - MCK - MM - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "108. MCD1_MO - MCD2_MSK - MCK - MM - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "109. MCD1_MO - MCD2_MSK - MCK - MM - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "110. MCD1_MO - MCD2_MSK - MCK - MM - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},

		{
			N: "111. MCD1_MO - MCD2_MSK - MCK - MM - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "112. MCD1_MO - MCD2_MSK - MCK - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "113. MCD1_MO - MCD2_MSK - MCK - MMTS - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "114. MCD1_MO - MCD2_MSK - MCK - MMTS - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "115. MCD1_MO - MCD2_MSK - MCK - MMTS - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "116. MCD1_MO - MCD2_MSK - MCK - MMTS - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 7,
				},
			},
		},

		{
			N: "117. MCD1_MO - MCD2_MSK - MCK - MMTS - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     7,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "118. MCD1_MO - MCD2_MSK - MCK - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},

		{
			N: "119. MCD1_MO - MCD2_MSK - MCK - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000620", //ХЛЕБНИКОВО
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "120. MCD1_MO - MCD2_MSK - MMTS - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "121. MCD1_MO - MCD2_MSK - MMTS - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "122. MCD1_MO - MCD2_MSK - MMTS - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "123. MCD1_MO - MCD2_MSK - MMTS - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000590", //ПОКРОВСКОЕ СТРЕШНЕВО (в будущем - Щукинская)
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000001", //КУРСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000700", //ТЕСТОВСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001060", //БЕГОВАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 6,
				},
			},
		},

		{
			N: "124. MCD1_MO - MCD2_MSK - MMTS - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     6,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "125. MCD1_MO - MCD2_MSK - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},

		{
			N: "126. MCD1_MO - MCD2_MSK - MCD1_MO",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000200", //КАЛИТНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     5,
					ExpectedSum: 700,
				},
			},
		},

		{
			N: "127. MCD1_MO - MCD2_MO - MM",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000720", //ОПАЛИХА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "128. MCD1_MO - MCD2_MO - MCK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000720", //ОПАЛИХА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "129. MCD1_MO - MCD2_MO - MMTS",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000720", //ОПАЛИХА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},

		{
			N: "130. MCD1_MO - MCD2_MO - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000720", //ОПАЛИХА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 5,
				},
			},
		},

		{
			N: "131. MCD1_MO - MCD2_MO - MCD1_MSK",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2000115", //ЛОБНЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4900,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					Terminal: &processing.Terminal{
						Station:   "2001410", //ГРАЖДАНСКАЯ
						Direction: processing.TerminalDirection_INGRESS,
					},
					Parent: 1,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MO,
					Terminal: &processing.Terminal{
						Station:   "2000720", //ОПАЛИХА
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress: 3,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Station:   "2000155", //ФИЛИ
						Direction: processing.TerminalDirection_INGRESS,
					},
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Station:   "2001340", //ВОДНИКИ
						Direction: processing.TerminalDirection_EGRESS,
					},
					Ingress:     5,
					ExpectedSum: 700,
				},
			},
		},
	}
)
