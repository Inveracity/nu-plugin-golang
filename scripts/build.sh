set -e
go build -o nu_plugin_gotry *.go
nu -c "register nu_plugin_gotry"
nu -c "nu-golang 2 234"
