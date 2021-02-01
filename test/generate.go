package test

import (
	"github.com/brianvoe/gofakeit"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/rabbit"
	passService "lab.dt.multicarta.ru/tp/pass-service/proto"
	"net/http"
	"strconv"
	"time"
)

var rpa *rabbit.Exchange

func init() {
	ps = passService.NewPassServiceProtobufClient(PassUrl, http.DefaultClient)
	gofakeit.Seed(time.Now().UnixNano())

	//rabbitConfig := rabbit.Config{
	//	ServiceName: "integration-test",
	//	Endpoint:    "amqp://guest:guest@localhost:5673/",
	//}
	//
	//rabbitPrimary := rabbit.MustNew(rabbitConfig)
	//authExchange := exchanges.MustExchangeByID(exchanges.ExchangeAuth)
	//rpa = rabbitPrimary.MustNewExchange(authExchange.Opts)
}

func CardExp() string {
	gofakeit.Seed(time.Now().UnixNano())
	return strconv.Itoa(gofakeit.Number(1000, 9999))
}

func CardBin() uint32 {
	gofakeit.Seed(time.Now().UnixNano())
	return uint32(gofakeit.Number(1000, 9999))
}

func CardPan() string {
	gofakeit.Seed(time.Now().UnixNano())
	return strconv.Itoa(gofakeit.CreditCardNumber())
}

func CardEmvCorrect() string {
	return "ca5zfeN8X0VQGfocwEl/8WI1IGyfrOPVST9rtzVBBpPDYVKm/f1/r+gyPwRBANXrWt13q/ADa6VEoQU8u+Og0FPk2IzepajqqPEpxzHGZANjE5fygv7Hk/kblwu6Ktxj76AhU3Te1nlr5LhrfcsOLr3LbEfr2YXOi8GRX2FC/AHNJRumHCyF6r7aBB4EwoAM/gBk53y1TIvM84eq7G4b+z4w0p/le+FXb4DzuryOsj633DVEaWtMbv4A+HgqpdubQyFizkbDmiRLlcEkXztuqEJX/c1jFrR4LhA4/gU9YPcd0YggZQ53gqJ8b57HljbVAosjwuvWE4JaG4sF71sRUg=="
}
func CardEmvIncorrect() string {
	return "WovYAaagxXfxea2U/5cMzTN1MEQIJZAUTovU5NDItD4qB6PyP0rzjg3nnXjKH3XVaVMjUSHJ5YiIJ2Ji4jX3E5bj+Ufe5BntEUhjqyphe8HQz+jJdhWfld+Bm61C8yeeq3qhYbf7zgswrh3d2Gd5L6h4PlKVbbCuGLI8KOmTmfoqQXX5dtF+ZHum7l5BUIvGn3nFj1Fbkye0iKcxHXyToZd/l/M9FuRy9/klAgKYPScYlYWRSwH2I5HZs5qDKJi/cXscLjYoF6h9xxRJMsXJr68BQ5E1bx9sG5mlBb3Pzeytns8Qct9pziuUdDUnFEd2xgx07ul7jNu40k9BARcyfQ=="
}
func CardEmvRefund() string {
	return "DYloYwM7siugRNTwXR+X844n1gM0vxQdmRloE3tREdgIMvJMHNzkK6G0P55vYSqLa7W3FkbjL0Rwu2XyiqFmumoyGAY/cMxMgHTVDCp7wuEpkPZo8PhRDwISBcvR1Roh8bh2omooXgSnjI2QvQFvV+bcjc3uI2XJ84mHnkOyQzO+pS/Ow67PykL8W7RLYsLDhnoa4556p7aUWKhPlaxxVs0M2HtermZZfsFnjRR0FpNiCxQvj28kHNjK2oiw/rqP0b6eBsWwBCZiTuhptJ99QJCtQ7zBA4iTk91Wze/qZgMyWrDNYOEOw9jW1wguDwsPHb5Bk2QrLCARfZgwkFjnLQ=="
}

