// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeTopologyRequest struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/volume#required Volume#required}
	Required *VolumeTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

