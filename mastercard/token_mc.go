package mastercard

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesTokensMC = test.Cases{
	{
		N:          "1. MM success authorization + Cancel",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeCorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N:          "2. MM unsuccess authorization + Cancel",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenIncorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N:          "3. MM-МCК-MCD1_MSК_МSК-ММТS-ММ",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000245", //РАБОЧИЙПОСЕЛОК
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 3,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
		},
	},
	//{
	//	N:          "4. Face/MM-MCK-MMTS",
	//	CardSystem: processing.CardSystem_MIR,
	//	CustomerId: uuid.New().String(),
	//	T: test.T{
	//		&test.RegisterFaceId{},
	//		&test.FaceIdRegistrationStatus{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MM_SUB,
	//		},
	//		&test.RegisterFaceId{},
	//		&test.FaceIdRegistrationStatus{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeFree,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
	//			Parent:      3,
	//		},
	//		&test.RegisterFaceId{},
	//		&test.FaceIdRegistrationStatus{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeFree,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
	//			Parent:      3,
	//		},
	//		&test.RegisterFaceId{},
	//		&test.FaceIdRegistrationStatus{},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MM_SUB,
	//		},
	//	},
	//},
	{
		N:          "5. MCD MSK1 - MCD MSK2 - MCD MO2", //
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000275",
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001270",
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000075",
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000460",
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     3,
				ExpectedSum: 800,
			},
		},
	},
	{
		N:          "6. MSK - MSK - MCK - MMTS - MMTS",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000185", //ЛИАНОЗОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
		},
	},
	{
		N:          "7. MCK - MCD1_MSK - MM-MCD2_MSK - MCD1_MSK",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001270", //ОКРУЖНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2002780", //ДЕПО
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000045", //ТЕКСТИЛЬЩИКИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 5,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001060", //БЕГОВАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 7,
			},
		},
	},
	{
		N:          "8. МGТ - МCD2_MО - MCD2_МСК - МGТ - МCК - МGТ - ММ - МGТ - ММТS - МGТ",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2003960",
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 5100,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2002780",
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
			},
		},
	},
	{
		N:          "9. MM - MCD1_MSK_МSК - MM - MМТS - MCD1_МSК_МО",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000275", //
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2002077", //
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2002077", //
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001101", //ИННОВАЦИОННЫЙЦЕНТР
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     6,
				ExpectedSum: 800,
			},
		},
	},
	{
		N:          "10. MMTS - MSK - MSK - MM - MCK - MSK - MO",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001140", //КУНЦЕВСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000009", //САВЕЛОВСКИЙВОКЗАЛ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001140", //КУНЦЕВСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     6,
				ExpectedSum: 800,
			},
		},
	},
	{
		N:          "11. MTPPK MASTER",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 2,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 2,
			},
		},
	},
	{
		N:          "12. MTPPK - MCD-MSK1 MASTER",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2001270", //ОКРУЖНАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4600,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8800,
			},
			&test.Complete{
				StartPass: 3,
				Passes: []int{
					4, 5,
				},
				Sum: 8800,
			},
		},
	},
	{
		N:          "13. МТППК - МЦД-МО1 MASTER",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000115", //ЛОБНЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 5100,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8800,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   3,
				ExpectedSum: 8800,
			},
			&test.Complete{
				StartPass: 3,
				Passes: []int{
					4, 5,
				},
				Sum: 8800,
			},
		},
	},
	{
		N:          "14. MTPPK-MCK-MM-MCK-MTPPK MASTER",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 8400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      2,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 8400,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					5,
				},
				Sum: 8400,
			},
		},
	},
	{
		N:          "15. МТППК - МГТ MASTER",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 42,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 42,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 42,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 42,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					3, 4, 6,
				},
				Sum: 42,
			},
		},
	},
	{
		N:          "16. ММТС - MTППК - ММТС - MTPPK - MMTS - MTPPK - COMPLETE",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeMCTokenCorrect,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				RequestType: test.RequestTypeOffline,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOffline,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				RequestType: test.RequestTypeOffline,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				AuthType:    test.AuthTypeMCTokenCorrect,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				AuthType:    test.AuthTypeMCTokenCorrect,
				RequestType: test.RequestTypeOffline,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   2,
				ExpectedSum: 6,
			},
			&test.Complete{
				StartPass: 2,
				Passes: []int{
					4, 6, 8,
				},
				Sum: 6,
			},
		},
	},
	{
		N:          "17. MCK success authorization + Cancel",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N:          "18. MCK success authorization + Cancel",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N:          "19. MM - MCK - MMTS success authorization + Cancel averything",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeMCTokenCorrect,
				Parent:      1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 3,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
}
