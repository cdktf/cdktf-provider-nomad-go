// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type NomadProviderHeaders struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs#name NomadProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs#value NomadProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

