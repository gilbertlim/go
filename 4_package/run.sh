#!/bin/bash

go module init package-test

go get github.com/google/uuid # get external package
# go.sum : 외부 패키지 버전에 대한 checksum 파일

go build -o package-test-0616

./package-test-0616
