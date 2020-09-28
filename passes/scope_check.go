package passes

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"time"
)

//Кейсы объема испытаний

var CasesScopeCheckPass = test.Cases{
	{
		N: "1. Запрет прохода для ББК МПС CUP и перевозчика МГТ",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "2. Запрет прохода для дебетовой ББК МПС CUP,  Эмитент из РФ",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "3. Запрет прохода по ББК, BIN которой находится в Стоп-листе",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "4. Offline-проход по ББК, BIN которой не находится в списке на online-авторизацию",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "5. Offline-проход по ББК, BIN которой находится в списке на online-авторизацию, PAN-хэш карты находится в white-листе",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "6. Online-проход при недоступном ТП",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "7. Online-проход при неполучении ответа от ТП при выполнении online-операций",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "8. Запрет online-прохода по ББК с неактивным счетом",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "9. Запрет online-прохода по ББК с недостаточным балансом",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "10. Offline-проход по ББК с недостаточным балансом",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "11. Offline-проходы в пределах одной поездки при добавлении ББК в СТОП-лист при первом проходе",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "12. Отмена регистрации поездки",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "13. Запрет прохода при повторном прикладывании",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "14. Online обработка по проходу, информация по которому уже обрабатывалась ранее. PayGate MM доступен.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "15. Online обработка по проходу, старше Т – 3 дня. PayGate MM доступен",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "16. Online обработка при попытке прохода по ББК с недостаточным балансом",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "17. Online обработка по проходу, при отсутствии других проходов в интервале времени +/- 90 минут от времени полученной операции. (ММ, ММТС, МЦК)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "18. Offline обработка по проходу, информация по которому уже обрабатывалась ранее",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "19. Offline обработка по проходу, старше Т – 3 дня", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "20. Offline обработка при попытке прохода по ББК с недостаточным балансом", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "21. Offline обработка по проходу, при отсутствии других проходов в интервале времени +/- 90 минут от времени полученной операции. (ММ, ММТС, МЦК)", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "22. Обработка прохода от Аэроэкспресс, который был обработан ранее", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "23. Обработка прохода от Аэроэкспресс, который был совершен раньше чем Т – 3 дня", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "24. Обработка прохода от Аэроэкспресс при отсутствии других проходов в промежутке +/- 300 минут", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MGT,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
		},
	},
	{
		N: "25. Одиночная поездка ММ",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
			},
		},
	},
	{
		N: "26. Одиночная поездка МЦК",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
			},
		},
	},
	{
		N: "27. Одиночная поездка ММТС",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
		},
	},
	{
		N: "28. Одиночная поездка за пределами МСК (Вход МЦД МО – Выход МЦД МСК)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
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
					Station:   "2000700", //ТЕСТОВСКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
			},
		},
	},
	{
		N: "29. Одиночная поездка за пределами МСК (Вход МЦД МСК – Выход МЦД МО)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000700", //ТЕСТОВСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				ExpectedSum: 700,
				Ingress:     1,
			},
		},
	},
	{
		N: "30. Одиночная поездка МЦД (Вход МЦД МСК – валидация МЦД МО)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000700", //ТЕСТОВСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     1,
				ExpectedSum: 700,
			},
		},
	},
	{
		N: "31. Одиночная поездка МЦД (МЦД МО валидация –  МЦД МО выход)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000685", //БАКОВКА
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
				Now:         test.NowFullDate(2019, 12, 31, 3, 30, 00),
				TimeToWait:  10 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000100", //ДОЛГОПРУДНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 12, 31, 4, 51, 00),
			},
		},
	},
	{
		N: "32. Одиночная поездка МЦД (МЦД МО валидация –  МЦД МО валидация)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
				Now:         test.NowFullDate(2019, 12, 31, 3, 30, 00),
				TimeToWait:  10 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000115", //ЛОБНЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 12, 31, 3, 30, 00),
			},
		},
	},
	{
		N: "33. Информация о выходе по МЦД пришла после закрытия периода комплексной поездки, но до списания комиссионного сбора.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_INGRESS,
				},
				ExpectedSum: 4900,
				Now:         test.NowFullDate(2019, 12, 31, 3, 30, 00),
				TimeToWait:  5 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Now:         test.NowFullDate(2019, 12, 31, 4, 10, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000600", //НОВОДАЧНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 12, 31, 3, 48, 00),
			},
		},
	},
	{
		N: "34. Комиссионный сбор при неосуществлении выхода с МЦД (МЦД МСК - выхода с МЦД не последовало)", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Direction: 0,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Direction: 0,
				},
			},
		},
	},
	{
		N: "35. Комиссионный сбор при неосуществлении выхода с МЦД (МЦД МО - выхода с МЦД не последовало)", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Direction: 0,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Direction: 0,
				},
			},
		},
	},
	{
		N: "36. Комиссионный сбор при незакрытой валидации прохода МЦД (Валидация МЦД МО – информации о входе или выходе не последовало)", //не описан
		T: test.T{
			processing.CardSystem_MIR,
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Direction: 0,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Direction: 0,
				},
			},
		},
	},
	{
		N: "37. Комплексная поездка во время перехода на новые транспортные сутки.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				ExpectedSum: 4900,
				Now:         test.NowFullDate(2019, 12, 31, 3, 30, 00),
				TimeToWait:  30 * time.Minute,
				Terminal: &processing.Terminal{
					Station:   "2000055", //ОДИНЦОВО
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2000155", //ФИЛИ
				},
				Now:     test.NowFullDate(2019, 12, 31, 4, 25, 00),
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 12, 31, 4, 40, 00),
			},
		},
	},
	{
		N: "38. Попытка повторного прохода для ММ на одном терминале.", //не описан
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 12, 2, 3, 30, 15),
				Terminal: &processing.Terminal{
					Id: "lacoste",
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 12, 2, 3, 30, 17),
				Terminal: &processing.Terminal{
					Id: "lacoste",
				},
			},
		},
	},
	{
		N: "39. Повторный проход для ММ на разных терминалах.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 12, 2, 3, 30, 15),
				Terminal: &processing.Terminal{
					Id:      "sirocco",
					Station: "123",
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 12, 2, 3, 30, 30),
				Terminal: &processing.Terminal{
					Id:      "multicarta",
					Station: "123",
				},
			},
		},
	},
	{
		N: "40. Дублирование звеньев в пределах МСК на МЦД в пределах одной линии (МЦД линия 1 – МЦД линия 1)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Station:   "2001140", //КУНЦЕВСКАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Now: test.NowFullDate(2019, 10, 10, 10, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000006", //БЕЛОРУССКАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Now:     test.NowFullDate(2019, 10, 10, 10, 00, 00),
				Ingress: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Direction: 1,
				},
				Now: test.NowFullDate(2019, 10, 10, 10, 00, 00),
			},
		},
	},
	{
		N: "41. Повторение вида поездки внутри комплексной поездки для ММ (ММ – МЦК - ММ)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 10, 10, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 10, 10, 20, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 10, 10, 40, 00),
			},
		},
	},
	{
		N: "42. Повторение вида поездки внутри комплексной поездки (МЦК – ММ -МЦК)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 10, 10, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 10, 10, 30, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 10, 10, 40, 00),
			},
		},
	},
	{
		N: "43. Дублирование звеньев в пределах МСК на МЦД в пределах разных линий (МЦД линия 1 – МЦД линия 2)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Station:   "2001060", //БЕГОВАЯ
					Direction: processing.TerminalDirection_INGRESS,
				},
				Now:        test.NowFullDate(2019, 10, 10, 10, 00, 00),
				TimeToWait: 5 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Station:   "2001270", //ОКРУЖНАЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 10, 10, 10, 28, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000200", //КАЛИТНИКИ
				},
				Parent:     1,
				Now:        test.NowFullDate(2019, 10, 10, 10, 40, 00),
				TimeToWait: 5 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2002780", //ДЕПО
				},
				Ingress: 3,
				Now:     test.NowFullDate(2019, 10, 10, 10, 56, 00),
			},
		},
	},
	{
		N: "44. Закрытие комплексной поездки по максимальному количеству пересадок в пределах МСК (ММТС – ММ – МЦК – МЦД МСК вход – МЦД МСК выход)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Parent:      1,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000200", //КАЛИТНИКИ
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2002780", //ДЕПО
				},
				Ingress: 4,
			},
		},
	},
	{
		N: "45. Закрытие комплексной поездки по максимальному количеству пересадок за пределами МСК (ММТС – МЦК – ММ – МЦД МСК вход – МЦД МО выход)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Parent:      1,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000075", //ТУШИНО
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				ExpectedSum: 700,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003960", //ПАВШИНО
				},
				Ingress: 4,
			},
		},
	},
	{
		N: "46. Закрытие комплексной поездки при проходе на ММТС (МЦК – ММ – ММТС)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 10, 10, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 10, 11, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 10, 11, 30, 00),
			},
		},
	},
	{
		N: "47. Поездка из Аэропорта Шереметьево на Аэроэкспрессе без информации о выходе (Аэропорт вход – информации о выходе не последовало)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "48. Поездка из Аэропорта Шереметьево на Аэроэкспрессе и выход на МЦД МО (Аэропорт вход – МЦД МО выход)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "49. Поездка из Аэропорта Шереметьево на Аэроэкспрессе и валидация на станции МЦД МО (Аэропорт вход – МЦД МО валидация без указания направления прохода)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "50. Поездка в Аэропорт Шереметьево на Аэроэкспрессе без информации о входе (Аэропорт выход)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "51. Поездка на МЦД МО и Аэроэкспресс в Аэропорт Шереметьево (МЦД МО вход - Аэропорт выход)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "52. Поездка из Аэропорта Шереметьево на Аэроэкспрессе и валидация на станции МЦД МО (Аэропорт вход – МЦД МО валидация)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "53. Поездка из Аэропорта Шереметьево на Аэроэкспрессе и пересадка на ММ (Аэропорт вход – ММ)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "54. Закрытие периода комплексной поездки при выходе за временный лимит 5 часов при наличии прохода на Аэроэкспресс. (Аэроэкспресс вход – МЦД МО валидация)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "55. Закрытие периода комплексной поездки при наличии прохода на выход на Аэроэкспресс. (МЦД МО валидация - Аэроэкспресс выход – МЦД МО валидация)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 0,
			},
		},
	},
	{
		N: "56. Открытие периода комплексной поездки при получении прохода на вход в Аэроэкспресс (ММ – Аэропорт вход)", //не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 2, 10, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_AE,
				SubCarrier:  carriers.SubCarrier_NONE_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 2, 11, 00, 00),
				Terminal: &processing.Terminal{
					Station:   "123",
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
		},
	},
	{
		N: "57. Открытие периода комплексной поездки при отсутствии прохода в интервале [-90 мин;+90 мин] относительно полученного прохода. (ММ – МЦК - ММ)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 2, 10, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 2, 11, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 2, 11, 30, 01),
			},
		},
	},
	{
		N: "58. Продление периода комплексной поездки по параметру времени посадки для поездки по МЦД. (МЦД МО вход – МЦД МСК выход – ММ - МЦК)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000155", //ФИЛИ
				},
				ExpectedSum: 4900,
				Now:         test.NowFullDate(2019, 06, 15, 12, 00, 00),
				TimeToWait:  30 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2001270", //ОКРУЖНАЯ
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 06, 15, 13, 00, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 06, 15, 13, 35, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 06, 15, 14, 00, 00),
			},
		},
	},
	{
		N: "59. В ТП пришел проход, ставший пересадкой внутри комплексной поездки, не изменив следующую поездку.",
		T: test.T{
			&test.Pass{
				PaymentType:      test.PaymentTypePayment,
				Carrier:          carriers.Carrier_MM,
				SubCarrier:       carriers.SubCarrier_MM_SUB,
				Now:              test.NowFullDate(2019, 12, 01, 10, 30, 00),
				ExpectedSum:      4200,
				IsComplexTimeout: true,
			},
			&test.Pass{
				PaymentType:      test.PaymentTypeFree,
				Carrier:          carriers.Carrier_MM,
				SubCarrier:       carriers.SubCarrier_MMTS_SUB,
				Now:              test.NowFullDate(2019, 12, 01, 11, 59, 59),
				Parent:           1,
				IsComplexTimeout: true,
			},
			&test.Pass{
				PaymentType:      test.PaymentTypeFree,
				Carrier:          carriers.Carrier_MM,
				SubCarrier:       carriers.SubCarrier_MMTS_SUB,
				Now:              test.NowFullDate(2019, 12, 01, 11, 10, 59),
				Parent:           1,
				IsComplexTimeout: true,
			},
		},
	},
	{
		N: "60. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000700", //ТЕСТОВСКАЯ
				},
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 20, 15, 00, 00),
				TimeToWait:  10 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 10, 20, 15, 33, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 20, 15, 51, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 20, 10, 51, 00),
			},
		},
	},
	{
		N: "61. После пересчета с учетом нового прохода Авторизационный проход существующей поездки стал пересадкой", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000700",
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "62. После пересчета с учетом нового прохода Пересадка стала Авторизационным проходом в существующей поездке", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000700", //ТЕСТОВСКАЯ
				},
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 20, 15, 00, 00),
				TimeToWait:  10 * time.Minute,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965", //ТИМИРЯЗЕВСКАЯ
				},
				Ingress: 1,
				Now:     test.NowFullDate(2019, 10, 20, 15, 33, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				Parent:      1,
				Now:         test.NowFullDate(2019, 10, 20, 15, 51, 00),
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				Now:         test.NowFullDate(2019, 10, 20, 10, 51, 00),
			},
		},
	},
	{
		N: "63. Отмена оплаты при пересчете при поступлении информации о проходе на Аэроэкспресс", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2000700",
				},
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "64. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "65. Коррекция стоимости поездки при смене тарифа в большую сторону у комплексных поездок после  пересчета при корректных поездках на МЦК", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "66. Успешная доавторизация операций в день Т + 4", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "67. Прекращение запросов на доавторизацию после наступления даты Т+14", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "68. Закрытие периода по Неавторизованным операциям", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "69. Возмещение средств за предоставление услуг перевозчикам", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "70. Проверка состава файла реестра Расхождений при их наличии", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "71. Проверка состава файла реестра Распределения денежных средств", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "72. Возврат денежных средств, переведенных перевозчику в случае возникновения спорной ситуации", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "73. Операция есть в Way4, но нет в ТП", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "74. Операции нет в Way4, но есть в ТП.", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "75. Успешный возврат средств, ошибочно некорректно списанных с ББК", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "76. Неуспешный возврат средств, ошибочно некорректно списанных с ББК", //Не описано
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2003965",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
}
