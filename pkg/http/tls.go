package http

import (
	"net"

	tls "github.com/refraction-networking/utls"
)

func dialTLS(network, addr string) (net.Conn, error) {
	netConn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	config := tls.Config{
		CipherSuites: []uint16{tls.TLS_AES_128_GCM_SHA256},
		MinVersion:   tls.VersionTLS13,
		ServerName: host,
	}

	tlsConn := tls.UClient(netConn, &config, tls.HelloGolang)
	if err = tlsConn.Handshake(); err != nil {
		return nil, err
	}

	return tlsConn, nil
}
