package resolve_test

import (
	"github.com/gavv/httpexpect"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/require"
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/crud"
	"lab.siroccotechnology.ru/tp/common/messages/registries"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"lab.siroccotechnology.ru/tp/integration-testing/user"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestReviseGetTaskList(t *testing.T) {
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

	t.Run("ReviseGetTaskList", func(t *testing.T) {
		_ = httpApi.POST("/twirp/proto.ApmAPIGateway/ReviseGetTaskList").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsKey("list")
	})
}

func TestReviseCreateTask(t *testing.T) {
	var (
		AT  string
		err error
	)
	if AT, err = user.GetAccessToken(t); err != nil {
		log.Error("Error to get Access Token")
		t.Error(err)
		return
	}

	orderNum := time.Now().Unix()

	httpApi := httpexpect.New(t, test.ApmApiUrl)
	request := &registries.StartReviseRequest{
		CarrierID: []carriers.Carrier{
			carriers.Carrier_MM,
		},
		OrderNum: uint64(orderNum),
		Date:     974906537000,
	}

	t.Run("ReviseCreateTask", func(t *testing.T) {
		_ = httpApi.POST("/twirp/proto.ApmAPIGateway/ReviseCreateTask").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Empty()
	})
}

func TestReviseResetTask(t *testing.T) {
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

	var obj *httpexpect.Object

	t.Run("ReviseGetTaskList", func(t *testing.T) {
		obj = httpApi.POST("/twirp/proto.ApmAPIGateway/ReviseGetTaskList").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsKey("list")
	})

	list := obj.Value("list").Array().Iter()

	for _, v := range list {
		raw := v.Raw()
		i, ok := raw.(map[string]interface{})

		require.True(t, ok)

		status, ok := i["status"]
		require.True(t, ok)

		if status == "error" {

			taskId, ok := i["Id"]
			require.True(t, ok)

			taskIdStr, ok := (taskId).(string)
			require.True(t, ok)
			taskInt, err := strconv.Atoi(taskIdStr)
			require.NoError(t, err)

			resetTaskrq := &registries.ResetTaskRequest{Id: uint64(taskInt)}

			t.Run("ReviseResetTask", func(t *testing.T) {
				obj = httpApi.
					POST("/twirp/proto.ApmAPIGateway/ReviseResetTask").
					WithHeader("Authorization", "Bearer "+AT).
					WithJSON(resetTaskrq).
					Expect().
					Status(http.StatusOK).
					JSON().Object().Empty()
			})

			return
		}
	}
}

func TestReviseGetTaskStatus(t *testing.T) {
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
		Limit:  3,
		Offset: 0,
	}

	var obj *httpexpect.Object

	t.Run("ReviseGetTaskList", func(t *testing.T) {
		obj = httpApi.POST("/twirp/proto.ApmAPIGateway/ReviseGetTaskList").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsKey("list")
	})

	list := obj.Value("list").Array().Iter()

	for _, v := range list {
		raw := v.Raw()
		i, ok := raw.(map[string]interface{})

		require.True(t, ok)

		var (
			status, taskId interface{}
			taskIdStr      string
		)
		status, ok = i["status"]
		taskId, ok = i["Id"]
		require.True(t, ok)

		taskIdStr, ok = (taskId).(string)
		require.True(t, ok)
		taskInt, err := strconv.Atoi(taskIdStr)
		require.NoError(t, err)

		getTaskStatusRq := &registries.GetReviseTaskStatusRequest{
			TaskId: int64(taskInt),
		}

		t.Run("ReviseResetTask", func(t *testing.T) {
			r := httpApi.
				POST("/twirp/proto.ApmAPIGatewayPublic/ReviseGetTaskStatus").
				WithHeader("Authorization", "Bearer "+AT).
				WithJSON(getTaskStatusRq).
				Expect().
				Status(http.StatusOK)

			raw := r.Body().Raw()
			response := &registries.GetReviseTaskStatusResponse{}
			err := jsonpb.Unmarshal(strings.NewReader(raw), response)
			require.NoError(t, err)

			require.True(t, response.Status == status)
		})
	}
}

func TestReviseGetTaskStages(t *testing.T) {
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
	request := &registries.EmptyMessage{}

	t.Run("ReviseGetTaskList", func(t *testing.T) {
		_ = httpApi.POST("/twirp/proto.ApmAPIGateway/ReviseGetStages").
			WithHeader("Authorization", "Bearer "+AT).
			WithJSON(request).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsKey("stages")
	})
}