func TokenMCCorrect() string {
	return "ufh9wgp7m4aSIqF/aMXqdjzvAFUO6pa9rHKTNjnaqTV3KdPKMPDEAjzqpZH6999H7Kh6YeyG5ix+HEJCLXizhxNvQ21LHHQwbM5df0ryWfNiDpyljLkxLI5dMl9qnPAYgaSPzswcqYQp9hlTbJdxp1ApDNhgTNHIQSsyiXxStIVqTVaGZuqias2SXFBTT6iaGv4wjyo6q600Otit1fGUQL4YcFu9eTWgnSMlKHJBMScLoFKx2uQkQOD2q7Ru4SBV7snTqJtOiPYPoW1Ed+3/iyvt3RwAA71dAp+vskHT3RDMbe/qO5Me4aMfs5H0mdiZ5HxIOHA1oIRnhApHEk/VhA=="
}
func TokenMCIncorrect() string {
	return "cEYXgoriIwrqygRbnY7uHYOypt0GepxG/FD+1XRKKA6xYnbGHIoi3oCt7Tp+oDFdGYKqxATJ4mt+T9+fMTxOBCf/0Gclzsmmj8EJZ2sscMlnLTzwwXLDPd/AnHKiE+ojLoAZupxjCkxJ23BFmTz1f1IWSrqxBKgXItMZLQXUYKgFlVLH7/5fmDXTS7+3dF3ncr894aU0PAPpUUE5N2xXNll36aKOusLWc4op/BkqRu0GZohheol+xc4PF6+ZiyTC+it9RcmriF5gm79dxASQhd03oSDFaqiwdsm4xlk1T5dg9tTJFj60zLkM6Q9Ylu34tUQLdS1WDVShcNH6K1+1FA=="
}

func MKEmulatorSuccess() string {
	return "GmbAJzvwhXYTskIojGZoyPZg2APDklnLmeJe043iKmECaQM2d2osjhp6w4ggJESeV1H4praqEXQW62DvkbbcwmBmbaNNHA+0v9aqWpy7x/+qQ56TfIjRbJI9Dp0HeONLN0LkADMAuPInx+vOSLO0g+QnyKOUVLIwFg3v7ZyNNPBWOS5Y05ujzrCMlZL+3kdN4rqmtEwQEN+Yijy4hGLowYetPf4r9MKFrmGyV3u/ufRJFFdRWgDYhr/swYrulA5wfFGzbVbCnCUwa+7tXsmQ07GCHI5xDHbkeBa/L8vu4sZfyt1HuQqog8dN79L945z8BH/ty5me5BR72u66NKJlNw=="
}
func MKEmulatorUnsuccess() string {
	return "DN1QUvKcya2IPuBQPGI6vR2zCqmfVcPCbOYn4Kts12gR/GlsJtEEwA/UR/iwTe/0AwmSb8nNtm4iTb26sNK32LrBeF8UevdTWNyC/USYGmbM2FnGtUS3OLzzXQRHGCizhpy1Qmi5JXKma7TKILUlH/PN2FG5f0eowK3Y+DljmseWQFiFlIE6M/jBrHUph8V/D4+ZFtphOc/A356r1W7qBVz09jbpACQ1UzIP5ufYhwB3G1cwWcmMXrGOmbCnolZkVY4FYUbZO2pNHv8JBZon29RRznPEKK/dcs3JJEoTO287X0OM7uC1hnFMrl05x/ymtbMVP6053QkMoL+kMN+l0A=="
}
func MKEmulatorUnsuccessWithReauth() string {
	return "b2X1GbUBWsFcfaLtjS9myMw94DLWYPKgiVgkP7rxkm/N4H6wqZbe/mfmhVIX+rjU3M1ISY9zgpFHUkqNrZLTFv7UlG3/gARhZ4fB+hrJdv7mfWfnAfbm7OknmRZiJpZzrcOiLJAhHaSnA75L1JzZdqC0gySK7qdW9nJTlCxvotNIxvztxAIb4Y9YUE/X0F5HDH4Im0yXzWa9RZB1hekxKkask5mncqTO8SHqt15btJDXoJUogFDFGkGCYhUfQJ8qNYc++Y+77ZoSniaN0bBkTR2gYwVkgB/0RXTYWxbj7U5bocNH3ztHNcc9TGHdPzyzM2Jjiy+DvLy7w0uM2uRVGQ=="
}
func MKEmulatorUnsuccessWithoutReauth() string {
	return "XOmg79U0VmKQdPWX58RWHpQsBWEJo2ELMettKURYIgLUF/BdAmIrAB4NJYyYqkbRhtzlU/ZGYq1ppdCZz7UjKs3GFf9c8Zt4dXMwicr33kaJtnRE36vQMWSYtgkVghCqYJ07S/A6LMq9WCJEHh9Jcz6SOyiOy0r4mbEbFzBIeLKgU0kN15rnxPhVG799uMFbieyFzxIswltKVtZK0sR+aNEef+iXGmATvz5bqKUlKtXBO6MZKVdJXt1Rtyem8hPjtUyuXUsxmklKvMEnYXySuAHxeXUXLTLc/lQPLrBhcrTMAzmCPHdEvqaHOG+XB1cEketZbgNENVLG03e7+khLOw=="
}
func MKEmulatorRandom() string {
	return "VB/kqiqfJnQ1S5TH0mq6fFWPi7uyKmGPbSgHRyUj4jjEvSOGkGFrjNUNgaVXPxScQhO2EYyQa180v6hSdB4CBUoevoiVTSt9Vr5pEX/FZJf94EEJVrFmYQtPXRjNaqckbuiVkKkjhEVP2TAdiEvqZSqzFcMTJuRBDWIzmNajGoRQz/seehttEDQe7qhldlvnIEoLrwZ9gFcv9Q352KlFQALWgQ5O8g8w7Jhr5xtV7AYGYqFxOPTrI54d6VnYzg4MHkHUxKDkYZhYFVbkPSdi8HQnHn9KG3OvZY/GjIH0c0Lg4Yem5lRdSjCHZKKEzRA3l+UPSfYabMocuOq3VMgFHA=="
}

