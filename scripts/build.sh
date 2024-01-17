#!/bin/sh
go build -o bin/nu_plugin_addv2 cmd/add/*.go

nu -c "register bin/nu_plugin_addv2"

nu -c "addv2 --help"
nu -c "addv2 1 2"
