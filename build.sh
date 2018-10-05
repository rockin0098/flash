#!/bin/sh

echo '###### building All ... ...'
go install -v -ldflags "-w -s" github.com/rockin0098/flash/GateServer
go install -v -ldflags "-w -s" github.com/rockin0098/flash/TLParser
