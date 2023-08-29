// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalvolume


type ExternalVolumeCapability struct {
	// Defines whether a volume should be available concurrently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/external_volume#access_mode ExternalVolume#access_mode}
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// The storage API that will be used by the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/1.4.20/docs/resources/external_volume#attachment_mode ExternalVolume#attachment_mode}
	AttachmentMode *string `field:"required" json:"attachmentMode" yaml:"attachmentMode"`
}

