package utils

import "crypto/tls"

func CreateTlsConfig(cert []byte, pkey []byte) (tlsConfig *tls.Config, err error) {
	config := tls.Config{
		Time: nil,
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.X509KeyPair(cert, pkey)
	if err != nil {
		println(err.Error())
		return
	}
	tlsConfig = &config
	return
}
