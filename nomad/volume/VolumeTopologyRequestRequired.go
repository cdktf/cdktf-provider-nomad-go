// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeTopologyRequestRequired struct {
	// topology block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.2.0/docs/resources/volume#topology Volume#topology}
	Topology interface{} `field:"required" json:"topology" yaml:"topology"`
}

