module lab.siroccotechnology.ru/tp/integration-testing

go 1.12

replace lab.siroccotechnology.ru/tp/common => ../common

replace lab.siroccotechnology.ru/tp/pass-service => ../pass-service

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible
	github.com/gavv/httpexpect v2.0.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/stretchr/testify v1.4.0
<<<<<<<<< Temporary merge branch 1
	github.com/valyala/fasthttp v1.6.0 // indirect
	lab.siroccotechnology.ru/tp/common v0.0.136
	lab.siroccotechnology.ru/tp/pass-service v0.0.70
=========
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	lab.siroccotechnology.ru/tp/common v0.0.137
	lab.siroccotechnology.ru/tp/pass-service v0.0.67
>>>>>>>>> Temporary merge branch 2
)
