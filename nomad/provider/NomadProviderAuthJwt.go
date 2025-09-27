// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type NomadProviderAuthJwt struct {
	// The name of the auth method to use for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs#auth_method NomadProvider#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
	// The externally issued authentication token to be exchanged for a Nomad ACL Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.1/docs#login_token NomadProvider#login_token}
	LoginToken *string `field:"required" json:"loginToken" yaml:"loginToken"`
}

