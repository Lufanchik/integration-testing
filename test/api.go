package test

import (
	"context"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/global"
	authService "lab.dt.multicarta.ru/tp/common/messages/auth"
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/common/messages/comments"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/messages/registries"
	protoResponse "lab.dt.multicarta.ru/tp/common/messages/response"
	"lab.dt.multicarta.ru/tp/common/messages/twpg"
	"lab.dt.multicarta.ru/tp/common/messages/user"
	webApi "lab.dt.multicarta.ru/tp/web-api-gateway/proto"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TapBySubCarrier(t *testing.T, p *Pass, card *processing.Card) *processing.TapRequest {
	req, resp := TapRequest(p.SubCarrier, card, p)
	u := "/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessTap"

	r := httpProcessingApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &processing.TapResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)

	resp.Id = response.Id
	resp.Created = response.Created
	p.id = response.Id

	require.Equal(t, resp, response)

	return req
}

func PassBySubCarrier(t *testing.T, tap *processing.TapRequest, p *Pass) uint64 {
	var requestedTime uint64
	switch p.RequestType {
	case RequestTypeOffline:
		req, resp := PassOfflineRequest(tap, p)
		requestedTime = req.Created
		u := "/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
		r := httpProcessingApi.POST(u).WithJSON(req).
			Expect().
			Status(http.StatusOK)
		object := r.Body().Raw()
		logRequest(u, r)

		response := &processing.OfflinePassResponse{}
		err := jsonpb.Unmarshal(strings.NewReader(object), response)
		require.NoError(t, err)

		resp.Id = response.Id
		resp.Created = response.Created
		p.id = response.Id

		require.Equal(t, resp, response)
	case RequestTypeOnline:
		req, resp := PassOnlineRequest(tap, p)
		requestedTime = req.Created
		u := "/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
		r := httpProcessingApi.POST(u).WithJSON(req).
			Expect().
			Status(http.StatusOK)

		object := r.Body().Raw()
		logRequest(u, r)

		response := &processing.OnlinePassResponse{}
		err := jsonpb.Unmarshal(strings.NewReader(object), response)
		require.NoError(t, err)

		resp.Id = response.Id
		resp.Created = response.Created
		p.id = response.Id

		require.Equal(t, resp, response)
	}
	return requestedTime
}

func Update(t *testing.T, p *Pass, up Updater) {
	ctx := context.Background()
	tap, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: p.id,
	})
	require.NoError(t, err)
	up.f(tap)
	_, err = ps.UpdatePass(ctx, tap)
	require.NoError(t, err)
}

