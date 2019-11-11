localhost: export PROCESSING_API_URL=http://localhost:9090
localhost: export PASS_URL=http://localhost:13380
localhost: export APM_API_URL=http://localhost:1340

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

apm:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestApmGateway$

complex_mck:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestComplexPassMCK$

complex_mm:
	go test -c -o ./bin/test
	./bin/test -test.v -test.run ^TestComplexPassMM$

full: simple apm complex_mck complex_mm

localhost: full
test: full
stage: full
prod: full