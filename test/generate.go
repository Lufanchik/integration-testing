package test

import (
	"github.com/brianvoe/gofakeit"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	passService "lab.dt.multicarta.ru/tp/pass-service/proto"
	"net/http"
	"strconv"
	"time"
)

func init() {
	ps = passService.NewPassServiceProtobufClient(PassUrl, http.DefaultClient)
	gofakeit.Seed(time.Now().UnixNano())
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

//func MasterCardTokenEmvCorrect() string {
//	return "ufh9wgp7m4aSIqF/aMXqdjzvAFUO6pa9rHKTNjnaqTV3KdPKMPDEAjzqpZH6999H7Kh6YeyG5ix+HEJCLXizhxNvQ21LHHQwbM5df0ryWfNiDpyljLkxLI5dMl9qnPAYgaSPzswcqYQp9hlTbJdxp1ApDNhgTNHIQSsyiXxStIVqTVaGZuqias2SXFBTT6iaGv4wjyo6q600Otit1fGUQL4YcFu9eTWgnSMlKHJBMScLoFKx2uQkQOD2q7Ru4SBV7snTqJtOiPYPoW1Ed+3/iyvt3RwAA71dAp+vskHT3RDMbe/qO5Me4aMfs5H0mdiZ5HxIOHA1oIRnhApHEk/VhA=="
//}
//func MasterCardTokenEmvIncorrect() string {
//	return "cEYXgoriIwrqygRbnY7uHYOypt0GepxG/FD+1XRKKA6xYnbGHIoi3oCt7Tp+oDFdGYKqxATJ4mt+T9+fMTxOBCf/0Gclzsmmj8EJZ2sscMlnLTzwwXLDPd/AnHKiE+ojLoAZupxjCkxJ23BFmTz1f1IWSrqxBKgXItMZLQXUYKgFlVLH7/5fmDXTS7+3dF3ncr894aU0PAPpUUE5N2xXNll36aKOusLWc4op/BkqRu0GZohheol+xc4PF6+ZiyTC+it9RcmriF5gm79dxASQhd03oSDFaqiwdsm4xlk1T5dg9tTJFj60zLkM6Q9Ylu34tUQLdS1WDVShcNH6K1+1FA=="
//}

func Card(system processing.CardSystem) *processing.Card {
	gofakeit.Seed(time.Now().UnixNano())
	card := &processing.Card{
		System: processing.CardSystem_VISA,
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
		System: processing.CardSystem_VISA,
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
		System: processing.CardSystem_VISA,
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
	}
}
