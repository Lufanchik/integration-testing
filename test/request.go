package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/google/uuid"
	"lab.dt.multicarta.ru/tp/common/global"
	authService "lab.dt.multicarta.ru/tp/common/messages/auth"
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/comments"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/messages/registries"
	"lab.dt.multicarta.ru/tp/common/messages/response"
	"lab.dt.multicarta.ru/tp/common/messages/twpg"
	webApi "lab.dt.multicarta.ru/tp/web-api-gateway/proto"
	"time"
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

	if p.EmptyEMV {
		response.Result = processing.PassStatus_FAILURE_INCORRECT_CARD
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
		return 4600
	}
	if p.SubCarrier == carriers.SubCarrier_MCK_SUB {
		return 4600
	}
	if p.SubCarrier == carriers.SubCarrier_MMTS_SUB {
		return 4600
	}
	if p.SubCarrier == carriers.SubCarrier_MCD1_MSK {
		return 4600
	}
	if p.SubCarrier == carriers.SubCarrier_MCD2_MSK {
		return 4600
	}
	if p.SubCarrier == carriers.SubCarrier_MCD1_MO {
		return 800
	}
	if p.SubCarrier == carriers.SubCarrier_MCD2_MO {
		return 800
	}
	if p.Carrier == carriers.Carrier_MM {
		return 4600
	}
	if p.Carrier == carriers.Carrier_MGT {
		return 4600
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
			if p.AuthType == AuthTypeMCTokenIncorrect {
				responseOR.Status = processing.AuthStatus_FAILURE_ISSUER
			}
		case PaymentTypeAggregate:
			responseOR.Status = processing.AuthStatus_SUCCESS_AGGREGATE
		}
	}

	if p.AuthType == AuthTypeIncorrect && p.PaymentType != PaymentTypeStartAggregate {
		responseOR.Status = processing.AuthStatus_FAILURE_ISSUER
	}

	if p.AuthType == AuthTypeMCTokenIncorrect && p.PaymentType != PaymentTypeStartAggregate {
		responseOR.Status = processing.AuthStatus_FAILURE_ISSUER
	}

	if p.Duration > 0 {
		request.Timeout = p.Duration
	}

	if p.EmptyEMV {
		responseOR.Result = processing.PassStatus_FAILURE_INCORRECT_CARD
		responseOR.Status = 0
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
			Approve: "CUSTOM_STATUS:PREAUTH-APPROVED",
			Cancel:  gofakeit.URL(),
			Decline: gofakeit.URL(),
		},
	}
}

func TWPGCreatePaymentLinkTest(pay *TWPGCreateAndPayOrderStep) *twpg.CreateTWPGPOrderRequest {
	return &twpg.CreateTWPGPOrderRequest{
		Id: pay.CustomerId,
		Urls: &twpg.TWPGUrls{
			Approve: gofakeit.URL(),
			Cancel:  gofakeit.URL(),
			Decline: gofakeit.URL(),
		},
		PaymentPurpose: twpg.CreateTWPGPOrderRequest_FACE_ID,
		Amount:         300,
		Merchant:       twpg.Merchant_FACE_ID,
	}
}

func TWPGOrderStatusRequest(os *TWPGOrderStatus) (*twpg.OrderStatusRequest, *twpg.OrderStatusResponse) {
	req := &twpg.OrderStatusRequest{
		OrderId: os.OrderId,
	}
	resp := &twpg.OrderStatusResponse{
		Status: os.OrderStatus,
	}

	return req, resp
}

func TWPGReverseOrderRequest(os *TWPGReverseOrder) *twpg.ReverseOrderRequest {
	req := &twpg.ReverseOrderRequest{
		OrderId: os.OrderId,
	}
	return req
}

func WebAPIFaceStatusRequest(faceCheck *FaceIdRegistrationStatus) (*twpg.RegistrationStatusRequest, *twpg.RegistrationStatusResponse) {
	req := &twpg.RegistrationStatusRequest{
		Id: faceCheck.Id,
	}

	resp := &twpg.RegistrationStatusResponse{
		Status: twpg.RegistrationStatusResponse_SUCCESS,
		Card: &twpg.CardInfo{
			Pan: "0900",
			Exp: "2101",
		},
	}

	return req, resp
}

func CardGetFullRequest(c *CardGetFull) (*cards.GetFullRequest, *cards.GetFullResponse) {
	req := &cards.GetFullRequest{
		Requests: make([]*cards.FullRequest, 1),
	}

	req.Requests[0] = &cards.FullRequest{
		Kind:     c.Kind,
		FileType: c.FileType,
	}

	resp := &cards.GetFullResponse{}

	return req, resp
}

func ReaderConfigurationRequest(c *ReaderConfiguration) (*webApi.ReaderRequest, *webApi.ReaderResponse) {
	req := &webApi.ReaderRequest{
		ServerStatus: &webApi.ServerStatus{
			MemoryUsage: 1,
			DiskUsage:   2,
			FreeSpace:   3,
			SystemTime:  4,
			ReadOnly:    false,
		},
		ReaderId: "22",
		Version:  "333",
		Lists: &webApi.Lists{
			FaceList: &webApi.FaceList{
				Time:     uint64(c.FaceList.Time.UnixNano()),
				FileType: c.FaceList.FileType,
			},
			StopList: &webApi.StopList{
				Time:     uint64(c.StopList.Time.UnixNano()),
				FileType: c.StopList.FileType,
			},
		},
	}

	resp := &webApi.ReaderResponse{}

	return req, resp
}

