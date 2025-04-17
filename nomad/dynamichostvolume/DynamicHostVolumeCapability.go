// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamichostvolume


type DynamicHostVolumeCapability struct {
	// An access mode available for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.0/docs/resources/dynamic_host_volume#access_mode DynamicHostVolume#access_mode}
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// An attachment mode available for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.0/docs/resources/dynamic_host_volume#attachment_mode DynamicHostVolume#attachment_mode}
	AttachmentMode *string `field:"required" json:"attachmentMode" yaml:"attachmentMode"`
}

