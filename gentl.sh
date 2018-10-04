#!/bin/sh

go run TLParser/main.go
gofmt -w proto/mtproto/tl/tl.class.id.go
