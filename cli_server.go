package main

import (
	"flag"
	"github.com/miaokobot/miaospeed/preconfigs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/miaokobot/miaospeed/service"
	"github.com/miaokobot/miaospeed/utils"
)

func InitConfigServer() *utils.GlobalConfig {
	gcfg := &utils.GCFG
	ecfg := &preconfigs.ECFG

	sflag := flag.NewFlagSet(cmdName+" server", flag.ExitOnError)
	sflag.StringVar(&gcfg.Token, "token", "", "specify the token used to sign request")
	sflag.StringVar(&gcfg.Binder, "bind", "", "bind a socket, can be format like 0.0.0.0:8080 or /tmp/unix_socket")
	sflag.UintVar(&gcfg.ConnTaskTreading, "connthread", 64, "parallel threads when processing normal connectivity tasks")
	sflag.Uint64Var(&gcfg.SpeedLimit, "speedlimit", 0, "speed ratelimit (in Bytes per Second), default with no limits")
	sflag.UintVar(&gcfg.PauseSecond, "pausesecond", 0, "pause such period after each speed job (seconds)")
	sflag.BoolVar(&gcfg.MiaoKoSignedTLS, "mtls", false, "enable miaoko certs for tls verification")
	sflag.BoolVar(&gcfg.NoSpeedFlag, "nospeed", false, "decline all speedtest requests")
	sflag.StringVar(&gcfg.MaxmindDB, "mmdb", "", "reroute all geoip query to local mmdbs. for example: test.mmdb,testcity.mmdb")

	//embeded values
	var ServerPublicKeyPath string
	var ServerPrivateKeyPath string
	var ScriptPredefinedPath string
	var ScriptGeoPath string
	var ScriptIpPath string

	sflag.StringVar(&ecfg.BuildToken, "buildtoken", preconfigs.BUILDTOKEN, "set build token")
	sflag.StringVar(&ServerPublicKeyPath, "serverpublickey", "", "set server public key")
	sflag.StringVar(&ServerPrivateKeyPath, "serverprivatekey", "", "set server private key")
	sflag.StringVar(&ScriptPredefinedPath, "scriptpredefined", "", "set predefined script")
	sflag.StringVar(&ScriptGeoPath, "scriptgeo", "", "set geo script")
	sflag.StringVar(&ScriptIpPath, "scriptip", "", "set ip script")

	whiteList := sflag.String("whitelist", "", "bot id whitelist, can be format like 1111,2222,3333")
	parseFlag(sflag)

	ReadAndFill(&ecfg.ServerPublicKey, ServerPublicKeyPath, preconfigs.MIAOKO_TLS_CRT)
	ReadAndFill(&ecfg.ServerPrivateKey, ServerPrivateKeyPath, preconfigs.MIAOKO_TLS_KEY)
	ReadAndFill(&ecfg.ScriptPredefined, ScriptPredefinedPath, preconfigs.PREDEFINED_SCRIPT)
	ReadAndFill(&ecfg.ScriptGeo, ScriptGeoPath, preconfigs.DEFAULT_GEOIP_SCRIPT)
	ReadAndFill(&ecfg.ScriptIp, ScriptIpPath, preconfigs.DEFAULT_IP_SCRIPT)

	gcfg.WhiteList = make([]string, 0)
	if *whiteList != "" {
		gcfg.WhiteList = strings.Split(*whiteList, ",")
	}

	return gcfg
}

func RunCliServer() {
	InitConfigServer()
	utils.DWarnf("MiaoSpeed speedtesting client %s", utils.VERSION)
	//社区构建的MiiaoSpeed增强版本，由@nodpai制作
	utils.DWarnf("Enhanced version by @nodpai")

	// load maxmind db
	if utils.LoadMaxMindDB(utils.GCFG.MaxmindDB) != nil {
		os.Exit(1)
	}

	// start task server
	go service.StartTaskServer()

	// start api server
	service.CleanUpServer()
	go service.InitServer()

	<-utils.MakeSysChan()

	// clean up
	service.CleanUpServer()
	utils.DLog("shutting down.")
}

func ReadAndFill(p *string, path string, defval string) {
	if path == "" {
		*p = defval
		return
	}

	//read file from path and fill p,if failed,use defval
	file, err := ioutil.ReadFile(path)
	if err != nil {
		utils.DWarnf("failed to read file %s, use default value", path)
		*p = defval
		return
	}
	*p = string(file)
}
