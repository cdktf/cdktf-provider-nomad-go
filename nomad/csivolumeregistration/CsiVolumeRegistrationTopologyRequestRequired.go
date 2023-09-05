// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolumeregistration


type CsiVolumeRegistrationTopologyRequestRequired struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.0.0/docs/resources/csi_volume_registration#topology CsiVolumeRegistration#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

