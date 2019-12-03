package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
)

func PassOfflineRequest(tap *processing.TapRequest, p *Pass) (*processing.OfflinePassRequest, *processing.OfflinePassResponse) {
	request := &processing.OfflinePassRequest{
		Id:      tap.Id,
		Created: Now(),
		Tap:     tap.Tap,
	}
	if p.Auth != nil {
		request.Auth = p.Auth
	}
	response := &processing.OfflinePassResponse{
		Result: processing.PassStatus_SUCCESS,
	}
	return request, response
}

func logRequest(u string, r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("%s, trace id: %s", u, trace.Raw()))
}

func PassOnlineRequest(tap *processing.TapRequest, p *Pass) (*processing.OnlinePassRequest, *processing.OnlinePassResponse) {
	request := &processing.OnlinePassRequest{
		Id:      tap.Id,
		Created: Now(),
		Tap:     tap.Tap,
	}
	if p.Auth != nil {
		request.Auth = p.Auth
	}
	response := &processing.OnlinePassResponse{
		Created: 0,
		Result:  processing.PassStatus_SUCCESS,
	}
	switch p.PaymentType {
	case PaymentTypePayment:
		response.Status = processing.AuthStatus_SUCCESS_AUTH
	case PaymentTypeFree:
		response.Status = processing.AuthStatus_SUCCESS_FREE
	}

	if p.AuthType == AuthTypeIncorrect {
		response.Status = processing.AuthStatus_FAILURE_ISSUER
	}

	if p.Duration > 0 {
		request.Timeout = p.Duration
	}

	return request, response
}

func CancelRequest(cl *Cancel, p *Pass) (*processing.CancelPassRequest, *processing.CancelPassResponse) {
	request := &processing.CancelPassRequest{
		Id:      p.id,
		Created: Now(),
		Reason:  cl.Reason,
	}

	response := &processing.CancelPassResponse{
		Id:     p.id,
		Result: processing.CancelPassResponse_SUCCESS,
	}

	return request, response
}

func ParkingRequest(card *processing.Card, pr *Parking) (*processing.CheckParkingRequest, *processing.CheckParkingResponse) {
	pr.R.Pan = card.Pan

	response := &processing.CheckParkingResponse{
		Result: pr.RP,
	}

	return pr.R, response
}

func AuthStatusRequest(p *Pass) (*processing.AuthRequest, *processing.AuthResponse) {
	request := &processing.AuthRequest{
		Id: p.id,
	}

	response := &processing.AuthResponse{
		Result: processing.AuthResponse_SUCCESS_RESULT,
	}
	switch p.PaymentType {
	case PaymentTypePayment:
		response.Status = processing.AuthResponse_SUCCESS_STATUS
		response.Auth = &processing.Auth{
			Sum:  p.ExpectedSum,
			Type: processing.Auth_CLASSIC,
		}
		response.Resolution = processing.AuthResponse_AUTHORIZED
	case PaymentTypeFree:
		response.Status = processing.AuthResponse_SUCCESS_FREE
	}

	if p.AuthType == AuthTypeIncorrect {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	return request, response
}

func TapRequest(c carriers.SubCarrier, card *processing.Card, p *Pass) (*processing.TapRequest, *processing.TapResponse) {
	timeNow := Now()
	request := &processing.TapRequest{
		Id:      p.carrierID,
		Created: timeNow,
		Tap: &processing.Tap{
			Created:    timeNow,
			Resolution: processing.TapResolution_APPROVE,
			Terminal: &processing.Terminal{
				Id:         gofakeit.HipsterWord(),
				Station:    gofakeit.BeerName(),
				Direction:  processing.TerminalDirection_INGRESS,
				SubCarrier: c,
			},
			Card: card,
			Sign: gofakeit.BeerStyle(),
		},
	}
	if p.Terminal != nil {
		if p.Terminal.Id == "" {
			p.Terminal.Id = gofakeit.HipsterWord()
		}
		request.Tap.Terminal = p.Terminal
	}
	response := &processing.TapResponse{
		Id:      "",
		Created: 0,
		Result:  processing.TapResponse_SUCCESS,
		Msg:     "",
	}
	return request, response
}
