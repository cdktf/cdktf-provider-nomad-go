// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamichostvolumeregistration


type DynamicHostVolumeRegistrationCapability struct {
	// An access mode available for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.2/docs/resources/dynamic_host_volume_registration#access_mode DynamicHostVolumeRegistration#access_mode}
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// An attachment mode available for the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.5.2/docs/resources/dynamic_host_volume_registration#attachment_mode DynamicHostVolumeRegistration#attachment_mode}
	AttachmentMode *string `field:"required" json:"attachmentMode" yaml:"attachmentMode"`
}

