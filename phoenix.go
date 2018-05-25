package phoenix

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/aporeto-inc/bahamut"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang"
)

// NewServer returns a new phoenix server with the given hook.
func NewServer(
	pluginsRegistry HooksRegistry,
	listenAddress string,
	caPool *x509.CertPool,
	serverCertificates []tls.Certificate,
	enableHealth bool,
	healthHandlerFunc bahamut.HealthServerFunc,
	healtListenAddress string,
) bahamut.Server {

	time.Local = time.UTC
	factory := func(i string, version int) elemental.Identifiable { return gaia.IdentifiableForIdentity(i) }
	registry := map[int]elemental.RelationshipsRegistry{0: gaia.Relationships()}

	options := []bahamut.Option{
		bahamut.OptRestServer(listenAddress),
		bahamut.OptTLS(serverCertificates, nil),
		bahamut.OptMTLS(caPool, tls.RequireAndVerifyClientCert),
		bahamut.OptModel(factory, registry),
		bahamut.OptTimeouts(60*time.Second, 120*time.Second, 240*time.Second),
	}

	if enableHealth {
		options = append(options, bahamut.OptHealthServer(healtListenAddress, healthHandlerFunc))
	}

	// Create a new Bahamut Server
	server := bahamut.New(options...)

	// Register all the processors
	bahamut.RegisterProcessorOrDie(server, newRemoteProcessorProcessor(pluginsRegistry), gaia.RemoteProcessorIdentity)

	return server
}
