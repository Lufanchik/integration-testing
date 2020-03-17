package parking

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var (
	CasesParkingPass = test.Cases{
		{
			N: "1.Parking_AuthTypeCorrect",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					AuthType:    test.AuthTypeCorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
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
				&test.Parking{
					RP: processing.CheckParkingResponse_SUCCESS,
					R: &processing.CheckParkingRequest{
						Stations: []string{"2", "1234"},
					},
				},
				&test.Parking{
					RP: processing.CheckParkingResponse_NOT_FOUND,
					R: &processing.CheckParkingRequest{
						Stations: []string{"2"},
					},
				},
			},
		},
		{
			N: "2.Parking_AuthTypeIncorrect",
			T: test.T{
				&test.Pass{
					PaymentType: test.PaymentTypePayment,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					AuthType:    test.AuthTypeIncorrect,
					Terminal: &processing.Terminal{
						Id:         "1",
						Station:    "1234",
						Direction:  processing.TerminalDirection_INGRESS,
						SubCarrier: carriers.SubCarrier_MM_SUB,
					},
				},
				&test.Pass{
					PaymentType: test.PaymentTypeFree,
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
				&test.Parking{
					RP: processing.CheckParkingResponse_SUCCESS,
					R: &processing.CheckParkingRequest{
						Stations: []string{"2", "1234"},
					},
				},
				&test.Parking{
					RP: processing.CheckParkingResponse_NOT_FOUND,
					R: &processing.CheckParkingRequest{
						Stations: []string{"2"},
					},
				},
			},
		},
	}
)