func CardCheckStopListRequest(cardCheck *CardStopList) (*cards.GetCardStatusRequest, *cards.GetCardStatusResponse) {
	req := &cards.GetCardStatusRequest{
		Pan: cardCheck.Pan,
	}

	resp := &cards.GetCardStatusResponse{
		Status: cardCheck.ExpectedStatus,
	}

	return req, resp
}

func ForceReauthRequest(fra *ForceReauth) (*authService.CardByPassIDRequest, *authService.AuthResponseEvent) {
	req := &authService.CardByPassIDRequest{
		PassId: fra.PassId,
	}
	resp := &authService.AuthResponseEvent{
		Response: &authService.AuthResponse{
			AuthStatus: processing.AuthStatus_SUCCESS_AUTH,
		},
		PassId: fra.PassId,
	}
	return req, resp
}

func AddCommentRequest(crud *CommentsCRUD) *comments.AddCommentRequest {
	request := &comments.AddCommentRequest{
		Comment: &comments.Comment{
			EntityId:  uuid.New().String(),
			UserId:    uuid.New().String(),
			RepliesTo: "",
			Body:      gofakeit.HackerPhrase(),
		},
	}

	return request
}

func GetProcessRevisePass() (*processing.OnlinePassEvent, *registries.ProcessRevisePassResponse) {
	rprReq := &processing.OnlinePassEvent{
		Request: &processing.OnlinePassRequest{
			Id:      uuid.New().String(),
			Created: uint64(time.Now().UnixNano()),
			Timeout: 0,
			Auth: &processing.Auth{
				Sum:  0,
				Type: processing.Auth_NOT_REQUIRED,
			},
			Tap: &processing.Tap{
				Created:    uint64(time.Now().UnixNano()),
				Resolution: 1,
				Terminal: &processing.Terminal{
					Id:        "terminal_id",
					Station:   "station_id",
					Direction: processing.TerminalDirection_INGRESS,
				},
				Card: &processing.Card{
					System: processing.CardSystem_UNKNOWN_SYSTEM,
					Type:   processing.CardType_UNKNOWN_TYPE,
					Pan:    CardPan(),
					Bin:    11111111,
					Exp:    "9912",
				},
				Sign: "123",
			},
		},
		Id:          "",
		Created:     uint64(time.Now().UnixNano()),
		CardId:      "",
		CarrierCode: carriers.Carrier_MM,
	}

	rprResp := &registries.ProcessRevisePassResponse{
		IsDuplicate: false,
	}

	return rprReq, rprResp
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

func CompleteWithCalculateRequest(cc *CompleteWithCalculate) (*processing.CompleteWithCalculateRequest, *processing.CompleteWithCalculateResponse) {
	request := &processing.CompleteWithCalculateRequest{
		Pan:  cc.Pan,
		Time: uint64(cc.Date.UnixNano()),
	}

	resp := &processing.CompleteWithCalculateResponse{}

	return request, resp
}

func McdRestoreRequest(mr *McdRestore) (*processing.McdRestoreRequest, *processing.McdRestoreResponse) {
	request := &processing.McdRestoreRequest{
		Pan:  mr.Pan,
		Time: uint64(mr.Date.UnixNano()),
	}

	resp := &processing.McdRestoreResponse{}

	return request, resp
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

func TWPGForceCheckRequest() (*response.EmptyMessage, *response.EmptyMessage) {
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

	if p.AuthType == AuthTypeUnsuccessWithReauth && !isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	if p.AuthType == AuthTypeUnsuccessWithoutReauth && !isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	if p.AuthType == AuthTypeIncorrect && p.isComplete && isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	if p.AuthType == AuthTypeMCTokenIncorrect && !isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	if p.AuthType == AuthTypeMCTokenIncorrect && p.isComplete && isAggregate(p) {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	if p.EmptyEMV {
		response.Resolution = processing.AuthResponse_NONE_RESOLUTION
		response.Status = processing.AuthResponse_NONE_STATUS
		response.Result = processing.AuthResponse_FAILURE_AUTH_NOT_FOUND
	}

	//aggregate by card system
	if isAggeregateByCardSystemCarrier(p.Carrier, p.card.System) {
		response.Status = processing.AuthResponse_SUCCESS_AGGREGATE

		response.Auth = &processing.Auth{
			Sum:  0,
			Type: processing.Auth_AGGREGATE,
		}
		response.Resolution = processing.AuthResponse_NONE_RESOLUTION
	}

	return request, response
}

//isAggeregateByCardSystemCarrier ?????????????? ?????????????????????????? ???????????? ?? ???????????? ?????????????????????? ?? ??????????
func isAggeregateByCardSystemCarrier(carrier carriers.Carrier, ps processing.CardSystem) bool {
	if carrier != carriers.Carrier_MM && carrier != carriers.Carrier_MGT {
		return false
	}

	return global.IsAggregate(ps, carrier)
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
				Id:         gofakeit.Country(),
				Station:    gofakeit.Color(),
				Direction:  processing.TerminalDirection_INGRESS,
				SubCarrier: c,
			},
			Card: card,
			Sign: "1.0.2",
		},
	}
	if p.Terminal != nil {
		if p.Terminal.Id == "" {
			p.Terminal.Id = gofakeit.HipsterWord()
		}
		request.Tap.Terminal = p.Terminal
	}
	if p.tapRequest != nil {
		request.Tap.Created = p.tapRequest.Tap.Created
	}
	response := &processing.TapResponse{
		Id:      "",
		Created: 0,
		Result:  processing.TapResponse_SUCCESS,
		Msg:     "",
	}

	if p.EmptyEMV {
		request.Tap.Card.Emv = ""
	}
	p.Terminal = request.Tap.Terminal
	return request, response
}
