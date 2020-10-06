package registry

import (
	"lab.dt.multicarta.ru/tp/common/global"
	"lab.dt.multicarta.ru/tp/common/messages/crud"
	"lab.dt.multicarta.ru/tp/common/messages/registries"
	"lab.dt.multicarta.ru/tp/common/messages/response"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"net/http"
)

var CasesReviseGetTaskList = test.Cases{
	{
		N: "200 get service schema",
		T: test.T{
			&test.Revise{
				Url:      "/twirp/proto.ReviseHttp/GetServiceSchema",
				Status:   http.StatusOK,
				Request:  &response.EmptyMessage{},
				Response: &crud.ServiceSchema{},
			},
		},
	},
	/*{
		N: "200 get task list",
		T: test.T{
			&test.Revise{
				Url:    "/twirp/proto.ReviseHttp/GetTasks",
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
	},*/
	{
		N: "200 get expires link",
		T: test.T{
			&test.Revise{
				Url:      "/twirp/proto.ReviseHttp/GetExpiresLink",
				Status:   http.StatusOK,
				Request:  &response.EmptyMessage{},
				Response: &registries.ExpiresLink{},
			},
		},
	},
	{
		N: "404 Reset Task",
		T: test.T{
			&test.Revise{
				Url:    "/twirp/proto.ReviseHttp/ResetTask",
				Status: http.StatusNotFound,
				Request: &registries.ResetTaskRequest{
					Id: 9999999,
				},
				Response: &registries.EmptyMessage{},
			},
		},
	},
	{
		N: "200 get stages",
		T: test.T{
			&test.Revise{
				Url:      "/twirp/proto.ReviseHttp/GetStages",
				Status:   http.StatusOK,
				Request:  &registries.EmptyMessage{},
				Response: &registries.GetReviseStagesResponse{},
			},
		},
	},
	{
		N: "404 get task status",
		T: test.T{
			&test.Revise{
				Url:    "/twirp/proto.ReviseHttp/GetTaskStatus",
				Status: http.StatusNotFound,
				Request: &registries.GetReviseTaskStatusRequest{
					TaskId: 9999999,
				},
				Response: &registries.GetReviseTaskStatusResponse{},
			},
		},
	},
	{
		N: "200 get Last By Order FileInfo",
		T: test.T{
			&test.Revise{
				Url:    "/twirp/proto.ReviseHttp/GetLastByOrderFileInfo",
				Status: http.StatusOK,
				Request: &registries.ReviseGetFileInfoRequest{
					FileType:    1,
					CarrierCode: 1,
					Date: &global.Date{
						Year:  1999,
						Month: 10,
						Day:   10,
					},
				},
				Response: &registries.ReviseGetFileInfoResponse{},
			},
		},
	},
	/*{
		N: "200 playground",
		T: test.T{
			&test.Revise{
				Url:    "/twirp/proto.ReviseHttp/Playground",
				Status: http.StatusOK,
				Request: &crud.EntitiesRequestV1{
					MainFields: &crud.Main{
						TableName: "task_manager",
						Fields:    []string{"id"},
					},
					Limit:  10,
					Offset: 0,
				},
				Response: &crud.Playground{},
			},
		},
	},*/
}
