package test

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/global"
	authService "lab.dt.multicarta.ru/tp/common/messages/auth"
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/comments"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/messages/registries"
	protoResponse "lab.dt.multicarta.ru/tp/common/messages/response"
	"lab.dt.multicarta.ru/tp/common/messages/twpg"
	"lab.dt.multicarta.ru/tp/common/messages/user"
	"lab.dt.multicarta.ru/tp/common/rabbit/routes"
	"lab.dt.multicarta.ru/tp/common/uuid"
	webApi "lab.dt.multicarta.ru/tp/web-api-gateway/proto"
	"net/http"
	"strings"
	"testing"
	"time"
)

var timeCount = 400

func TapBySubCarrier(t *testing.T, p *Pass, card *processing.Card) (*processing.TapRequest, *processing.TapResponse) {
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

	return req, response
}

type PassResponser interface {
	GetId() string
}

func PassBySubCarrier(t *testing.T, tap *processing.TapRequest, p *Pass) (uint64, PassResponser) {
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
		return requestedTime, response
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

		if p.EmptyEMV {
			resp.Msg = response.Msg
		}

		require.Equal(t, resp, response)
		return requestedTime, response
	}
	panic(errors.New("not found type pass"))
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
	fmt.Printf("pass: %+v\n", *p)
	if ingress != nil {
		fmt.Printf("ingress: %+v\n", *ingress)
	}
	if p.secondParent != nil {
		fmt.Printf("secondParent: %+v\n", *p.secondParent)
	}
	ctx := context.Background()
	//time.Sleep(time.Second * 10)
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
		IsNotInitComplex:  false,
	}

	if p.timeToWait != 0 {
		expectPass.CreatedAtCarrier = NanoToMicro(p.tapRequest.Tap.Created + uint64(p.timeToWait.Nanoseconds()))
	}

	if ingress != nil {
		expectPass.IngressId = ingress.id
		expectPass.IngressCarrierTapId = ingress.carrierID
	}

	if p.isParent {
		expectPass.IsComplex = true
		expectPass.ParentComplexId = p.id
		expectPass.IsNotInitComplex = false
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

	//aggregate by card system(VISA, MASTERCARD)
	if isAggeregateByCardSystemCarrier(p.Carrier, p.card.System) {
		//проход является агрегационным
		expectPass.IsAggregate = true

		//aggregation_id указывает на открывайющий агрегацию проход
		if p.aggregate != nil && !isFirst {
			expectPass.AggregateId = p.aggregate.id
		}

		//до комплита
		if isFirst {
			expectPass.IsAuth = false
			expectPass.Sum = 0
			expectPass.IsAuthSend = false
		} else {
			//после комплита
			expectPass.Sum = getSumByCarrier(p)
			expectPass.IsAuthSend = true
		}

		//является ли проход инициирующим агрегацию
		if p.IsInitAggregate {
			//чеккард идет по инициирующему агрегацию проходу
			expectPass.IsSuccessCheckCard = true
			if !isFirst {
				//после комплита мы можем сравнить сумму агрегации ,посланную на авторизацию
				expectPass.SumAggregate = p.ExpectedSumAggregate
				expectPass.IsInitAggregate = true
			}
		} else {
			expectPass.IsSuccessCheckCard = false
		}
	}

	if isAggregate(p) {
		expectPass.IsAuth = false
		expectPass.IsAggregate = true
	}

	if p.Parent > 0 {
		expectPass.IsComplex = true
		expectPass.ParentComplexId = parent.id
		expectPass.IsNotInitComplex = true
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

	if !isFirst && p.secondParent != nil {
		expectPass.IsComplex = true
		expectPass.ParentComplexId = p.secondParent.id
		expectPass.IsNotInitComplex = true
	}

	expectPass.IsComplexTimeout = global.IsComplexTimeout(global.UnixNanoToLocalizedTime(expectPass.CreatedAtCarrier))

	passDB, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: p.id,
	})

	isEqual := assert.ObjectsAreEqual(expectPass, passDB)
	counter := 0

	//if p.EmptyEMV {
	//	expectPass.Sum = 0
	//}

	for !isEqual {
		fmt.Println("pass not equal - " + t.Name())
		time.Sleep(TimeAfterRequest)
		passDB, err = ps.GetPass(ctx, &pass.PassRequest{
			Id: p.id,
		})

		if passDB != nil {
			expectPass.Updated = passDB.Updated
		}

		if passDB != nil && passDB.TerminalDirection == processing.TerminalDirection_EGRESS && passDB.IsComplex {

		}

		isEqual = assert.ObjectsAreEqual(expectPass, passDB)
		counter++
		if counter > timeCount {
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

	if p.EmptyEMV {
		resp.Auth = response.Auth
	}

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
		fmt.Println("auth status not equal")
		time.Sleep(TimeAfterRequest)
		resp, response, err = GetAuthStatus(p)
		isEqual = assert.ObjectsAreEqual(resp, response)
		counter++
		if counter > timeCount {
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

func CompleteCalcApi(t *testing.T, aggregatePasses []*Pass, pan string) {
	//check passes pan
	for _, value := range aggregatePasses {
		require.NotNil(t, value.card)
		require.Equal(t, pan, value.card.Pan)
	}

	//made complete
	req, resp := CompleteWithCalculateRequest(pan)
	url := "/twirp/proto.WebAPIGateway/ProcessCompleteWithCalculate"
	r := httpWebApi.POST(url).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(url, r)
	response := &processing.CompleteWithCalculateResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)
	require.Equal(t, resp, response)

	//check passes after complete
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
			if counter > timeCount {
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

func ProcessRevisePassRequest(t *testing.T, prp *ProcessRevisePass) {
	prpReq, prpResp := GetProcessRevisePass()
	prp.req = prpReq
	prpR, err := ps.ProcessRevisePass(context.Background(), prpReq)
	require.NoError(t, err)
	require.Equal(t, prpResp, prpR)
	ctx := context.Background()

	id, err := uuid.GenerateUUID(ctx, prpReq.Request)
	passDB, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: id,
	})
	require.NoError(t, err)
	prp.pass = passDB

	expectPass := &pass.Pass{
		Id:                id,
		UserId:            prpReq.Request.Tap.Card.Pan,
		Kind:              processing.TapType_PASS_ONLINE,
		CarrierTapId:      prpReq.Request.Id,
		CarrierCode:       prpReq.CarrierCode,
		CarrierResolution: prpReq.Request.Tap.Resolution,
		TerminalId:        prpReq.Request.Tap.Terminal.Id,
		TerminalStation:   prpReq.Request.Tap.Terminal.Station,
		TerminalDirection: prpReq.Request.Tap.Terminal.Direction,
		Sign:              prpReq.Request.Tap.Sign,
		IsCancel:          false,
		IsComplexTimeout:  false,
		CreatedAtRequest:  passDB.CreatedAtRequest,
		Updated:           passDB.Updated,
		CreatedAtCarrier:  NanoToMicro(prpReq.Request.Tap.Created),
		IsComplex:         false,
		ParentComplexId:   "",
		IsComplexCarrier:  global.IsComplexCarrier(prpReq.CarrierCode),
		IsPass:            true,
		IsFree:            false,
		Sum:               4600,
		IsAuth:            false,
		PassType:          pass.PassType_PASS_BBK,
		PaySystem:         prpReq.Request.Tap.Card.System,
	}

	isEqual := assert.ObjectsAreEqual(expectPass, passDB)
	counter := 0

	for !isEqual {
		fmt.Println("pass not equal")
		time.Sleep(TimeAfterRequest)
		passDB, err = ps.GetPass(ctx, &pass.PassRequest{
			Id: id,
		})

		if passDB != nil {
			expectPass.Updated = passDB.Updated
		}

		isEqual = assert.ObjectsAreEqual(expectPass, passDB)
		counter++
		if counter > timeCount {
			break
		}
	}

	require.Equal(t, expectPass, passDB)
	require.NoError(t, err)
}

func CommentsCheck(t *testing.T, crud *CommentsCRUD) {
	atClaims := GetAtClaims()
	buf, err := json.Marshal(atClaims.Info)

	// Add comment
	addReq := AddCommentRequest(crud)
	addCommentURL := "/twirp/proto.CommentsService/AddComment"
	addR := httpCommentService.POST(addCommentURL).
		WithHeader("user", string(buf)).
		WithJSON(addReq).
		Expect().
		Status(http.StatusOK)
	logRequest(addCommentURL, addR)

	addCommentResponse := &comments.AddCommentResponse{}
	object := addR.Body().Raw()
	err = jsonpb.Unmarshal(strings.NewReader(object), addCommentResponse)
	require.NoError(t, err)
	addReq.Comment.Id = addCommentResponse.CommentId

	fmt.Printf("EntityID: %s\n", addReq.Comment.EntityId)

	// Get comments
	getReq := GetCommentsRequest(addReq)
	getCommenstURL := "/twirp/proto.CommentsService/GetComments"
	getR := httpCommentService.POST(getCommenstURL).WithJSON(getReq).
		Expect().
		Status(http.StatusOK)
	object = getR.Body().Raw()
	logRequest(getCommenstURL, getR)

	getCommentsResponse := &comments.GetCommentsResponse{}
	err = jsonpb.Unmarshal(strings.NewReader(object), getCommentsResponse)
	require.NoError(t, err)

	count := len(getCommentsResponse.Comments)
	require.Equal(t, 1, count)
	comment := getCommentsResponse.Comments[0]

	require.Equal(t, addReq.Comment.Body, comment.Body)
	require.Equal(t, addReq.Comment.EntityId, comment.EntityId)
	require.Equal(t, atClaims.Info.Uuid, comment.UserId)

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

func ReaderConfigurationSend(t *testing.T, c *ReaderConfiguration) {
	req, _ := ReaderConfigurationRequest(c)
	u := "/twirp/proto.WebAPIGateway/ReaderConfiguration"

	r := httpWebApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	actualResponse := &webApi.ReaderResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), actualResponse)
	require.NoError(t, err)
	require.NotNil(t, actualResponse.Lists)
	require.Equal(t, true, len(actualResponse.Lists.FaceList.Url) > 0)

	r = httpWebApi.GET("/download/s3").WithQuery("id", actualResponse.Lists.FaceList.Id).
		Expect().
		Status(http.StatusOK)

	object = r.Body().Raw()
	logRequest(u, r)

	h := md5.New()
	h.Write([]byte(r.Body().Raw()))
	require.NoError(t, err)

	hash := fmt.Sprintf("%x", h.Sum(nil))
	require.Equal(t, actualResponse.Lists.FaceList.Md5, hash)
	fmt.Println(hash)

	r = httpWebApi.GET("/download/s3").WithQuery("id", actualResponse.Lists.StopList.Id).
		Expect().
		Status(http.StatusOK)

	object = r.Body().Raw()
	logRequest(u, r)

	s := md5.New()
	s.Write([]byte(r.Body().Raw()))
	require.NoError(t, err)

	hashS := fmt.Sprintf("%x", s.Sum(nil))
	fmt.Println(hashS)
	require.Equal(t, actualResponse.Lists.StopList.Md5, hashS)
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

func AuthResponseGo(t *testing.T, ra *AuthResponse) {
	rta, err := routes.QueueByID(routes.RouteAuthResponse)
	require.NoError(t, err)
	rtat := rta.RouteByParam("MM")
	aResp := &authService.AuthResponseEvent{
		Response: &authService.AuthResponse{
			Sum:              4600,
			Rrn:              "000022371574",
			AuthId:           "d5f2e434-525b-401e-88a2-6c754b7255e2",
			ParentAuthId:     "",
			AuthStatus:       processing.AuthStatus_SUCCESS_AUTH,
			TransactionType:  authService.TransactionType_TRANSACTION_AUTH,
			CreatedAtRequest: uint64(time.Date(2020, time.August, 28, 9, 00, 00, 00, time.UTC).UnixNano()),
			Invoice:          "6039136839172976",
			Type:             1,
			FacePurchase:     nil,
			FacePurchase2:    nil,
			IsToken:          false,
		},
		CarrierCode:       carriers.Carrier_MM,
		PassId:            "3b422d69-b7d8-50be-baf7-6579ee0299c6",
		CreatedAtCarrier:  uint64(time.Now().UnixNano()),
		AggregatePassesId: nil,
	}

	var serverSpan opentracing.Span
	options := []opentracing.StartSpanOption{}
	serverSpan = opentracing.StartSpan("HTTP request", options...)
	defer serverSpan.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)

	err = rpa.Publish(ctx, rtat, aResp)
	require.NoError(t, err)
}

func ReAuthGo(t *testing.T, ra *ReAuth) {
	passes := []string{
		"f92d7d9a-8d31-5bc8-be93-1afe8c5280c1",
		"f9287a0b-6bd4-551d-89c5-5c7d68ed1d41",
		"f6751d00-00fe-5359-a31d-17a843ddc1d3",
		"f3697a7f-748e-5d2b-a1e7-b292722ae4c5",
		"e6360ae9-2f46-5577-8efe-a816a4c30379",
		"dc9ee339-39a1-5882-aac8-d0e12af48182",
		"d8adfd46-bfff-5560-b015-954d68e5aa5a",
		"d441eac8-0641-506a-90ad-011b828be201",
		"d4ce596a-ea46-5738-b775-cf71954e67f2",
		"d0a6319e-aba3-500f-9110-aa05135a3454",
		"cce75812-9b42-5036-9d6c-d99467fe41e6",
		"c60f16e1-8dbf-551d-8e11-17a34ecec823",
		"bc679fa2-ff8f-5503-bde2-f4a2cc199431",
		"bc2a3876-3af6-5ff4-a0a5-099a26e0e0ec",
		"b58b667d-a93d-5c0c-90d8-fe1e1613a126",
		"b2be5509-38f4-51dc-b2f3-ed0b89f3cdbb",
		"a9c3c81c-76bb-5d5b-aae1-5d444e749e22",
		"bdf42cf6-a0a8-5e10-87ef-8047b1003ca3",
		"92237ec3-9e67-5caa-a5d3-7f6f6c7b1d4f",
		"8ece21aa-f5ff-5794-a596-03e2a6394741",
		"88368740-971d-5318-a02d-311539d093f1",
		"7fd7bbb6-7763-54a3-b6c4-6d8eb7d9ac26",
		"7acb6c8a-b55e-5882-99cc-85ce4dd9881e",
		"76673c19-04ce-5239-b5f4-51e8ba03dc8a",
		"753ebd51-87e7-505b-aefd-9acea178cac7",
		"70dd5425-b2e6-5f28-8632-3a2ff3c70f9b",
		"6de626da-8987-5a82-85d8-333d398cfc2a",
		"6dd50a07-6846-5734-aa70-ea7dcfcee5b8",
		"5fbf8bc5-be59-5c11-a7e3-1027b74b2e1f",
		"564fb7b1-0d9e-5e50-a751-8e45fc9b3ae2",
		"54024db8-88db-5a32-b5da-f1e5faa9dae8",
		"4eecab9f-a5ea-5778-9934-71eba24f5f73",
		"4cc56b64-0a93-5ba1-bab4-0f0c9b8b4a49",
		"4bb5efd7-42fb-546f-b83c-207650471976",
		"475d012d-104f-5f18-a68d-32e77c948896",
		"36cc26db-3178-5778-95af-bebd49c8da6f",
		"31e5e9cc-d9a6-5d05-bfd6-8558b3123e6a",
		"2c0b96da-1e79-5193-801a-15c686d1a7e2",
		"2a54f4b4-ee3b-5e7a-a997-e9786525ae79",
		"2465a00f-7ff9-50ba-b919-aa0994d0da74",
		"1b8d2fed-801e-5bea-ae81-7b7ad687a4ad",
		"19110975-ce86-539a-99e7-b65d1d175afa",
		"1451ec4a-b186-55cb-a238-4dae410c51c7",
		"08ec31a0-f9e7-5d17-88a2-df0040a412e1",
		"027fd919-4649-5cb1-819e-fcc0e3d20dbd",
		"01178010-6322-5ee0-9435-9e3a5f95bdb4",
		"0036ce5f-2c73-59c8-a53f-e06b06c11001",
	}

	for _, v := range passes {
		fmt.Printf("start: %v\n", v)
		passDB, err := ps.ProcessRepeatPass(context.Background(), &pass.PassRequest{
			Id: v,
		})

		require.NoError(t, err)
		fmt.Println(passDB.String())
		fmt.Printf("done: %v\n", v)
	}
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