func ValidatePass(t *testing.T, p *Pass, parent *Pass, ingress *Pass, isFirst bool) {
	ctx := context.Background()

	expectPass := &pass.Pass{
		Id:                p.id,
		UserId:            p.card.Pan,
		Kind:              0,
		CarrierTapId:      p.tapRequest.Id,
		CarrierCode:       p.Carrier,
		CarrierResolution: p.tapRequest.Tap.Resolution,
		TerminalId:        p.tapRequest.Tap.Terminal.Id,
		TerminalStation:   p.tapRequest.Tap.Terminal.Station,
		TerminalDirection: p.tapRequest.Tap.Terminal.Direction,
		Sign:              p.tapRequest.Tap.Sign,
		IsCancel:          false,
		IsComplexTimeout:  false,
		CreatedAtRequest:  NanoToMicro(p.timeRequest),
		CreatedAtCarrier:  NanoToMicro(p.tapRequest.Tap.Created),
		IsComplex:         false,
		ParentComplexId:   "",
		IsComplexCarrier:  global.IsComplexCarrier(p.Carrier),
		IsPass:            true,
		IsFree:            false,
		CarrierCodeSub:    p.SubCarrier,
		Sum:               getSumByCarrier(p),
		IsAuth:            false,
		PassType:          GetPassType(p),
		PaySystem:         p.card.System,
	}

	if p.timeToWait != 0 {
		expectPass.CreatedAtCarrier = NanoToMicro(p.tapRequest.Tap.Created + uint64(p.timeToWait.Nanoseconds()))
	}

	if ingress != nil {
		expectPass.IngressId = ingress.id
	}

	if p.isParent {
		expectPass.IsComplex = true
	}

	switch p.RequestType {
	case RequestTypeOffline:
		expectPass.Kind = processing.TapType_PASS_OFFLINE
	case RequestTypeOnline:
		expectPass.Kind = processing.TapType_PASS_ONLINE
	}

	switch p.PaymentType {
	case PaymentTypeFree:
		expectPass.IsFree = true
		expectPass.IsAuth = false
	case PaymentTypePayment:
		expectPass.IsFree = false
		expectPass.IsAuth = true
	case PaymentTypePrepayed:
		expectPass.IsFree = false
		expectPass.IsAuth = false
	}

	if isAggregate(p) {
		expectPass.IsAuth = false
		expectPass.IsAggregate = true
	}

	if p.Parent > 0 {
		expectPass.IsComplex = true
		expectPass.ParentComplexId = parent.id
	}

	if p.AuthType == AuthTypeIncorrect {
		expectPass.IsAuth = false
	}

	if p.AuthType == AuthTypeMCTokenIncorrect {
		expectPass.IsAuth = false
	}

	if p.AuthType == AuthTypeUnsuccessWithReauth {
		expectPass.IsAuth = false
	}

	if p.AuthType == AuthTypeUnsuccessWithoutReauth {
		expectPass.IsAuth = false
	}

	if p.isCancel {
		expectPass.IsCancel = true

		if p.AuthType == AuthTypeRefund {
			expectPass.CancelType = processing.CancelType_CANCEL_REFUND
		}
	}

	if p.PaymentType == PaymentTypeStartAggregate && p.AuthType == AuthTypeCorrect && !isFirst {
		expectPass.IsAuth = true
	}

	if p.PaymentType == PaymentTypeStartAggregate && p.AuthType == AuthTypeMCTokenCorrect && !isFirst {
		expectPass.IsAuth = true
	}

	if p.aggregate != nil && p.PaymentType == PaymentTypeAggregate && p.aggregate.AuthType == AuthTypeCorrect && !isFirst {
		expectPass.IsAuth = true
	}

	if p.aggregate != nil && p.PaymentType == PaymentTypeAggregate && p.aggregate.AuthType == AuthTypeMCTokenCorrect && !isFirst {
		expectPass.IsAuth = true
	}

	if p.aggregate != nil && p.PaymentType == PaymentTypeAggregate && p.aggregate.AuthType == AuthTypeCorrect && !isFirst {
		expectPass.AggregateId = p.aggregate.id
	}

	if p.aggregate != nil && p.PaymentType == PaymentTypeAggregate && p.aggregate.AuthType == AuthTypeMCTokenCorrect && !isFirst {
		expectPass.AggregateId = p.aggregate.id
	}

	if p.PaymentType == PaymentTypeStartAggregate && p.AuthType == AuthTypeCorrect && !isFirst {
		expectPass.AggregateId = expectPass.Id
	}

	if p.PaymentType == PaymentTypeStartAggregate && p.AuthType == AuthTypeMCTokenCorrect && !isFirst {
		expectPass.AggregateId = expectPass.Id
	}

	if p.PaymentType == PaymentTypeStartAggregate && isFirst {
		expectPass.Sum = 0
	}

	if p.PaymentType == PaymentTypePrepayed {
		expectPass.Sum = 0
	}

	if p.PaymentType == PaymentTypeStartAggregate && !isFirst {
		expectPass.Sum = getSumByCarrier(p)
		expectPass.SumAggregate = getSumByCarrier(p)
	}

	if p.PaymentType == PaymentTypeStartAggregate && !isFirst && p.AuthType == AuthTypeIncorrect {
		expectPass.Sum = 0
		expectPass.SumAggregate = getSumByCarrier(p)
	}

	if p.PaymentType == PaymentTypeStartAggregate && !isFirst && p.AuthType == AuthTypeMCTokenIncorrect {
		expectPass.Sum = 0
		expectPass.SumAggregate = getSumByCarrier(p)
	}

	if p.PaymentType == PaymentTypeStartAggregate {
		expectPass.IsInitAggregate = true
	}

	expectPass.IsComplexTimeout = global.IsComplexTimeout(global.UnixNanoToLocalizedTime(expectPass.CreatedAtCarrier))

	passDB, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: p.id,
	})

	isEqual := assert.ObjectsAreEqual(expectPass, passDB)
	counter := 0

	for !isEqual {
		time.Sleep(TimeAfterRequest)
		passDB, err = ps.GetPass(ctx, &pass.PassRequest{
			Id: p.id,
		})

		if passDB != nil {
			expectPass.Updated = passDB.Updated
		}

		isEqual = assert.ObjectsAreEqual(expectPass, passDB)
		counter++
		if counter > 200 {
			break
		}
	}

	require.Equal(t, expectPass, passDB)
	require.NoError(t, err)
}

