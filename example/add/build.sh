#!/bin/bash
set -e

# Build the binary
go build -o nu_plugin_golang *.go

# Register it in Nushell
nu -c "register nu_plugin_golang"

# Run it
nu -c "nu-golang 2 234"
