// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeMountOptions struct {
	// The file system type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/volume#fs_type Volume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The flags passed to mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/nomad/2.4.0/docs/resources/volume#mount_flags Volume#mount_flags}
	MountFlags *[]*string `field:"optional" json:"mountFlags" yaml:"mountFlags"`
}

