package registry

import (
	"github.com/gavv/httpexpect"
	"github.com/prometheus/common/log"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/crud"
	"lab.siroccotechnology.ru/tp/common/messages/registries"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"lab.siroccotechnology.ru/tp/integration-testing/user"
	"net/http"
	"testing"
)

func TestResolveGetTaskList(t *testing.T) {
	var (
		AT  string
		err error
	)
	if AT, err = user.GetAccessToken(t); err != nil {
		log.Error("Error to get Access Token")
		t.Error(err)
		return
	}

	httpApi := httpexpect.New(t, test.ApmApiUrl)
	request := &crud.EntitiesRequest{
		Filters: []*crud.Filter{
			&crud.Filter{
				Field:    "stage",
				Value:    "copy_to_storage",
				Operator: 3,
			},
		},
		Limit:  10,
		Offset: 0,
	}

	t.Run("ResolveGetTaskList", func(t *testing.T) {
		_ = httpApi.POST("/twirp/proto.ApmAPIGateway/ResolveGetTaskList").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsKey("list")
	})
}

func TestResolveCreateTask(t *testing.T) {
	var (
		AT  string
		err error
	)
	if AT, err = user.GetAccessToken(t); err != nil {
		log.Error("Error to get Access Token")
		t.Error(err)
		return
	}

	httpApi := httpexpect.New(t, test.ApmApiUrl)
	request := &registries.StartResolveRequest{
		CarrierID: []carriers.Carrier{
			carriers.Carrier_MM,
		},
		Date: 974906537000,
	}

	t.Run("ResolveCreateTask", func(t *testing.T) {
		_ = httpApi.POST("/twirp/proto.ApmAPIGateway/ResolveCreateTask").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Empty()
	})
}
