// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalvolume


type ExternalVolumeTopologyRequest struct {
	// preferred block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/external_volume#preferred ExternalVolume#preferred}
	Preferred *ExternalVolumeTopologyRequestPreferred `field:"optional" json:"preferred" yaml:"preferred"`
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/external_volume#required ExternalVolume#required}
	Required *ExternalVolumeTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

