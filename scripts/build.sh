set -e
go build -o nu_plugin_golang *.go
nu -c "register nu_plugin_golang"
nu -c "nu-golang 2 234"
