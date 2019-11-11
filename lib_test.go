package http_test

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	Cases       [][]interface{}
)

//варианты шагов, который мы можем использовать в кейсах
type (
	//генерация прохода
	Pass struct {
		//сгенерированный ид прохода, заполняется автоматически после создания
		ID string
		//сгенерированный ид прохода от перевозчика, заполняется автоматически после создания
		CarrierID string
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
		//функция времени
		Now func() uint64
		//время отведенное на запрос в миллисекундах, работает только с онлайном
		Duration uint32
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
)

const (
	PaymentTypeFree PaymentType = iota + 1
	PaymentTypeFullPayment
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
	case PaymentTypeFullPayment:
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

func AuthStatusRequest(p *Pass) (*processing.AuthRequest, *processing.AuthResponse) {
	request := &processing.AuthRequest{
		Id: p.ID,
	}

	response := &processing.AuthResponse{
		Result:     processing.AuthResponse_SUCCESS_RESULT,
		Resolution: processing.AuthResponse_AUTHORIZED,
	}
	switch p.PaymentType {
	case PaymentTypeFullPayment:
		response.Status = processing.AuthResponse_SUCCESS_STATUS
		response.Auth = &processing.Auth{
			Sum:  p.ExpectedSum,
			Type: processing.Auth_CLASSIC,
		}
	case PaymentTypeFree:
		response.Status = processing.AuthResponse_SUCCESS_FREE
	}

	if p.AuthType == AuthTypeIncorrect {
		response.Resolution = processing.AuthResponse_FAILURE
	}

	return request, response
}

func TapRequest(c carriers.SubCarrier, card *processing.Card, carrierID string) (*processing.TapRequest, *processing.TapResponse) {
	timeNow := Now()
	request := &processing.TapRequest{
		Id:      carrierID,
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
	response := &processing.TapResponse{
		Id:      "",
		Created: 0,
		Result:  processing.TapResponse_SUCCESS,
		Msg:     "",
	}
	return request, response
}

func TapBySubCarrier(t *testing.T, p *Pass, card *processing.Card) *processing.TapRequest {
	req, resp := TapRequest(p.SubCarrier, card, p.CarrierID)

	t.Run(p.Carrier.String()+"/twirp/sirocco.ProcessingAPI/ProcessTap - sub carrier: "+p.SubCarrier.String(), func(t *testing.T) {
		object := httpProcessingApi.POST("/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessTap").WithJSON(req).
			Expect().
			Status(http.StatusOK).Body().Raw()

		response := &processing.TapResponse{}
		err := jsonpb.Unmarshal(strings.NewReader(object), response)
		require.NoError(t, err)

		resp.Id = response.Id
		resp.Created = response.Created
		p.ID = response.Id

		assert.Equal(t, resp, response)
	})

	return req
}

func PassBySubCarrier(t *testing.T, tap *processing.TapRequest, p *Pass) uint64 {
	var requestedTime uint64
	switch p.RequestType {
	case RequestTypeOffline:
		req, resp := PassOfflineRequest(tap, p)
		requestedTime = req.Created
		t.Run(p.Carrier.String()+"/twirp/sirocco.ProcessingAPI/ProcessOfflinePass - sub carrier: "+p.SubCarrier.String(), func(t *testing.T) {
			object := httpProcessingApi.POST("/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessOfflinePass").WithJSON(req).
				Expect().
				Status(http.StatusOK).Body().Raw()

			response := &processing.OfflinePassResponse{}
			err := jsonpb.Unmarshal(strings.NewReader(object), response)
			require.NoError(t, err)

			resp.Id = response.Id
			resp.Created = response.Created
			p.ID = response.Id

			assert.Equal(t, resp, response)
		})
	case RequestTypeOnline:
		req, resp := PassOnlineRequest(tap, p)
		requestedTime = req.Created
		t.Run(p.Carrier.String()+"/twirp/sirocco.ProcessingAPI/ProcessOnlinePass - sub carrier: "+p.SubCarrier.String(), func(t *testing.T) {
			object := httpProcessingApi.POST("/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/ProcessOnlinePass").WithJSON(req).
				Expect().
				Status(http.StatusOK).Body().Raw()

			response := &processing.OnlinePassResponse{}
			err := jsonpb.Unmarshal(strings.NewReader(object), response)
			require.NoError(t, err)

			resp.Id = response.Id
			resp.Created = response.Created
			p.ID = response.Id

			assert.Equal(t, resp, response)
		})
	}
	return requestedTime
}

func Update(t *testing.T, p *Pass, up Updater) {
	ctx := context.Background()
	tap, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: p.ID,
	})
	require.NoError(t, err)
	up.f(tap)
	_, err = ps.UpdatePass(ctx, tap)
	require.NoError(t, err)
}

func ValidatePass(t *testing.T, tap *processing.TapRequest, card *processing.Card, p *Pass, parent *Pass, timeRequest uint64) {
	ctx := context.Background()
	passDB, err := ps.GetPass(ctx, &pass.PassRequest{
		Id: p.ID,
	})
	require.NoError(t, err)

	expectPass := &pass.Pass{
		Id:                p.ID,
		UserId:            card.Pan,
		Kind:              0,
		CarrierTapId:      tap.Id,
		CarrierCode:       p.Carrier,
		CarrierResolution: tap.Tap.Resolution,
		TerminalId:        tap.Tap.Terminal.Id,
		TerminalStation:   tap.Tap.Terminal.Station,
		TerminalDirection: tap.Tap.Terminal.Direction,
		Sign:              tap.Tap.Sign,
		IsCancel:          false,
		IsComplexTimeout:  false,
		CreatedAtRequest:  NanoToMicro(timeRequest),
		CreatedAtCarrier:  NanoToMicro(tap.Tap.Created),
		IsComplex:         false,
		ParentComplexId:   "",
		IsComplexCarrier:  global.IsComplexCarrier(p.Carrier),
		IsPass:            true,
		IsFree:            false,
		CarrierCodeSub:    p.SubCarrier,
		Sum:               p.ExpectedSum,
		IsAuth:            false,
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
	case PaymentTypeFullPayment:
		expectPass.IsFree = false
		expectPass.IsAuth = true
	}

	if p.Parent > 0 {
		expectPass.IsComplex = true
		expectPass.ParentComplexId = parent.ID
	}

	if p.AuthType == AuthTypeIncorrect {
		expectPass.IsAuth = false
	}

	assert.Equal(t, expectPass, passDB)
	require.NoError(t, err)
}

func AuthStatus(t *testing.T, p *Pass) {
	req, resp := AuthStatusRequest(p)
	t.Run(p.Carrier.String()+"/twirp/sirocco.ProcessingAPI/AuthStatus - sub carrier: "+p.SubCarrier.String(), func(t *testing.T) {
		object := httpProcessingApi.POST("/" + p.Carrier.String() + "/twirp/sirocco.ProcessingAPI/AuthStatus").WithJSON(req).
			Expect().
			Status(http.StatusOK).Body().Raw()

		response := &processing.AuthResponse{}
		err := jsonpb.Unmarshal(strings.NewReader(object), response)
		require.NoError(t, err)

		resp.Created = response.Created
		resp.Info = response.Info
		assert.Equal(t, resp, response)
	})
}

func NanoToMicro(tm uint64) uint64 {
	return tm / uint64(time.Microsecond) * 1000
}

func AbsGetRegistryApi(t *testing.T, registry *AbsGetRegistry) {
	req := registries.AbsRegistryRequest{}
	t.Run("twirp/proto.ApmAPIGateway/AbsGetRegistry", func(t *testing.T) {
		httpApmApi.POST("/twirp/proto.ApmAPIGateway/AbsGetRegistry").WithJSON(req).
			Expect().
			Status(http.StatusForbidden)
	})
}

func LoginApi(t *testing.T, lg *Login) {
	req := user.LoginRequest{}
	//resp := user.JWTResponse{}
	t.Run("twirp/proto.ApmAPIGatewayPublic/Login", func(t *testing.T) {
		_ = httpApmApi.POST("/twirp/proto.ApmAPIGatewayPublic/Login").WithJSON(req).
			Expect().
			Status(http.StatusNotFound).Body().Raw()
	})
}

func Run(t *testing.T, cases Cases) {
	httpProcessingApi = httpexpect.New(t, processingApiUrl)
	httpApmApi = httpexpect.New(t, apmApiUrl)
	for _, scenario := range cases {
		card := Card()
		carrierID := uuid.New().String()
		for _, step := range scenario {
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
				p.CarrierID = carrierID
				tapReq := TapBySubCarrier(t, p, card)
				timeRequest := PassBySubCarrier(t, tapReq, p)
				var parent *Pass
				if p.Parent > 0 {
					pr, ok := (scenario[p.Parent-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					parent = pr
				}
				//ждем сообщения из rabbit mq
				time.Sleep(time.Millisecond * 400)
				ValidatePass(t, tapReq, card, p, parent, timeRequest)
				AuthStatus(t, p)
			}

			//Updater
			u, ok := step.(Updater)
			if ok {
				pu, ok := (scenario[u.target-1]).(*Pass)
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
		}
	}
}
