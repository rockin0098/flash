#!/bin/sh

go run TLParser/main.go
gofmt -w proto/mtproto/tl.class.id.go
gofmt -w proto/mtproto/tl.object.all.go
