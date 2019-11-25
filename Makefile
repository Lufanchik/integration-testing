local: export PROCESSING_API_URL=http://localhost:9090
local: export PASS_URL=http://localhost:13380
local: export APM_API_URL=http://localhost:1340

stage: export PROCESSING_API_URL=http://processing-api-gateway.stage.svc.cluster.local:9090
stage: export PASS_URL=http://pass-service.stage.svc.cluster.local:13380
stage: export APM_API_URL=http://apm-api-gateway-public.stage.svc.cluster.local:1340

test: export PROCESSING_API_URL=http://processing-api-gateway.test.svc.cluster.local:9090
test: export PASS_URL=http://pass-service.test.svc.cluster.local:13380
test: export APM_API_URL=http://apm-api-gateway-public.test.svc.cluster.local:1340

prod: export PROCESSING_API_URL=http://processing-api-gateway.production.svc.cluster.local:9090
prod: export PASS_URL=http://pass-service.production.svc.cluster.local:13380
prod: export APM_API_URL=http://apm-api-gateway-public.production.svc.cluster.local:1340

simple:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestSimplePass$

simple_complex:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestSimpleComplexPass$

apm:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestApmGateway$

complex_wrong_time:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestWrongTimeComplexPass$

metro_complex_mcd:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestComplexMCD$

metro_complex_mcd_2:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestComplexMCD2$

metro_complex_mm:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestMetroComplexMM$

metro_complex_mck:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestMetroComplexMCK$

metro_complex_mmts:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestMetroComplexMMTS$

revise:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestRevise$

resolve:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestResolve$

parking:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestParking$

complex: simple_complex metro_complex_mm metro_complex_mck metro_complex_mmts metro_complex_mcd metro_complex_mcd_2
full: simple simple_complex apm complex parking revise resolve
local: resolve
test: full
stage: full
prod: full
