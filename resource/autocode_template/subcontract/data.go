package subcontract

import _ "embed"

//go:embed service_enter.go.tpl
var Server []byte

//go:embed router_enter.go.tpl
var Router []byte

//go:embed api_enter.go.tpl
var API []byte
