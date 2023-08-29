// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeTopologyRequestRequiredTopology struct {
	// Define attributes for the topology request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/volume#segments Volume#segments}
	Segments *map[string]*string `field:"required" json:"segments" yaml:"segments"`
}

