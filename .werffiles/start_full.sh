#!/usr/bin/env bash
go test -c -o ./bin/test
./bin/test -test.run ^TestFull$