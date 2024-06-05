// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolumeregistration


type CsiVolumeRegistrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/csi_volume_registration#create CsiVolumeRegistration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/csi_volume_registration#delete CsiVolumeRegistration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

