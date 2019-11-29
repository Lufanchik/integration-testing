package http_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"lab.siroccotechnology.ru/tp/common/global"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/pass"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/common/messages/registries"
	"lab.siroccotechnology.ru/tp/common/messages/user"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	passService "lab.siroccotechnology.ru/tp/pass-service/proto"
)

var (
	httpProcessingApi *httpexpect.Expect
	httpApmApi        *httpexpect.Expect
	ps                passService.PassService
)

type (
	PaymentType uint8
	RequestType uint8
	AuthType    uint8
	Cases       []Case
	T           []interface{}
)

//варианты шагов, который мы можем использовать в кейсах
type (
	Case struct {
		N string
		T T
	}
	//генерация прохода
	Pass struct {
		//тип оплаты
		PaymentType PaymentType
		//тип прохода
		RequestType RequestType
		//тип авторизации
		AuthType AuthType
		//ид саб перевозчика
		SubCarrier carriers.SubCarrier
		//ид перевозчика
		Carrier carriers.Carrier
		//данные об авторизации
		Auth *processing.Auth
		//сумма, на которую мы ожидаем авторизацию
		ExpectedSum uint32
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		Parent int
		//ссылка на вход, для валидации связи
		Ingress int
		//функция времени
		Now func() uint64
		//время отведенное на запрос в миллисекундах, работает только с онлайном
		Duration uint32
		//данные о терминале
		Terminal *processing.Terminal
		//время, которое мы ждяли перед входом
		TimeToWait time.Duration

		id          string
		carrierID   string
		tapRequest  *processing.TapRequest
		timeRequest uint64
		card        *processing.Card
		parent      *Pass
		ingress     *Pass
		isParent    bool
		timeToWait  time.Duration
	}

	//генерация прохода
	PassCheck struct {
		//тип оплаты
		PaymentType PaymentType
		//сумма, на которую мы ожидаем авторизацию
		ExpectedSum uint32
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		Parent int
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		Target int
	}

	//модификация прохода
	Updater struct {
		//функция модификатора
		f func(tap *pass.Pass)
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		target int
	}

	//получение реестров абс
	AbsGetRegistry struct {
	}

	//логин
	Login struct {
	}

	//отмена
	Cancel struct {
		Target int
		Reason processing.CancelPassRequest_CancelPassReason
	}

	//проверка парковки
	Parking struct {
		rp processing.CheckParkingResponse_Result
		r  *processing.CheckParkingRequest
	}
)

const (
	PaymentTypeFree PaymentType = iota + 1
	PaymentTypePayment
)

const (
	RequestTypeNone RequestType = iota
	RequestTypeOffline
	RequestTypeOnline
)

const (
	AuthTypeCorrect AuthType = iota
	AuthTypeIncorrect
)

func init() {
	ps = passService.NewPassServiceProtobufClient(passUrl, http.DefaultClient)
	gofakeit.Seed(time.Now().UnixNano())
}

func CardExp() string {
	return strconv.Itoa(gofakeit.Number(1000, 9999))
}

func CardBin() uint32 {
	return uint32(gofakeit.Number(1000, 9999))
}

func CardPan() string {
	return strconv.Itoa(gofakeit.CreditCardNumber())
}

func CardEmvCorrect() string {
	return "ca5zfeN8X0VQGfocwEl/8WI1IGyfrOPVST9rtzVBBpPDYVKm/f1/r+gyPwRBANXrWt13q/ADa6VEoQU8u+Og0FPk2IzepajqqPEpxzHGZANjE5fygv7Hk/kblwu6Ktxj76AhU3Te1nlr5LhrfcsOLr3LbEfr2YXOi8GRX2FC/AHNJRumHCyF6r7aBB4EwoAM/gBk53y1TIvM84eq7G4b+z4w0p/le+FXb4DzuryOsj633DVEaWtMbv4A+HgqpdubQyFizkbDmiRLlcEkXztuqEJX/c1jFrR4LhA4/gU9YPcd0YggZQ53gqJ8b57HljbVAosjwuvWE4JaG4sF71sRUg=="
}

func CardEmvIncorrect() string {
	return "WovYAaagxXfxea2U/5cMzTN1MEQIJZAUTovU5NDItD4qB6PyP0rzjg3nnXjKH3XVaVMjUSHJ5YiIJ2Ji4jX3E5bj+Ufe5BntEUhjqyphe8HQz+jJdhWfld+Bm61C8yeeq3qhYbf7zgswrh3d2Gd5L6h4PlKVbbCuGLI8KOmTmfoqQXX5dtF+ZHum7l5BUIvGn3nFj1Fbkye0iKcxHXyToZd/l/M9FuRy9/klAgKYPScYlYWRSwH2I5HZs5qDKJi/cXscLjYoF6h9xxRJMsXJr68BQ5E1bx9sG5mlBb3Pzeytns8Qct9pziuUdDUnFEd2xgx07ul7jNu40k9BARcyfQ=="
}

var Now func() uint64
var NowBackup = func() uint64 {
	return uint64(time.Now().UnixNano())
}

var NowCustom = func(hour, min int) func() uint64 {
	now := time.Now()
	return func() uint64 {
		return uint64(time.Date(
			now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.UTC).UnixNano())
	}
}

var NowCustomDate = func(month, day, hour, min int) func() uint64 {
	now := time.Now()
	return func() uint64 {
		return uint64(time.Date(
			now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.UTC).UnixNano())
	}
}

