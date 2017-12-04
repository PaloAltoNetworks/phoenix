package phoenix

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/aporeto-inc/bahamut"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/rufusmodels/v1/golang"
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
	bahamutConfig.ReSTServer.IdleTimeout = 120 * time.Second
	bahamutConfig.ReSTServer.ListenAddress = listenAddress
	bahamutConfig.ReSTServer.ReadTimeout = 5 * time.Second
	bahamutConfig.ReSTServer.WriteTimeout = 5 * time.Second
	bahamutConfig.TLS.ClientCAPool = caPool
	bahamutConfig.TLS.RootCAPool = caPool
	bahamutConfig.TLS.ServerCertificates = serverCertificates
	bahamutConfig.TLS.AuthType = tls.RequireAndVerifyClientCert
	bahamutConfig.ProfilingServer.Disabled = true
	bahamutConfig.HealthServer.Disabled = !enableHealth
	bahamutConfig.HealthServer.HealthHandler = healthHandlerFunc
	bahamutConfig.HealthServer.ListenAddress = healtListenAddress
	bahamutConfig.Model.RelationshipsRegistry = map[int]elemental.RelationshipsRegistry{0: rufusmodels.Relationships()}
	bahamutConfig.Model.IdentifiablesFactory = func(i string, version int) elemental.Identifiable { return rufusmodels.IdentifiableForIdentity(i) }
	bahamutConfig.WebSocketServer.Disabled = true

	// Create a new Bahamut Server
	server := bahamut.NewServer(bahamutConfig)

	// Register all the processors
	bahamut.RegisterProcessorOrDie(server, newRemoteProcessorProcessor(pluginsRegistry), rufusmodels.RemoteProcessorIdentity)

	return server
}
