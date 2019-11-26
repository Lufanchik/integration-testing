package http

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

package http_test

import (
"lab.siroccotechnology.ru/tp/common/messages/carriers"
"lab.siroccotechnology.ru/tp/common/messages/processing"
"testing"
)

func TestComplexMCD(t *testing.T) {
	Run(t, casesComplexMCD)
}

var (
	casesComplexMCD = Cases{
//gusmanov test case
	{
N: "MCK - MM - MM", //1
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MM_SUB,
ExpectedSum: 4200,
},
},
},

{
N: "MCK - MM - MCK", //2
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
},
},

{
N: "MCK - MM - MMTS - MM", //3
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MMTS_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MM_SUB,
ExpectedSum: 1,
},
},
},

{
N: "MCK - MM - MMTS - MCK", //4
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MMTS_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier: carriers.Carrier_MM,
SubCarrier: carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
},
},

{
N: "MCK - MM - MMTS - MMTS", //5
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MMTS_SUB,
Parent: 1,
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
N: "MCK - MM - MMTS - MCD_MSK", //6
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MMTS_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000275", //СЕТУНЬ
Direction: processing.TerminalDirection_INGRESS,
},
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001060", //БЕГОВАЯ
Direction: processing.TerminalDirection_EGRESS,
},
Ingress: 4,
},
},
},

{
N: "MCK - MM - MMTS - MCD_MO", //7
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MMTS_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2002077", //марк
Direction: processing.TerminalDirection_INGRESS,
},
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MO,
Terminal: &processing.Terminal{
Station: "2000115", //лобня
Direction: processing.TerminalDirection_EGRESS,
},
Ingress: 4,
ExpectedSum: 700,
},
},
},

{
N: "MCK - MM - MCD_MSK - MM - MM", //8
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000245", //рабочий поселок
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
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
N: "MCK - MM - MCD_MSK - MM - MCK", //9
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001140", //Кунцевская (бывш. КУНЦЕВО 1)
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000455", //ДЕГУНИНО
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
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
N: "MCK - MM - MCD_MSK - MM - MMTS", //10
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000700", //ТЕСТОВСКАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2003965", //ТИМИРЯЗЕВСКАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
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
N: "MCK - MM - MCD_MSK - MM - MCD_MSK", //11
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
},
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000245", //РАБОЧИЙ ПОСЕЛОК
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 6,
},
},
},

{
N: "MCK - MM - MCD_MSK - MM - MCD_MO", //12
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
},
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MO,
Terminal: &processing.Terminal{
Station: "2000115", //ЛОБНЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 6,
ExpectedSum: 700,
},
},
},

{
N: "MCK - MM - MCD_MSK - MCK", //13
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
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
N: "MCK - MM - MCD_MSK - MMTS - MM", //14
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
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
Carrier: 	 carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
ExpectedSum: 4200,
},
},
},

{
N: "MCK - MM - MCD_MSK - MMTS - MCK", //15
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
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
Carrier: 	 carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
},
},

{
N: "MCK - MM - MCD_MSK - MMTS - MMTS", //16
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
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
Carrier: 	 carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MMTS_SUB,
ExpectedSum: 4200,
},
},
},

{
N: "MCK - MM - MCD_MSK - MMTS - MCD_MSK", //17
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
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
Station: "2001060", //БЕГОВАЯ
Direction: processing.TerminalDirection_INGRESS,
},
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2003965", //ТИМИРЯЗЕВСКАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 6,
},
},
},

{
N: "MCK - MM - MCD_MSK - MMTS - MCD_MO", //18
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
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
Station: "2001060", //БЕГОВАЯ
Direction: processing.TerminalDirection_INGRESS,
},
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MO,
Terminal: &processing.Terminal{
Station: "2001775", //ШЕРЕМЕТЬЕВСКАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 6,
ExpectedSum: 700,
},
},
},

{
N: "MCK - MM - MCD_MSK1 - MCD_MSK2 - MM", //19
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_INGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD2_MSK,
Terminal: &processing.Terminal{
Station: "2001425", //КРАСНЫЙ БАЛТИЕЦ
Direction: processing.TerminalDirection_EGRESS,
},
Ingress: 5,
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
N: "MCK - MM - MCD_MSK1 - MCD_MSK2 - MCK", //20
T: T{
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MM_SUB,
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2000155", //ФИЛИ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD1_MSK,
Terminal: &processing.Terminal{
Station: "2001270", //ОКРУЖНАЯ
Direction: processing.TerminalDirection_EGRESS,
},
Ingress: 3,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD2_MSK,
Terminal: &processing.Terminal{
Station: "2000009", //САВЕЛОВСКИЙ ВОКЗАЛ
Direction: processing.TerminalDirection_INGRESS,
},
Parent: 1,
},
&Pass{
PaymentType: PaymentTypeFree,
Carrier:     carriers.Carrier_MCD,
SubCarrier:  carriers.SubCarrier_MCD2_MSK,
Terminal: &processing.Terminal{
Station: "2000605", //Рижская
Direction: processing.TerminalDirection_EGRESS,
},
Ingress: 5,
},
&Pass{
PaymentType: PaymentTypePayment,
Carrier:     carriers.Carrier_MM,
SubCarrier:  carriers.SubCarrier_MCK_SUB,
ExpectedSum: 4200,
},
},
},




}
)
