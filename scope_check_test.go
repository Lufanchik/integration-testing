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
					PaymentType: PaymentTypeFullPayment,
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
	}
)
