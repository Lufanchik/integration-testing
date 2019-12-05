package passes

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"time"
)

var CasesScopeCheckPass = test.Cases{
	{
		N: "25. Одиночная поездка ММ",
		T: test.T{
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
		N: "26. Одиночная поездка МЦК",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
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
				ExpectedSum: 4200,
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
		N: "31. Одиночная поездка МЦД (МЦД МО валидация –  МЦД МО выход)",
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
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Station:   "2000115", //ЛОБНЯ
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress: 1,
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
			},
		},
	},
	{
		N: "33. Информация о выходе по МЦД пришла после закрытия периода комплексной поездки, но до списания комиссионного сбора.",
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
		N: "34. Комиссионный сбор при неосуществлении выхода с МЦД (МЦД МСК - выхода с МЦД не последовало)",
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
					Station:   "2000055",
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
					Station:   "2000155",
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
		N: "38. Попытка повторного прохода для ММ на одном терминале.",
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
				Now:         test.NowFullDate(2019, 12, 2, 3, 30, 30),
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
		N: "40. Комплексная поездка во время перехода на новые транспортные сутки.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Id:        "sirocco",
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				ExpectedSum: 4200,
				Terminal: &processing.Terminal{
					Id:        "multicarta",
					Direction: processing.TerminalDirection_EGRESS,
				},
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
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
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
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
					Direction: 1,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: 2,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: 1,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: 2,
				},
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
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: 1,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: 2,
				},
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
				ExpectedSum: 0,
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
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				ExpectedSum: 0,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
				},
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				ExpectedSum: 700,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
				},
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
		N: "47. Поездка из Аэропорта Шереметьево на Аэроэкспрессе без информации о выходе (Аэропорт вход – информации о выходе не последовало)",
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
		N: "48. Поездка из Аэропорта Шереметьево на Аэроэкспрессе и выход на МЦД МО (Аэропорт вход – МЦД МО выход)",
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
		N: "58. Продление периода комплексной поездки по параметру времени посадки для поездки по МЦД. (МЦД МО вход – МЦД МСК выход – ММ - МЦК)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MO,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_INGRESS,
					Station:   "2002077",
				},
				ExpectedSum: 4900,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD1_MSK,
				Terminal: &processing.Terminal{
					Direction: processing.TerminalDirection_EGRESS,
					Station:   "2000275",
				},
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 0,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "59. В ТП пришел проход, ставший пересадкой внутри комплексной поездки, не изменив следующую поездку.",
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
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
		},
	},
	{
		N: "60. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
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
		N: "61. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
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
		N: "62. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
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
		N: "64. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
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
