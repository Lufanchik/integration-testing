package passes

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesCancel = test.Cases{
	{
		N: "1. ММ успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "2. МЦК успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "3. ММТС успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "4. МСК-МСК2 успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000235", //ДМИТРОВСКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "5. МСК-МО2 успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     1,
				ExpectedSum: 700,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "6. МО-МСК2 успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "7. МО-МО2 успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2002952", //СИЛИКАТНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "8. МСК-МСК успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "9. МСК-МО успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000685", //БАКОВКА
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     1,
				ExpectedSum: 700,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "10. МО-МСК успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000685", //БАКОВКА
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "11. МО-МО успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000600", //НОВОДАЧНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "12. МЦК ММ успешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 3,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 5,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 7,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},

	//КЕЙСЫ С ДОАВТОРИЗАЦИЕЙ

	{
		N: "13. ММ неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "14. МЦК неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "15. ММТС неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "16. МСК-МСК2 неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000235", //ДМИТРОВСКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "17. МСК-МО2 неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     1,
				ExpectedSum: 700,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "18. МО-МСК2 неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "19. МО-МО2 неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2002952", //СИЛИКАТНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "20. МСК-МСК неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "21. МСК-МО неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000685", //БАКОВКА
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     1,
				ExpectedSum: 700,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "22. МО-МСК неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2000685", //БАКОВКА
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000155", //ФИЛИ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "23. МО-МО неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				AuthType:    test.AuthTypeIncorrect,
				Terminal: &processing.Terminal{
					Station:   "2001101", //ИННОВАЦИОННЫЙ ЦЕНТР
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000600", //НОВОДАЧНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Cancel{
				Target: 2,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
	{
		N: "24. МЦК ММ неуспешная авторизация + Отмена",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 3,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 5,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeIncorrect,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 7,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
}