func Card() *processing.Card {
	return &processing.Card{
		System: processing.CardSystem_VISA,
		Type:   processing.CardType_DEBIT,
		Pan:    CardPan(),
		Bin:    CardBin(),
		Exp:    CardExp(),
		Token: &processing.Token{
			Type: processing.Token_SAMSUNG_PAY,
		},
	}
}

func GenerateEmv(card *processing.Card, p *Pass) {
	switch p.AuthType {
	case AuthTypeCorrect:
		card.Emv = CardEmvCorrect()
	case AuthTypeIncorrect:
		card.Emv = CardEmvIncorrect()
	}
}

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
	pr.r.Pan = card.Pan

	response := &processing.CheckParkingResponse{
		Result: pr.rp,
	}

	return pr.r, response
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

func logRequest(u string, r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("%s, trace id: %s", u, trace.Raw()))
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

func ValidatePass(t *testing.T, p *Pass, parent *Pass, ingress *Pass) {
	ctx := context.Background()
	passDB, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: p.id,
	})
	require.NoError(t, err)

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
		Sum:               p.ExpectedSum,
		IsAuth:            false,
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
	}

	if p.Parent > 0 {
		expectPass.IsComplex = true
		expectPass.ParentComplexId = parent.id
	}

	if p.AuthType == AuthTypeIncorrect {
		expectPass.IsAuth = false
	}

	require.Equal(t, expectPass, passDB)
	require.NoError(t, err)
}

func AuthStatus(t *testing.T, p *Pass) {
	req, resp := AuthStatusRequest(p)
	u := "/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/AuthStatus"
	r := httpProcessingApi.POST(u).WithJSON(req).
		Expect().
		Status(http.StatusOK)

	object := r.Body().Raw()
	logRequest(u, r)

	response := &processing.AuthResponse{}
	err := jsonpb.Unmarshal(strings.NewReader(object), response)
	require.NoError(t, err)

	resp.Created = response.Created
	resp.Info = response.Info
	require.Equal(t, resp, response)
}

func NanoToMicro(tm uint64) uint64 {
	return tm / uint64(time.Microsecond) * 1000
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
	ValidatePass(t, target, parent, nil)
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

func Run(t *testing.T, cases Cases) {
	httpProcessingApi = httpexpect.New(t, processingApiUrl)
	httpApmApi = httpexpect.New(t, apmApiUrl)
	t.Parallel()
	for _, scenario := range cases {
		card := Card()
		carrierID := uuid.New().String()
		fmt.Println(scenario.N)
		fmt.Println(card.String())
		for N, step := range scenario.T {
			fmt.Println(N + 1)
			t.Run("case: "+scenario.N, func(t *testing.T) {
				//Pass
				p, ok := step.(*Pass)
				if ok {
					if p.Now != nil {
						Now = p.Now
					} else {
						Now = NowBackup
					}
					if p.RequestType == RequestTypeNone {
						p.RequestType = globalRequestType
					}
					GenerateEmv(card, p)
					p.carrierID = carrierID
					tapReq := TapBySubCarrier(t, p, card)
					timeRequest := PassBySubCarrier(t, tapReq, p)
					var parent, ingress *Pass
					if p.Parent > 0 {
						pr, ok := (scenario.T[p.Parent-1]).(*Pass)
						if !ok {
							t.Fail()
						}
						parent = pr
						parent.isParent = true
						parent.timeToWait = parent.TimeToWait
					}
					if p.Ingress > 0 {
						ing, ok := (scenario.T[p.Ingress-1]).(*Pass)
						if !ok {
							t.Fail()
						}
						ingress = ing
						ingress.timeToWait = ingress.TimeToWait
					}
					p.tapRequest = tapReq
					p.timeRequest = timeRequest
					p.card = card
					p.parent = parent
					p.ingress = ingress

					time.Sleep(time.Millisecond * 200)

					ValidatePass(t, p, parent, ingress)
					AuthStatus(t, p)

					if parent != nil {
						ValidatePass(t, parent, parent.parent, parent.ingress)
						AuthStatus(t, parent)
					}

					if ingress != nil {
						ValidatePass(t, ingress, ingress.parent, ingress.ingress)
						AuthStatus(t, ingress)
					}
				}

				//Updater
				u, ok := step.(Updater)
				if ok {
					pu, ok := (scenario.T[u.target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					Update(t, pu, u)
				}

				//AbsGetRegistry
				agr, ok := step.(*AbsGetRegistry)
				if ok {
					AbsGetRegistryApi(t, agr)
				}

				//AbsGetRegistry
				lg, ok := step.(*Login)
				if ok {
					LoginApi(t, lg)
				}

				//PassCheck
				pc, ok := step.(*PassCheck)
				if ok {
					target, ok := (scenario.T[pc.Target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					parent, ok := (scenario.T[pc.Parent-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					time.Sleep(time.Millisecond * 200)
					PassCheckApi(t, pc, target, parent)
				}

				//Cancel
				cl, ok := step.(*Cancel)
				if ok {
					target, ok := (scenario.T[cl.Target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					CancelApi(t, cl, target)
				}

				//Parking
				pr, ok := step.(*Parking)
				if ok {
					ParkingApi(t, card, pr)
				}
			})
			if t.Failed() {
				break
			}
		}
		if t.Failed() {
			break
		}
	}
}
