package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/aporeto-inc/addedeffect/lombric"
)

// Configuration is a struct used to load the config file
type Configuration struct {
	CAPoolPath               string `mapstructure:"cacert"              desc:"Path to the CA certificate"                   required:"true"`
	ListenAddress            string `mapstructure:"listen"              desc:"Listening address"                            default:":443"`
	ServerCertificatePath    string `mapstructure:"server-cert"         desc:"Path to the server certificate"               required:"true"`
	ServerCertificateKeyPath string `mapstructure:"server-cert-key"     desc:"Path to the server certificate key"           required:"true"`

	CAPool            *x509.CertPool
	ServerCertificate tls.Certificate
}

// Prefix returns the configuration prefix.
func (c *Configuration) Prefix() string { return "phoenixexample" }

// NewConfiguration returns a new Configuration.
func NewConfiguration() *Configuration {

	c := &Configuration{}
	lombric.Initialize(c)

	var err error
	c.CAPool, err = x509.SystemCertPool()
	if err != nil {
		panic("Unable to load system CA pool: " + err.Error())
	}

	cadata, err := ioutil.ReadFile(c.CAPoolPath)
	if err != nil {
		panic("Unable to read CA file: " + err.Error())
	}
	c.CAPool.AppendCertsFromPEM(cadata)

	c.ServerCertificate, err = tls.LoadX509KeyPair(c.ServerCertificatePath, c.ServerCertificateKeyPath)
	if err != nil {
		panic("Unable to load the server certificate: " + err.Error())
	}

	return c
}
