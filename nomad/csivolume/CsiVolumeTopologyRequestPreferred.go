// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume


type CsiVolumeTopologyRequestPreferred struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/csi_volume#topology CsiVolume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

