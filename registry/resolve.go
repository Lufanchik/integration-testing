package registry

import (
	"lab.siroccotechnology.ru/tp/common/messages/crud"
	"lab.siroccotechnology.ru/tp/common/messages/registries"
	"lab.siroccotechnology.ru/tp/common/messages/response"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"net/http"
)

var CasesResolveGetTaskList = test.Cases{
	{
		N: "200 get service schema",
		T: test.T{
			&test.Resolve{
				Url:      "/twirp/proto.ResolveHttp/GetServiceSchema",
				Status:   http.StatusOK,
				Request:  &response.EmptyMessage{},
				Response: &crud.ServiceSchema{},
			},
		},
	},
	{
		N: "200 get expires link",
		T: test.T{
			&test.Resolve{
				Url:      "/twirp/proto.ResolveHttp/GetExpiresLink",
				Status:   http.StatusOK,
				Request:  &response.EmptyMessage{},
				Response: &registries.ExpiresLink{},
			},
		},
	},
	{
		N: "200 get tasks",
		T: test.T{
			&test.Resolve{
				Url:    "/twirp/proto.ResolveHttp/GetTasks",
				Status: http.StatusOK,
				Request: &crud.EntitiesRequestV1{
					MainFields: &crud.Main{
						TableName: "task_manager",
						Fields:    []string{"id"},
					},
					Limit:  10,
					Offset: 0,
				},
				Response: &crud.EntitiesResponse{},
			},
		},
	},
}
