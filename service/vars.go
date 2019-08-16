package service

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"github.com/zoenion/common/conf"
	configpb "github.com/zoenion/common/proto/config"
	servicepb "github.com/zoenion/common/proto/service"
	"google.golang.org/grpc/credentials"
)

type Vars struct {
	Name     string
	Dir      string
	Domain   string
	IP       string
	Insecure bool

	ConfigServer    string
	configClient    configpb.ConfigClient
	configChanInput chan conf.Map
	lastState       conf.Map

	Registry       string
	Namespace      string
	RegistryID     string
	registryClient servicepb.RegistryClient

	GRPCAuthorityAddress        string
	AuthorityCertPath           string
	AuthorityCredentials        string
	authorityCert               *x509.Certificate
	authorityGRPCAuthentication credentials.PerRPCCredentials

	GatewayGRPCPort string
	GatewayHTTPPort string

	gRPCAuthorityCredentials credentials.TransportCredentials
	serviceCert              *x509.Certificate
	serviceKey               crypto.PrivateKey
}

type ConfigVars struct {
	Dir  string
	Name string
}

func (v *Vars) GRPCAddress() string {
	addr := v.Domain
	if addr == "" {
		addr = v.IP
	}
	return fmt.Sprintf("%s:%s", addr, v.GatewayGRPCPort)
}

func (v *Vars) HTTPAddress() string {
	addr := v.Domain
	if addr == "" {
		addr = v.IP
	}
	return fmt.Sprintf("%s:%s", addr, v.GatewayHTTPPort)
}