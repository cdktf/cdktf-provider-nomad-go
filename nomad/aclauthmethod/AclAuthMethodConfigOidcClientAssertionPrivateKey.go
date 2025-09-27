// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod


type AclAuthMethodConfigOidcClientAssertionPrivateKey struct {
	// Specific 'kid' header to set on the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#key_id AclAuthMethod#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Name of the header the IDP will use to find the cert to verify the JWT signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#key_id_header AclAuthMethod#key_id_header}
	KeyIdHeader *string `field:"optional" json:"keyIdHeader" yaml:"keyIdHeader"`
	// An x509 certificate PEM to derive a key ID header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#pem_cert AclAuthMethod#pem_cert}
	PemCert *string `field:"optional" json:"pemCert" yaml:"pemCert"`
	// Path to an x509 certificate PEM on Nomad servers to derive a key ID header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#pem_cert_file AclAuthMethod#pem_cert_file}
	PemCertFile *string `field:"optional" json:"pemCertFile" yaml:"pemCertFile"`
	// RSA private key PEM to use to sign the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#pem_key AclAuthMethod#pem_key}
	PemKey *string `field:"optional" json:"pemKey" yaml:"pemKey"`
	// Path to an RSA private key PEM on Nomad servers to use to sign the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#pem_key_file AclAuthMethod#pem_key_file}
	PemKeyFile *string `field:"optional" json:"pemKeyFile" yaml:"pemKeyFile"`
}

