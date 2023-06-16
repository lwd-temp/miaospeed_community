package preconfigs

import (
	"crypto/tls"
	"crypto/x509"
)

func MakeSelfSignedTLSServer() *tls.Config {
	cert, _ := tls.X509KeyPair([]byte(ECFG.ServerPublicKey), []byte(ECFG.ServerPrivateKey))

	// Construct a tls.config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// Other options
	}

	return tlsConfig
}

func MiaokoRootCAPrepare() *x509.CertPool {
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(MIAOKO_ROOT_CA)
	return rootCAs
}
