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
	./bin/test -test.run ^TestSimplePass$

simple_complex:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestSimpleComplexPass$

apm:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestApmGateway$

complex_wrong_time:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestWrongTimeComplexPass$

metro_complex_mcd:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestComplexMCD$

metro_complex_2mcd:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestComplex2MCD$

metro_complex_time_mcd:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestComplexTimeMCD$

metro_complex_mm:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestMetroComplexMM$

metro_complex_mck:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestMetroComplexMCK$

metro_complex_mmts:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestMetroComplexMMTS$

offline_metro_complex_mck_test:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestOfflineMetroComplexMCK$

offline_metro_complex_mm_test:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestOfflineMetroComplexMM$

offline_metro_complex_mmts_test:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestOfflineMetroComplexMMTS$

revise:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestRevise$

resolve:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestResolve$

parking:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestParking$

scope_check:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestScopeCheckPass$


offline: offline_metro_complex_mck_test offline_metro_complex_mm_test offline_metro_complex_mmts_test
complex: metro_complex_mmts metro_complex_mm simple_complex metro_complex_mck metro_complex_mcd metro_complex_2mcd offline scope_check
full: simple simple_complex apm complex parking revise resolve
local: metro_complex_time_mcd
test: metro_complex_mmts
stage: full
prod: full
