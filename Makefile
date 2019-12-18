local: export PROCESSING_API_URL=http://localhost:9090
local: export PASS_URL=http://localhost:13380
local: export APM_API_URL=http://localhost:1340
local: export WEB_API_URL=http://localhost:1344

test: export PROCESSING_API_URL=http://processing-api-gateway.test.svc.cluster.local:9090
test: export PASS_URL=http://pass-service.test.svc.cluster.local:13380
test: export APM_API_URL=http://apm-api-gateway-public.test.svc.cluster.local:1340
test: export WEB_API_URL=http://web-api-gateway.test.svc.cluster.local:1344

full:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestFull$

simple:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestSimple$

local: full
test: simple
