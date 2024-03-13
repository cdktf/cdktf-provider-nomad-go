// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolumeregistration


type CsiVolumeRegistrationTopologyRequest struct {
	// required block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/csi_volume_registration#required CsiVolumeRegistration#required}
	Required *CsiVolumeRegistrationTopologyRequestRequired `field:"optional" json:"required" yaml:"required"`
}

