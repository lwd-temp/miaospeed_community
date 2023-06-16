package preconfigs

type EmbedConfig struct {
	BuildToken       string //utils.BUILDTOKEN
	ServerPublicKey  string //preconfigs.MIAOKO_TLS_CRT
	ServerPrivateKey string //preconfigs.MIAOKO_TLS_KEY
	ScriptPredefined string //engine.PREDEFINED_SCRIPT
	ScriptGeo        string //engine.DEFAULT_GEOIP_SCRIPT
	ScriptIp         string //engine.DEFAULT_IP_SCRIPT
}

var ECFG EmbedConfig
