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

// BaseConfiguration represent the bare minimal configuration to get a phoenix server running
type BaseConfiguration struct {
	CAPoolPath               string `mapstructure:"cacert"              desc:"Path to the CA certificate"                   required:"true"`
	ListenAddress            string `mapstructure:"listen"              desc:"Listening address"                            default:":443"`
	ServerCertificatePath    string `mapstructure:"server-cert"         desc:"Path to the server certificate"               required:"true"`
	ServerCertificateKeyPath string `mapstructure:"server-cert-key"     desc:"Path to the server certificate key"           required:"true"`
}
