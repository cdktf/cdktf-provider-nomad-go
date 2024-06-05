// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csivolume


type CsiVolumeMountOptions struct {
	// The file system type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/csi_volume#fs_type CsiVolume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The flags passed to mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.3.0/docs/resources/csi_volume#mount_flags CsiVolume#mount_flags}
	MountFlags *[]*string `field:"optional" json:"mountFlags" yaml:"mountFlags"`
}

