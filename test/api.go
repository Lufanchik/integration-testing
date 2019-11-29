package test

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
	"lab.siroccotechnology.ru/tp/common/global"
	"lab.siroccotechnology.ru/tp/common/messages/pass"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/common/messages/registries"
	"lab.siroccotechnology.ru/tp/common/messages/user"
	"net/http"
	"strings"
	"testing"
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