func GetAuthStatus(p *Pass) (*processing.AuthResponse, *processing.AuthResponse, error) {
	req, resp := AuthStatusRequest(p)
	u := "/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/AuthStatus"
	r := httpProcessingApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &processing.AuthResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)

	resp.Created = response.Created
	resp.Info = response.Info

	return resp, response, err
}

func AuthStatus(t *testing.T, p *Pass) {
	if p.faceId != "" {
		return
	}
	// Чтобы всегда при проходах по МТ игнорировать проверку AuthStatus
	if p.PassType == pass.PassType_PASS_MT {
		return
	}

	resp, response, err := GetAuthStatus(p)

	isEqual := assert.ObjectsAreEqual(resp, response)
	counter := 0
	for !isEqual {
		time.Sleep(TimeAfterRequest)
		resp, response, err = GetAuthStatus(p)
		isEqual = assert.ObjectsAreEqual(resp, response)
		counter++
		if counter > 200 {
			break
		}
	}

	require.Equal(t, resp, response)
	require.NoError(t, err)
}

func ResolveTestApi(t *testing.T, _case *Resolve) {
	var r *httpexpect.Response

	switch _case.Url {
	// GetSchema
	case "/twirp/proto.ResolveHttp/GetServiceSchema":
		r = httpResolveService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().
			ContainsKey("tables").
			ContainsKey("relations")
	case "/twirp/proto.ResolveHttp/GetExpiresLink":
		r = httpResolveService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().
			ContainsKey("link_expires")
	case "/twirp/proto.ResolveHttp/GetTasks":
		r = httpResolveService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
	default:
		t.Error("Error to find test handler for url:", _case.Url)
		return
	}

	logRequest(_case.Url, r)
}

func ReviseTestApi(t *testing.T, _case *Revise) {
	var r *httpexpect.Response

	switch _case.Url {
	// GetSchema
	case "/twirp/proto.ReviseHttp/GetServiceSchema":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().
			ContainsKey("tables").
			ContainsKey("relations")
	case "/twirp/proto.ReviseHttp/GetTasks":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object()
	case "/twirp/proto.ReviseHttp/GetExpiresLink":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().
			ContainsKey("link_expires")
	case "/twirp/proto.ReviseHttp/ResetTask":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().NotContainsKey("meta")
	case "/twirp/proto.ReviseHttp/GetStages":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)

		r.JSON().Object().ContainsKey("stages")
	case "/twirp/proto.ReviseHttp/GetTaskStatus":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().NotContainsKey("meta")
	case "/twirp/proto.ReviseHttp/GetLastByOrderFileInfo":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().ContainsKey("file")
	case "/twirp/proto.ReviseHttp/Playground":
		r = httpReviseService.POST(_case.Url).
			WithJSON(_case.Request).
			Expect().
			Status(_case.Status)
		r.JSON().Object().ContainsKey("query").ContainsKey("values")
	default:
		t.Error("Error to find test handler for url:", _case.Url)
		return
	}

	logRequest(_case.Url, r)
}

func AbsGetRegistryApi(t *testing.T, registry *AbsGetRegistry) {
	req := registries.AbsRegistryRequest{}
	u := "/twirp/proto.ApmAPIGateway/AbsGetRegistry"
	r := httpApmApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusForbidden)

	logRequest(u, r)
}

func LoginApi(t *testing.T, lg *Login) {
	req := user.LoginRequest{}
	u := "/twirp/proto.ApmAPIGatewayPublic/Login"
	r := httpApmApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusNotFound)

	logRequest(u, r)
}

func PassCheckApi(t *testing.T, pc *PassCheck, target *Pass, parent *Pass) {
	target.PaymentType = pc.PaymentType
	target.ExpectedSum = pc.ExpectedSum
	ValidatePass(t, target, parent, nil, true)
	AuthStatus(t, target)
}

