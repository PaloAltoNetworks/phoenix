package main

import (
	"crypto/tls"
	"time"

	"github.com/aporeto-inc/phoenix"
)

func main() {

	cfg := NewConfiguration()
	time.Local = time.UTC

	phoenix.NewServer(
		phoenix.NewHooksRegistry(examplePluginFunc),
		":44334",
		cfg.CAPool,
		[]tls.Certificate{cfg.ServerCertificate},
	).Start()

}
