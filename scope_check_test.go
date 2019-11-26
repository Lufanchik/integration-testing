package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

func TestScopeCheckPass(t *testing.T) {
	Run(t, casesScopeCheckPass)
}

var (
	casesScopeCheckPass = Cases{
		{
			N: "Запрет на проезд в МГТ для CUP",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MGT,
					ExpectedSum: 4200,
					card: &processing.Card{
						System: 1,
						Type:   1,
						Pan:    "C8421CE3207A01F74EF458196493E3BC1A5E1C4F867427A453E2148C9300DC0E",
						Bin:    81719999,
						Exp:    "3010",
						Emv:    "ZcwPGStKUWN5YUGbwoxTLaptyRM4+tYhqb1dUKPwpJK5lYBXN0w3Kv2XuleiGyu/yJV3ffgsQp+UkQCPw/vR0g7OoZTT7X87umEPIqVaxz0kf8CG0niv3N6Doxd4Clpc8aRldkwQtJNHemP338BwvBNmgAzMfHkSbrsGyWnXdOwENR3CqOdCXvC0I2AuK9OsNQ64hmJL19SNo91USNY3cv0eeLdEK+xsynB2riUaV/R1KSSRVjAuT/tG99IWvwGfoijEA7to6428u5SlAhaq4CNMIaZGS/5NKk0BVF8UHIb+CcF7Ui4E0MgY7vq9WTpHz7OwZS5ywC5JyTEvL8cjsQ==",
					},
				},
			},
		},
		{
			N: "Запрет прохода для дебетовой ББК МПС CUP,  Эмитент из РФ",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MGT,
					ExpectedSum: 4200,
					card: &processing.Card{
						System: processing.CardSystem_MIR,
						Type:   processing.CardType_UNKNOWN_TYPE,
						Pan:    "C8421CE3207A01F74EF458196493E3BC1A5E1C4F867427A453E2148C9300DC0E",
						Bin:    81719999,
						Exp:    "3010",
						Emv:    "ZcwPGStKUWN5YUGbwoxTLaptyRM4+tYhqb1dUKPwpJK5lYBXN0w3Kv2XuleiGyu/yJV3ffgsQp+UkQCPw/vR0g7OoZTT7X87umEPIqVaxz0kf8CG0niv3N6Doxd4Clpc8aRldkwQtJNHemP338BwvBNmgAzMfHkSbrsGyWnXdOwENR3CqOdCXvC0I2AuK9OsNQ64hmJL19SNo91USNY3cv0eeLdEK+xsynB2riUaV/R1KSSRVjAuT/tG99IWvwGfoijEA7to6428u5SlAhaq4CNMIaZGS/5NKk0BVF8UHIb+CcF7Ui4E0MgY7vq9WTpHz7OwZS5ywC5JyTEvL8cjsQ==",
					},
				},
			},
		},
		{
			N: "Запрет прохода по ББК, BIN которой находится в Стоп-листе",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,

					card: &processing.Card{
						Bin: CardBin(),
					},
				},
			},
		},
		{
			N: "25. Одиночная поездка ММ",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "26. Одиночная поездка МЦК",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "27. Одиночная поездка ММТС",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOffline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "28. Одиночная поездка за пределами МСК (Вход МЦД МО – Выход МЦД МСК)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 1,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: 2,
					},
				},
			},
		},
		{
			N: "29. Одиночная поездка за пределами МСК (Вход МЦД МСК – Выход МЦД МО)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 1,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Direction: 2,
					},
				},
			},
		},
		{
			N: "30. Одиночная поездка МЦД (Вход МЦД МСК – валидация МЦД МО)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 1,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Direction: 0,
					},
				},
			},
		},
		{
			N: "31. Одиночная поездка МЦД (МЦД МО валидация –  МЦД МО выход)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 0,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Direction: 2,
					},
				},
			},
		},
		{
			N: "32. Одиночная поездка МЦД (МЦД МО валидация –  МЦД МО валидация)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 0,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Direction: 0,
					},
				},
			},
		},
		{
			N: "33. Информация о выходе по МЦД пришла после закрытия периода комплексной поездки, но до списания комиссионного сбора.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 0,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
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
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 0,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
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
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: 2,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
				},
			},
		},
		{
			N: "38. Попытка повторного прохода для ММ на одном терминале.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
				},
			},
		},
		{
			N: "39. Повторный проход по для ММ на разных терминалах.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Station: "456",
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					Terminal: &processing.Terminal{
						Station: "457",
					},
				},
			},
		},
		{
			N: "40. Комплексная поездка во время перехода на новые транспортные сутки.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 1,
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					ExpectedSum: 4200,
					Terminal: &processing.Terminal{
						Direction: 2,
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
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
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
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
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
			},
		},
		{
			N: "42. Повторение вида поездки внутри комплексной поездки (МЦК – ММ -МЦК)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
			},
		},
		{
			N: "43. Дублирование звеньев в пределах МСК на МЦД в пределах разных линий (МЦД линия 1 – МЦД линия 2)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					ExpectedSum: 4900,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					ExpectedSum: 0,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					ExpectedSum: 0,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					ExpectedSum: 0,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
					},
				},
			},
		},
		{
			N: "44. Закрытие комплексной поездки по максимальному количеству пересадок в пределах МСК (ММТС – ММ – МЦК – МЦД МСК вход – МЦД МСК выход)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					ExpectedSum: 0,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					ExpectedSum: 0,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
					},
				},
			},
		},
		{
			N: "45. Закрытие комплексной поездки по максимальному количеству пересадок за пределами МСК (ММТС – МЦК – ММ – МЦД МСК вход – МЦД МО выход)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD2_MSK,
					ExpectedSum: 0,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
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
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 0,
				},
			},
		},
		{
			N: "47. Поездка из Аэропорта Шереметьево на Аэроэкспрессе без информации о выходе (Аэропорт вход – информации о выходе не последовало)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 0,
				},
			},
		},
		{
			N: "48. Поездка из Аэропорта Шереметьево на Аэроэкспрессе и выход на МЦД МО (Аэропорт вход – МЦД МО выход)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 0,
				},
			},
		},
		{
			N: "57. Открытие периода комплексной поездки при отсутствии прохода в интервале [-90 мин;+90 мин] относительно полученного прохода. (ММ – МЦК - ММ)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574425800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574427800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574431800000000000,
						},
					},
				},
			},
		},
		{
			N: "58. Продление периода комплексной поездки по параметру времени посадки для поездки по МЦД. (МЦД МО вход – МЦД МСК выход – ММ - МЦК)",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MO,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
						Station:   "2002077",
					},
					ExpectedSum: 4900,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574434800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
						Station:   "2000275",
					},
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574438400000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428600000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428000000000000,
						},
					},
				},
			},
		},
		{
			N: "59. В ТП пришел проход, ставший пересадкой внутри комплексной поездки, не изменив следующую поездку.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
						Station:   "2000700",
					},
					ExpectedSum: 4900,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574425800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
						Station:   "2003965",
					},
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574427800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428500000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428000000000000,
						},
					},
				},
			},
		},
		{
			N: "60. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574425800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428500000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428000000000000,
						},
					},
				},
			},
		},
		{
			N: "61. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
						Station:   "2000700",
					},
					ExpectedSum: 4900,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574425800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
						Station:   "2003965",
					},
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574427800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428500000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574425500000000000,
						},
					},
				},
			},
		},
		{
			N: "62. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_INGRESS,
						Station:   "2000700",
					},
					ExpectedSum: 4900,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574425800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
						Station:   "2003965",
					},
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574427800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428500000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428100000000000,
						},
					},
				},
			},
		},
		{
			N: "64. В ТП пришел проход, ставший новой поездкой, не изменив следующую комплексную поездку.",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MCD,
					SubCarrier:  carriers.SubCarrier_MCD1_MSK,
					Terminal: &processing.Terminal{
						Direction: processing.TerminalDirection_EGRESS,
						Station:   "2003965",
					},
					ExpectedSum: 4900,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574427800000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 0,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428500000000000,
						},
					},
				},
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					tapRequest: &processing.TapRequest{
						Tap: &processing.Tap{
							Created: 1574428100000000000,
						},
					},
				},
			},
		},
	}
)
