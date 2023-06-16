package preconfigs

import (
	_ "embed"
)

//go:embed embeded/miaokoCA/miaoko.crt
var MIAOKO_TLS_CRT string

//go:embed embeded/miaokoCA/miaoko.key
var MIAOKO_TLS_KEY string

//go:embed embeded/ca-certificates.crt
var MIAOKO_ROOT_CA []byte

//go:embed embeded/predefined.js
var PREDEFINED_SCRIPT string

//go:embed embeded/default_geoip.js
var DEFAULT_GEOIP_SCRIPT string

//go:embed embeded/default_ip.js
var DEFAULT_IP_SCRIPT string

//go:embed embeded/BUILDTOKEN.key
var BUILDTOKEN string
