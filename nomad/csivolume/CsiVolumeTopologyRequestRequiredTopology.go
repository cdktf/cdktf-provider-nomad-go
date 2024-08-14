// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume


type CsiVolumeTopologyRequestRequiredTopology struct {
	// Define the attributes for the topology request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.1/docs/resources/csi_volume#segments CsiVolume#segments}
	Segments *map[string]*string `field:"required" json:"segments" yaml:"segments"`
}

