// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume


type CsiVolumeTopologyRequest struct {
	// preferred block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#preferred CsiVolume#preferred}
	Preferred *CsiVolumeTopologyRequestPreferred `field:"optional" json:"preferred" yaml:"preferred"`
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#required CsiVolume#required}
	Required *CsiVolumeTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

