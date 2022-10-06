package externalvolume


type ExternalVolumeMountOptions struct {
	// The file system type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/external_volume#fs_type ExternalVolume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The flags passed to mount.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/nomad/r/external_volume#mount_flags ExternalVolume#mount_flags}
	MountFlags *[]*string `field:"optional" json:"mountFlags" yaml:"mountFlags"`
}

