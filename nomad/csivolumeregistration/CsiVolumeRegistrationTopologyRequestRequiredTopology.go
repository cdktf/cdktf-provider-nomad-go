// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolumeregistration


type CsiVolumeRegistrationTopologyRequestRequiredTopology struct {
	// Define attributes for the topology request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/csi_volume_registration#segments CsiVolumeRegistration#segments}
	Segments *map[string]*string `field:"required" json:"segments" yaml:"segments"`
}