func CancelApi(t *testing.T, cl *Cancel, target *Pass) {
	req, resp := CancelRequest(cl, target)
	u := "/" + target.Carrier.String() + "/twirp/sirocco.ProcessingAPI/CancelPass"
	r := httpProcessingApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &processing.CancelPassResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)
	resp.Created = response.Created
	require.Equal(t, resp, response)
}

func ParkingApi(t *testing.T, card *processing.Card, pr *Parking) {
	req, resp := ParkingRequest(card, pr)
	u := "/mm/twirp/sirocco.ProcessingAPI/CheckParking"
	r := httpProcessingApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &processing.CheckParkingResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)
	require.Equal(t, resp, response)
}

func CompleteApi(t *testing.T, pass *Pass, passes []*Pass, sum int) {
	req, resp := CompleteRequest(pass, passes, sum)
	u := "/" + pass.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessComplete"
	r := httpProcessingApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &processing.CompleteResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	resp.Created = response.Created
	require.NoError(t, err)
	require.Equal(t, resp, response)
}

func WebAPI(t *testing.T, card *processing.Card, passes []*Pass) {
	log.Infof(card.Pan)

	counter := 0
	var response *webApi.PassesResponse
	for {
		req := WebAPIPassesRequest(card)
		u := "/twirp/proto.WebAPIGateway/GetPasses"
		r := httpWebApi.POST(u).WithJSON(req).
			Expect().
			Status(http.StatusOK)

		object := r.Body().Raw()
		logRequest(u, r)

		response = &webApi.PassesResponse{}
		err := jsonpb.Unmarshal(strings.NewReader(object), response)
		require.NoError(t, err)

		if len(response.Passes) == 0 {
			time.Sleep(TimeAfterRequest)
			counter++
			if counter > 200 {
				break
			} else {
				continue
			}
		}

		break
	}

	require.Equal(t, len(passes), len(response.Passes), "passes count doesn't match")

	responsePassesMap := map[string]struct{}{}
	for _, p := range response.Passes {
		if p.Status == webApi.PassStatus_PAID {
			require.NotEmpty(t, p.AuthTime)
		} else {
			require.Empty(t, p.AuthTime)
		}

		responsePassesMap[p.Id] = struct{}{}
	}

	for _, p := range passes {
		_, ok := responsePassesMap[p.id]
		require.True(t, ok, "pass not found: %s", p.id)
	}
}

func CommentsCheck(t *testing.T, crud *CommentsCRUD) {
	// Add comment
	addReq := AddCommentRequest(crud)
	addCommentURL := "/twirp/proto.CommentsService/AddComment"
	addR := httpCommentService.POST(addCommentURL).WithJSON(addReq).
		Expect().
		Status(http.StatusOK)
	logRequest(addCommentURL, addR)

	fmt.Printf("EntityID: %s\n", addReq.Comment.EntityId)

	// Get comments
	getReq := GetCommentsRequest(addReq)
	getCommenstURL := "/twirp/proto.CommentsService/GetComments"
	getR := httpCommentService.POST(getCommenstURL).WithJSON(getReq).
		Expect().
		Status(http.StatusOK)
	object := getR.Body().Raw()
	logRequest(getCommenstURL, getR)

	getCommentsResponse := &comments.GetCommentsResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), getCommentsResponse)
	require.NoError(t, err)

	count := len(getCommentsResponse.Comments)
	require.Equal(t, 1, count)
	comment := getCommentsResponse.Comments[0]

	require.Equal(t, addReq.Comment.Body, comment.Body)
	require.Equal(t, addReq.Comment.EntityId, comment.EntityId)
	require.Equal(t, addReq.Comment.UserId, comment.UserId)

	// Delete comments
	delReq := DelCommentRequest(comment)
	delCommentURL := "/twirp/proto.CommentsService/DeleteComment"
	delR := httpCommentService.POST(delCommentURL).WithJSON(delReq).
		Expect().
		Status(http.StatusOK)
	logRequest(delCommentURL, delR)

	// Check deletion
	getR = httpCommentService.POST(getCommenstURL).WithJSON(getReq).
		Expect().
		Status(http.StatusOK)
	object = getR.Body().Raw()
	logRequest(getCommenstURL, getR)

	getCommentsResponse = &comments.GetCommentsResponse{}
	err = jsonpb.Unmarshal(strings.NewReader(object), getCommentsResponse)
	require.NoError(t, err)

	count = len(getCommentsResponse.Comments)
	require.Equal(t, 1, count)
	comment = getCommentsResponse.Comments[0]

	require.Equal(t, "", comment.Body)
	require.Equal(t, true, comment.IsDeleted)
}

