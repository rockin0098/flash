#!/bin/sh

echo '###### building All ... ...'
go install -v -ldflags "-w -s" github.com/rockin0098/meow/GateServer
go install -v -ldflags "-w -s" github.com/rockin0098/meow/TLParser
