#!/usr/bin/env bash

go get -u github.com/davelaursen/present-plus
go get -u ./...
#present -http '127.0.0.1:4000' -notes
present-plus