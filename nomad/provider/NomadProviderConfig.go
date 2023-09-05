// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type NomadProviderConfig struct {
	// URL of the root of the target Nomad agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#address NomadProvider#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#alias NomadProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A path to a PEM-encoded certificate authority used to verify the remote agent's certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#ca_file NomadProvider#ca_file}
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// PEM-encoded certificate authority used to verify the remote agent's certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#ca_pem NomadProvider#ca_pem}
	CaPem *string `field:"optional" json:"caPem" yaml:"caPem"`
	// A path to a PEM-encoded certificate provided to the remote agent; requires use of key_file or key_pem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#cert_file NomadProvider#cert_file}
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// PEM-encoded certificate provided to the remote agent; requires use of key_file or key_pem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#cert_pem NomadProvider#cert_pem}
	CertPem *string `field:"optional" json:"certPem" yaml:"certPem"`
	// Consul token to validate Consul Connect Service Identity policies specified in the job file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#consul_token NomadProvider#consul_token}
	ConsulToken *string `field:"optional" json:"consulToken" yaml:"consulToken"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#headers NomadProvider#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// HTTP basic auth configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#http_auth NomadProvider#http_auth}
	HttpAuth *string `field:"optional" json:"httpAuth" yaml:"httpAuth"`
	// A set of environment variables that are ignored by the provider when configuring the Nomad API client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#ignore_env_vars NomadProvider#ignore_env_vars}
	IgnoreEnvVars *map[string]interface{} `field:"optional" json:"ignoreEnvVars" yaml:"ignoreEnvVars"`
	// A path to a PEM-encoded private key, required if cert_file or cert_pem is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#key_file NomadProvider#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
	// PEM-encoded private key, required if cert_file or cert_pem is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#key_pem NomadProvider#key_pem}
	KeyPem *string `field:"optional" json:"keyPem" yaml:"keyPem"`
	// Region of the target Nomad agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#region NomadProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ACL token secret for API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#secret_id NomadProvider#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// Skip TLS verification on client side.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#skip_verify NomadProvider#skip_verify}
	SkipVerify interface{} `field:"optional" json:"skipVerify" yaml:"skipVerify"`
	// Vault token if policies are specified in the job file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs#vault_token NomadProvider#vault_token}
	VaultToken *string `field:"optional" json:"vaultToken" yaml:"vaultToken"`
}

