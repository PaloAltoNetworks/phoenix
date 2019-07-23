// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package phoenix

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia"
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

	modelManagers := map[int]elemental.ModelManager{0: gaia.Manager(), 1: gaia.Manager()}

	options := []bahamut.Option{
		bahamut.OptRestServer(listenAddress),
		bahamut.OptTLS(serverCertificates, nil),
		bahamut.OptMTLS(caPool, tls.RequireAndVerifyClientCert),
		bahamut.OptModel(modelManagers),
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
