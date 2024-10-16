// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalvolume


type ExternalVolumeTopologyRequestRequiredTopology struct {
	// Define the attributes for the topology request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/external_volume#segments ExternalVolume#segments}
	Segments *map[string]*string `field:"required" json:"segments" yaml:"segments"`
}

