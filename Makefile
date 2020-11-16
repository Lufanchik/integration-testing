local: export PROCESSING_API_URL=http://localhost:9090
local: export PASS_URL=http://localhost:13380
local: export APM_API_URL=http://localhost:1340
local: export WEB_API_URL=http://localhost:1344
local: export AUTH_SERVICE_URL=http://localhost:9091
local: export REVISE_SERVICE_URL=http://localhost:1338
local: export RESOLVE_SERVICE_URL=http://localhost:1335
local: export TWPG_SERVICE_URL=http://localhost:1312
local: export COMMENTS_SERVICE_URL=http://localhost:5447
local: export CARD_SERVICE_URL=http://localhost:1480

test: export PROCESSING_API_URL=http://processing-api-gateway.test.svc.cluster.local:9090
test: export PASS_URL=http://pass-service.test.svc.cluster.local:13380
test: export APM_API_URL=http://apm-api-gateway-public.test.svc.cluster.local:1340
test: export WEB_API_URL=http://web-api-gateway.test.svc.cluster.local:1344
test: export AUTH_SERVICE_URL=http://auth-service.test.svc.cluster.local:9091
test: export REVISE_SERVICE_URL=http://revise-service.test.svc.cluster.local:1338
test: export RESOLVE_SERVICE_URL=http://resolve-service.test.svc.cluster.local:1335
test: export TWPG_SERVICE_URL=http://twpg-service.test.svc.cluster.local:1312
test: export COMMENTS_SERVICE_URL=http://comment-service.test.svc.cluster.local:5447
test: export CARD_SERVICE_URL=http://card-service.test.svc.cluster.local:1480

prod: export PASS_URL=http://pass-service.production.svc.cluster.local:13380

full:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestFull$

simple:
	go test -c -o ./bin/test
	./bin/test -test.run ^TestSimple$

local: full
prod: simple
test: simple
