package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	webApi "lab.siroccotechnology.ru/tp/web-api-gateway/proto"
)

func PassOfflineRequest(tap *processing.TapRequest, p *Pass) (*processing.OfflinePassRequest, *processing.OfflinePassResponse) {
	request := &processing.OfflinePassRequest{
		Id:      tap.Id,
		Created: Now(),
		Tap:     tap.Tap,
	}
	request.Auth = auth(p)

	response := &processing.OfflinePassResponse{
		Result: processing.PassStatus_SUCCESS,
	}
	return request, response
}

func logRequest(u string, r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("%s, trace id: %s", u, trace.Raw()))
}

func auth(p *Pass) *processing.Auth {
	if p.Auth != nil {
		return p.Auth
	} else {
		na := &processing.Auth{
			Sum:  p.ExpectedSum,
			Type: processing.Auth_CLASSIC,
		}
		if p.PaymentType == PaymentTypeStartAggregate {
			na.Type = processing.Auth_AGGREGATE
		}
		return na
	}
}

func PassOnlineRequest(tap *processing.TapRequest, p *Pass) (*processing.OnlinePassRequest, *processing.OnlinePassResponse) {
	request := &processing.OnlinePassRequest{
		Id:      tap.Id,
		Created: Now(),
		Tap:     tap.Tap,
	}

	request.Auth = auth(p)

	responseOR := &processing.OnlinePassResponse{
		Created: 0,
		Result:  processing.PassStatus_SUCCESS,
	}
	switch p.PaymentType {
	case PaymentTypePayment:
		responseOR.Status = processing.AuthStatus_SUCCESS_AUTH
	case PaymentTypeFree:
		responseOR.Status = processing.AuthStatus_SUCCESS_FREE
	case PaymentTypeStartAggregate:
		responseOR.Status = processing.AuthStatus_SUCCESS_AUTH
	case PaymentTypeAggregate:
		responseOR.Status = processing.AuthStatus_SUCCESS_AGGREGATE
	}

	if p.AuthType == AuthTypeIncorrect && p.PaymentType != PaymentTypeStartAggregate && p.PaymentType != PaymentTypeAggregate {
		responseOR.Status = processing.AuthStatus_FAILURE_ISSUER
	}

	if p.Duration > 0 {
		request.Timeout = p.Duration
	}

	return request, responseOR
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

func WebAPIRequest(card *processing.Card) *webApi.PassesRequest {
	return &webApi.PassesRequest{
		Hash: card.Pan,
	}
}

func CompleteRequest(pass *Pass, passes []*Pass, sum int) (*processing.CompleteRequest, *processing.CompleteResponse) {
	request := &processing.CompleteRequest{
		Id:      pass.id,
		Created: Now(),
		Amount:  uint32(sum),
	}
	var ids []string
	for _, v := range passes {
		ids = append(ids, v.id)
	}

	request.Pass = ids

	response := &processing.CompleteResponse{
		Result: processing.CompleteResponse_SUCCESS,
	}

	return request, response
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

	if isAggregate(p) {
		response.Auth = &processing.Auth{}
	}

	return request, response
}

func isAggregate(p *Pass) bool {
	return p.PaymentType == PaymentTypeStartAggregate || p.PaymentType == PaymentTypeAggregate
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
