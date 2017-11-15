package phoenix

// BaseConfiguration represent the bare minimal configuration to get a phoenix server running
type BaseConfiguration struct {
	CAPoolPath               string `mapstructure:"cacert"              desc:"Path to the CA certificate"                   required:"true"`
	ListenAddress            string `mapstructure:"listen"              desc:"Listening address"                            default:":443"`
	ServerCertificatePath    string `mapstructure:"server-cert"         desc:"Path to the server certificate"               required:"true"`
	ServerCertificateKeyPath string `mapstructure:"server-cert-key"     desc:"Path to the server certificate key"           required:"true"`
}
