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
	github.com/valyala/fasthttp v1.6.0 // indirect
	lab.siroccotechnology.ru/tp/common v0.0.136
	lab.siroccotechnology.ru/tp/pass-service v0.0.70
)
