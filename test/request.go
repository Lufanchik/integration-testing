package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/google/uuid"
	authService "lab.dt.multicarta.ru/tp/common/messages/auth"
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/comments"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/messages/response"
	"lab.dt.multicarta.ru/tp/common/messages/twpg"
	webApi "lab.dt.multicarta.ru/tp/web-api-gateway/proto"
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
			Sum:  getSumByCarrier(p),
			Type: processing.Auth_CLASSIC,
		}
		if p.PaymentType == PaymentTypeStartAggregate {
			na.Type = processing.Auth_AGGREGATE
			na.Sum = 0
		}
		return na
	}
}

func getSumByCarrier(p *Pass) uint32 {
	if p.ExpectedSum > 0 {
		return p.ExpectedSum
	}
	if p.PaymentType == PaymentTypeFree {
		return 0
	}
	if p.SubCarrier == carriers.SubCarrier_MM_SUB {
		return 4400
	}
	if p.SubCarrier == carriers.SubCarrier_MCK_SUB {
		return 4400
	}
	if p.SubCarrier == carriers.SubCarrier_MMTS_SUB {
		return 4400
	}
	if p.SubCarrier == carriers.SubCarrier_MCD1_MSK {
		return 4200
	}
	if p.SubCarrier == carriers.SubCarrier_MCD2_MSK {
		return 4200
	}
	if p.SubCarrier == carriers.SubCarrier_MCD1_MO {
		return 4900
	}
	if p.SubCarrier == carriers.SubCarrier_MCD2_MO {
		return 4900
	}
	if p.Carrier == carriers.Carrier_MM {
		return 4400
	}
	if p.Carrier == carriers.Carrier_MGT {
		return 4400
	}
	if p.Carrier == carriers.Carrier_MTPPK {
		return p.ExpectedSum
	}
	return 999
}

func GetPassType(p *Pass) pass.PassType {
	if p != nil {
		return p.PassType
	}
	return pass.PassType_PASS_BBK
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
	case PaymentTypePrepayed:
		responseOR.Status = processing.AuthStatus_NONE_AUTH

	}

	if p.isComplete {
		switch p.PaymentType {
		case PaymentTypeStartAggregate:
			responseOR.Status = processing.AuthStatus_SUCCESS_AUTH
			if p.AuthType == AuthTypeIncorrect {
				responseOR.Status = processing.AuthStatus_FAILURE_ISSUER
			}
		case PaymentTypeAggregate:
			responseOR.Status = processing.AuthStatus_SUCCESS_AGGREGATE
		}
	}

	if p.AuthType == AuthTypeIncorrect && p.PaymentType != PaymentTypeStartAggregate {
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

func WebAPIPassesRequest(card *processing.Card) *webApi.PassesRequest {
	return &webApi.PassesRequest{
		Hash: card.Pan,
	}
}

func WebAPIRegisterRequest(fcl *RegisterFaceId, card *processing.Card) *authService.FaceIdRegisterRequest {
	return &authService.FaceIdRegisterRequest{
		Id: card.Pan,
		Urls: &authService.TWPGUrls{
			Approve: gofakeit.URL(),
			Cancel:  gofakeit.URL(),
			Decline: gofakeit.URL(),
		},
	}
}

func WebAPIFaceStatusRequest(faceCheck *FaceIdRegistrationStatus) (*twpg.RegistrationStatusRequest, *twpg.RegistrationStatusResponse) {
	req := &twpg.RegistrationStatusRequest{
		Id: faceCheck.Id,
	}

	resp := &twpg.RegistrationStatusResponse{
		Status: twpg.RegistrationStatusResponse_SUCCESS,
		Card: &twpg.CardInfo{
			Pan: "4150********0900",
			Exp: "2101",
		},
	}

	return req, resp
}

func AddCommentRequest(crud *CommentsCRUD) *comments.AddCommentRequest {
	request := &comments.AddCommentRequest{
		Comment: &comments.Comment{
			Id:        uuid.New().String(),
			EntityId:  uuid.New().String(),
			UserId:    uuid.New().String(),
			RepliesTo: "",
			Body:      gofakeit.HackerPhrase(),
		},
	}

	return request
}

func GetCommentsRequest(add *comments.AddCommentRequest) *comments.GetCommentsRequest {
	request := &comments.GetCommentsRequest{
		EntityId: add.Comment.EntityId,
	}

	return request
}

func DelCommentRequest(c *comments.Comment) *comments.DeleteCommentRequest {
	request := &comments.DeleteCommentRequest{
		CommentId: c.Id,
	}

	return request
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

func FaceForceCheckRequest() (*response.EmptyMessage, *response.EmptyMessage) {
	return &response.EmptyMessage{}, &response.EmptyMessage{}
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
			Sum:  getSumByCarrier(p),
			Type: processing.Auth_CLASSIC,
		}
		response.Resolution = processing.AuthResponse_AUTHORIZED
	case PaymentTypeFree:
		response.Status = processing.AuthResponse_SUCCESS_FREE
	case PaymentTypeStartAggregate:
		response.Status = processing.AuthResponse_SUCCESS_AGGREGATE
		response.Auth = &processing.Auth{
			Sum:  0,
			Type: processing.Auth_AGGREGATE,
		}
	case PaymentTypeAggregate:
		response.Status = processing.AuthResponse_SUCCESS_AGGREGATE
		response.Auth = &processing.Auth{
			Sum:  getSumByCarrier(p),
			Type: processing.Auth_AGGREGATE,
		}
	}

	if p.isComplete {
		switch p.PaymentType {
		case PaymentTypeStartAggregate:
			response.Status = processing.AuthResponse_SUCCESS_STATUS
			response.Auth = &processing.Auth{
				Sum:  getSumByCarrier(p),
				Type: processing.Auth_AGGREGATE,
			}
			response.Resolution = processing.AuthResponse_AUTHORIZED
		case PaymentTypeAggregate:
			if p.aggregate != nil && p.aggregate.AuthType == AuthTypeCorrect {
				response.Status = processing.AuthResponse_SUCCESS_STATUS
				response.Auth = &processing.Auth{
					Sum:  getSumByCarrier(p),
					Type: processing.Auth_AGGREGATE,
				}
				response.Resolution = processing.AuthResponse_AUTHORIZED
			}
		}
	}

	if p.AuthType == AuthTypeIncorrect && !isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	if p.AuthType == AuthTypeIncorrect && p.isComplete && isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
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