func FaceApiCheckStatus(t *testing.T, faceCheck *FaceIdRegistrationStatus) {
	req, expectedResponse := WebAPIFaceStatusRequest(faceCheck)
	u := "/twirp/proto.WebAPIGateway/FaceIdRegistrationStatus"

	r := httpWebApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	actualResponse := &twpg.RegistrationStatusResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), actualResponse)
	require.NoError(t, err)

	require.Equal(t, expectedResponse, actualResponse)
}

func CardApiGetFull(t *testing.T, c *CardGetFull) {
	req, _ := CardGetFullRequest(c)
	u := "/twirp/proto.CardService/GetFull"

	r := httpCardService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	actualResponse := &cards.GetFullResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), actualResponse)
	require.NoError(t, err)

	require.Equal(t, 1, len(actualResponse.Fulls))
}

func CardCheckStopList(t *testing.T, cardCheck *CardStopList) {
	req, expectedResponse := CardCheckStopListRequest(cardCheck)
	u := "/twirp/proto.CardService/GetCardStatus"

	r := httpCardService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	actualResponse := &cards.GetCardStatusResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), actualResponse)
	require.NoError(t, err)

	require.Equal(t, expectedResponse, actualResponse)
}

func ForceReauthCall(t *testing.T, fra *ForceReauth) {
	req, expectedResponse := ForceReauthRequest(fra)
	u := "/twirp/sirocco.AuthAPI/ActiveReAuth"

	r := httpAuthService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	fmt.Printf("pass_id: %s", fra.PassId)
	object := r.Body().Raw()
	fmt.Println(object)
	logRequest(u, r)

	actualResponse := &authService.AuthResponseEvent{}
	err := jsonpb.Unmarshal(strings.NewReader(object), actualResponse)
	require.NoError(t, err)

	require.Equal(t, expectedResponse.Response.AuthStatus, actualResponse.Response.AuthStatus)
}

func FaceApiGetRegisterLink(t *testing.T, card *processing.Card, fcl *RegisterFaceId) {
	req := WebAPIRegisterRequest(fcl, card)
	u := "/twirp/proto.WebAPIGateway/RegisterFaceId"
	r := httpWebApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &authService.CreateTWPGOrderResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)

	require.NotEmpty(t, response.Redirect)
	log.Infof("Redirect URL: %s", response.Redirect)

	fcl.RedirectURL = response.Redirect

	err = twpgForceCheck()
	require.NoError(t, err)
	time.Sleep(time.Second)
}

func TWPGCreateAndPayOrder(t *testing.T, pay *TWPGCreateAndPayOrderStep) {
	req := TWPGCreatePaymentLinkTest(pay)
	u := "/twirp/proto.TWPGService/CreatePaymentLink"
	r := httpTWPGService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &twpg.CreateTWPGOrderResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)

	require.NotEmpty(t, response.Redirect)
	log.Infof("Redirect URL: %s", response.Redirect)

	pay.CustomerId = response.Redirect
	pay.OrderId = response.Id

	err = twpgForceCheck()
	require.NoError(t, err)
	time.Sleep(time.Second)
}

func TWPGCheckOrderStatus(t *testing.T, os *TWPGOrderStatus) {
	req, expectedResponse := TWPGOrderStatusRequest(os)
	u := "/twirp/proto.TWPGService/OrderStatus"
	r := httpTWPGService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &twpg.OrderStatusResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)
	require.Equal(t, expectedResponse, response)
}

func TWPGReverse(t *testing.T, os *TWPGReverseOrder) {
	req := TWPGReverseOrderRequest(os)
	u := "/twirp/proto.TWPGService/ReverseOrder"
	r := httpTWPGService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	logRequest(u, r)
	err := twpgForceCheck()
	require.NoError(t, err)
	time.Sleep(time.Second)
}

func twpgForceCheck() error {
	req, _ := TWPGForceCheckRequest()
	u := "/twirp/proto.TWPGService/FaceForceCheck"
	r := httpTWPGService.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &protoResponse.EmptyMessage{}
	return jsonpb.Unmarshal(strings.NewReader(object), response)
}
