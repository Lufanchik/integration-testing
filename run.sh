#!/usr/bin/env bash

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/apm-api-gateway && \
make run &> /home/vorozhcovgs/logs/apm-api-gateway.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/user-service && \
make run &> /home/vorozhcovgs/logs/user-service.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/tidpool-service && \
make run &> /home/vorozhcovgs/logs/tidpool-service.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/mk-emulator && \
make run &> /home/vorozhcovgs/logs/mk-emulator.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/processing-api-gateway && \
make run &> /home/vorozhcovgs/logs/processing-api-gateway.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/pass-service && \
make run &> /home/vorozhcovgs/logs/pass-service.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/calculator-service && \
make run &> /home/vorozhcovgs/logs/calculator-service.logs && \

cd /home/vorozhcovgs/go/src/lab.siroccotechnology.ru/tp/auth-service && \
make run &> /home/vorozhcovgs/logs/auth-service.logs