func Card(system processing.CardSystem) *processing.Card {
	gofakeit.Seed(time.Now().UnixNano())
	card := &processing.Card{
		System: processing.CardSystem_MIR,
		Type:   processing.CardType_DEBIT,
		Pan:    CardPan(),
		Bin:    CardBin(),
		Exp:    CardExp(),
		Token: &processing.Token{
			Type: processing.Token_SAMSUNG_PAY,
		},
	}
	if system != processing.CardSystem_NONE_SYSTEM {
		card.System = system
	}
	return card
}

func FaceCard(system processing.CardSystem, pan string) *processing.Card {
	gofakeit.Seed(time.Now().UnixNano())
	card := &processing.Card{
		System: processing.CardSystem_MIR,
		Type:   processing.CardType_FACEID,
		Pan:    pan,
		Bin:    CardBin(),
		Exp:    CardExp(),
		Token: &processing.Token{
			Type: processing.Token_SAMSUNG_PAY,
		},
	}
	if system != processing.CardSystem_NONE_SYSTEM {
		card.System = system
	}
	return card
}

func MTCard(system processing.CardSystem) *processing.Card {
	gofakeit.Seed(time.Now().UnixNano())
	card := &processing.Card{
		System: processing.CardSystem_MIR,
		Type:   processing.CardType_MT,
		Pan:    CardPan(),
		Bin:    CardBin(),
		Exp:    CardExp(),
		Token: &processing.Token{
			Type: processing.Token_SAMSUNG_PAY,
		},
	}
	if system != processing.CardSystem_NONE_SYSTEM {
		card.System = system
	}
	return card
}

func GenerateEmv(card *processing.Card, p *Pass) {
	switch p.AuthType {
	case AuthTypeCorrect:
		card.Emv = CardEmvCorrect()
	case AuthTypeIncorrect:
		card.Emv = CardEmvIncorrect()
	case AuthTypeRefund:
		card.Emv = CardEmvRefund()
	case AuthTypeMCTokenCorrect:
		card.Emv = TokenMCCorrect()
	case AuthTypeMCTokenIncorrect:
		card.Emv = TokenMCIncorrect()
	case AuthTypeMKEmulatorSuccess:
		card.Emv = MKEmulatorSuccess()
	case AuthTypeMKEmulatorUnsuccess:
		card.Emv = MKEmulatorUnsuccess()
	case AuthTypeMKEmulatorRandom:
		card.Emv = MKEmulatorRandom()
	case AuthTypeUnsuccessWithReauth:
		card.Emv = MKEmulatorUnsuccessWithReauth()
	case AuthTypeUnsuccessWithoutReauth:
		card.Emv = MKEmulatorUnsuccessWithoutReauth()
	}
}
