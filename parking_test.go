package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

var (
	parkingPass = Cases{
		{
			N: "Parking",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&Pass{
					PaymentType: PaymentTypeFree,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					Parent:      1,
					Terminal: &processing.Terminal{
						Id:         "2",
						Station:    "4321",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MMTS_SUB,
					},
				},
				&Parking{
					rp: processing.CheckParkingResponse_SUCCESS,
					r: &processing.CheckParkingRequest{
						Stations: []string{"2", "1234"},
					},
				},
				&Parking{
					rp: processing.CheckParkingResponse_NOT_FOUND,
					r: &processing.CheckParkingRequest{
						Stations: []string{"2"},
					},
				},
			},
		},
	}
)

func TestParking(t *testing.T) {
	Run(t, parkingPass)
}
