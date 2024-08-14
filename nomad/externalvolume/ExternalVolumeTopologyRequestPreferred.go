// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalvolume


type ExternalVolumeTopologyRequestPreferred struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/external_volume#topology ExternalVolume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

