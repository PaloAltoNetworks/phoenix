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

	// Bahamut API Server configuration
	bahamutConfig := bahamut.Config{}
	bahamutConfig.ReSTServer.Disabled = false
	bahamutConfig.ReSTServer.DisableKeepalive = false
	bahamutConfig.ReSTServer.IdleTimeout = 240 * time.Second
	bahamutConfig.ReSTServer.ListenAddress = listenAddress
	bahamutConfig.ReSTServer.ReadTimeout = 60 * time.Second
	bahamutConfig.ReSTServer.WriteTimeout = 120 * time.Second
	bahamutConfig.TLS.ClientCAPool = caPool
	bahamutConfig.TLS.RootCAPool = caPool
	bahamutConfig.TLS.ServerCertificates = serverCertificates
	bahamutConfig.TLS.AuthType = tls.RequireAndVerifyClientCert
	bahamutConfig.HealthServer.Disabled = !enableHealth
	bahamutConfig.HealthServer.HealthHandler = healthHandlerFunc
	bahamutConfig.HealthServer.ListenAddress = healtListenAddress
	bahamutConfig.Model.RelationshipsRegistry = map[int]elemental.RelationshipsRegistry{0: gaia.Relationships()}
	bahamutConfig.Model.IdentifiablesFactory = func(i string, version int) elemental.Identifiable { return gaia.IdentifiableForIdentity(i) }
	bahamutConfig.PushServer.Disabled = true

	// Create a new Bahamut Server
	server := bahamut.NewServer(bahamutConfig)

	// Register all the processors
	bahamut.RegisterProcessorOrDie(server, newRemoteProcessorProcessor(pluginsRegistry), gaia.RemoteProcessorIdentity)

	return server
}
