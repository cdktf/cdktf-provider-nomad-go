// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aclauthmethod


type AclAuthMethodConfigOidcClientAssertion struct {
	// The source of the key Nomad will use to sign the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#key_source AclAuthMethod#key_source}
	KeySource *string `field:"required" json:"keySource" yaml:"keySource"`
	// List of audiences to accept the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#audience AclAuthMethod#audience}
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// Additional headers to include on the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#extra_headers AclAuthMethod#extra_headers}
	ExtraHeaders *map[string]*string `field:"optional" json:"extraHeaders" yaml:"extraHeaders"`
	// Algorithm of the key used to sign the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#key_algorithm AclAuthMethod#key_algorithm}
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// private_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs/resources/acl_auth_method#private_key AclAuthMethod#private_key}
	PrivateKey *AclAuthMethodConfigOidcClientAssertionPrivateKey `field:"optional" json:"privateKey" yaml:"privateKey"`
}

