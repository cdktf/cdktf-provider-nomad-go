// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume


type CsiVolumeCapability struct {
	// Defines whether a volume should be available concurrently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#access_mode CsiVolume#access_mode}
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// The storage API that will be used by the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/csi_volume#attachment_mode CsiVolume#attachment_mode}
	AttachmentMode *string `field:"required" json:"attachmentMode" yaml:"attachmentMode"`
}

